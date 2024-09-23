package orchestrator

import (
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/thirdweb-dev/indexer/internal/common"
	"github.com/thirdweb-dev/indexer/internal/storage"
	"github.com/thirdweb-dev/indexer/internal/worker"
)

const DEFAULT_FAILURES_PER_POLL = 10
const DEFAULT_FAILURE_TRIGGER_INTERVAL = 1000

type FailureRecoverer struct {
	failuresPerPoll   int
	triggerIntervalMs int
	storage           storage.IStorage
	rpc               common.RPC
}

func NewFailureRecoverer(rpc common.RPC, storage storage.IStorage) *FailureRecoverer {
	failuresPerPoll, err := strconv.Atoi(os.Getenv("FAILURES_PER_POLL"))
	if err != nil || failuresPerPoll == 0 {
		failuresPerPoll = DEFAULT_FAILURES_PER_POLL
	}
	triggerInterval, err := strconv.Atoi(os.Getenv("FAILURE_TRIGGER_INTERVAL"))
	if err != nil || triggerInterval == 0 {
		triggerInterval = DEFAULT_FAILURE_TRIGGER_INTERVAL
	}
	return &FailureRecoverer{
		triggerIntervalMs: triggerInterval,
		failuresPerPoll:   failuresPerPoll,
		storage:           storage,
		rpc:               rpc,
	}
}

func (fr *FailureRecoverer) Start() {
	interval := time.Duration(fr.triggerIntervalMs) * time.Millisecond
	ticker := time.NewTicker(interval)

	log.Debug().Msgf("Failure Recovery running")
	go func() {
		for range ticker.C {
			blockFailures, err := fr.storage.OrchestratorStorage.GetBlockFailures(fr.failuresPerPoll)
			if err != nil {
				log.Error().Err(err).Msg("Failed to get block failures")
				continue
			}

			blocksToTrigger := make([]*big.Int, 0, len(blockFailures))
			for _, blockFailure := range blockFailures {
				blocksToTrigger = append(blocksToTrigger, blockFailure.BlockNumber)
			}

			log.Debug().Msgf("Triggering Failure Recoverer for trigger: %v", blocksToTrigger)
			worker := worker.NewWorker(fr.rpc, fr.storage)
			results := worker.Run(blocksToTrigger)
			fr.handleBlockResults(blockFailures, results)
		}
	}()

	// Keep the program running (otherwise it will exit)
	select {}
}

func (fr *FailureRecoverer) handleBlockResults(blockFailures []common.BlockFailure, results []worker.BlockResult) {
	err := fr.storage.OrchestratorStorage.DeleteBlockFailures(blockFailures)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting block failures")
		return
	}
	blockFailureMap := make(map[*big.Int]common.BlockFailure)
	for _, failure := range blockFailures {
		blockFailureMap[failure.BlockNumber] = failure
	}
	var newBlockFailures []common.BlockFailure
	for _, result := range results {
		if result.Error != nil {
			prevBlockFailure, ok := blockFailureMap[result.BlockNumber]
			failureCount := 1
			if ok {
				failureCount = prevBlockFailure.FailureCount + 1
			}
			newBlockFailures = append(newBlockFailures, common.BlockFailure{
				BlockNumber:   result.BlockNumber,
				FailureReason: result.Error.Error(),
				FailureTime:   time.Now(),
				ChainId:       fr.rpc.ChainID,
				FailureCount:  failureCount,
			})
		}
	}
	fr.storage.OrchestratorStorage.StoreBlockFailures(newBlockFailures)
}
