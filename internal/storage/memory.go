package storage

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"

	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/thirdweb-dev/indexer/internal/common"
)

type MemoryConnectorConfig struct {
	MaxItems int
	Prefix   string
}

type MemoryConnector struct {
	cache *lru.Cache[string, string]
}

func NewMemoryConnector(cfg *MemoryConnectorConfig) (*MemoryConnector, error) {
	if cfg != nil && cfg.MaxItems <= 0 {
		return nil, fmt.Errorf("maxItems must be greater than 0")
	}

	maxItems := 1000
	if cfg != nil && cfg.MaxItems > 0 {
		maxItems = cfg.MaxItems
	}

	cache, err := lru.New[string, string](maxItems)
	if err != nil {
		return nil, fmt.Errorf("failed to create LRU cache: %w", err)
	}

	return &MemoryConnector{
		cache: cache,
	}, nil
}

func (m *MemoryConnector) GetLatestPolledBlockNumber() (uint64, error) {
	blockNumber, ok := m.cache.Get("latest_polled_block_number")
	if !ok {
		return math.MaxUint64, nil // this will overflow to genesis
	}
	return strconv.ParseUint(blockNumber, 10, 64)
}

func (m *MemoryConnector) StoreLatestPolledBlockNumber(blockNumber uint64) error {
	m.cache.Add("latest_polled_block_number", strconv.FormatUint(blockNumber, 10))
	return nil
}

func (m *MemoryConnector) StoreBlockFailures(failures []common.BlockFailure) error {
	for _, failure := range failures {
		failureJson, err := common.BlockFailureToString(failure)
		if err != nil {
			return err
		}
		m.cache.Add(fmt.Sprintf("block_failure:%d", failure.BlockNumber), failureJson)
	}
	return nil
}

func (m *MemoryConnector) GetBlockFailures(limit int) ([]common.BlockFailure, error) {
	blockFailures := []common.BlockFailure{}
	for _, key := range m.cache.Keys() {
		if len(blockFailures) >= limit {
			break
		}
		if strings.HasPrefix(key, "block_failure:") {
			value, ok := m.cache.Get(key)
			if ok {
				blockFailure, err := common.StringToBlockFailure(value)
				if err != nil {
					return nil, err
				}
				blockFailures = append(blockFailures, blockFailure)
			}
		}
	}
	return blockFailures, nil
}

func (m *MemoryConnector) DeleteBlockFailures(failures []common.BlockFailure) error {
	for _, failure := range failures {
		m.cache.Remove(fmt.Sprintf("block_failure:%d", failure.BlockNumber))
	}
	return nil
}

func (m *MemoryConnector) InsertBlocks(blocks []common.Block) error {
	for _, block := range blocks {
		blockJson, err := json.Marshal(block)
		if err != nil {
			return err
		}
		m.cache.Add(fmt.Sprintf("block:%d", block.Number), string(blockJson))
	}
	return nil
}

func (m *MemoryConnector) GetBlocks(qf QueryFilter) ([]common.Block, error) {
	blocks := []common.Block{}
	limit := getLimit(qf)
	blockNumbersToCheck := getBlockNumbersToCheck(qf)
	for _, key := range m.cache.Keys() {
		if len(blocks) >= int(limit) {
			break
		}
		if isKeyForBlock(key, "block:", blockNumbersToCheck) {
			value, ok := m.cache.Get(key)
			if ok {
				block := common.Block{}
				err := json.Unmarshal([]byte(value), &block)
				if err != nil {
					return nil, err
				}
				blocks = append(blocks, block)
			}
		}
	}
	return blocks, nil
}

func (m *MemoryConnector) InsertTransactions(txs []common.Transaction) error {
	for _, tx := range txs {
		txJson, err := json.Marshal(tx)
		if err != nil {
			return err
		}
		m.cache.Add(fmt.Sprintf("transaction:%s", tx.Hash), string(txJson))
	}
	return nil
}

func (m *MemoryConnector) GetTransactions(qf QueryFilter) ([]common.Transaction, error) {
	txs := []common.Transaction{}
	limit := getLimit(qf)
	blockNumbersToCheck := getBlockNumbersToCheck(qf)
	for _, key := range m.cache.Keys() {
		if len(txs) >= limit {
			break
		}
		if isKeyForBlock(key, "transaction:", blockNumbersToCheck) {
			value, ok := m.cache.Get(key)
			if ok {
				tx := common.Transaction{}
				err := json.Unmarshal([]byte(value), &tx)
				if err != nil {
					return nil, err
				}
				txs = append(txs, tx)
			}
		}
	}
	return txs, nil
}

func (m *MemoryConnector) InsertEvents(events []common.Log) error {
	for _, event := range events {
		eventJson, err := json.Marshal(event)
		if err != nil {
			return err
		}
		m.cache.Add(fmt.Sprintf("event:%s-%d", event.TransactionHash, event.Index), string(eventJson))
	}
	return nil
}

func (m *MemoryConnector) GetEvents(qf QueryFilter) ([]common.Log, error) {
	events := []common.Log{}
	limit := getLimit(qf)
	blockNumbersToCheck := getBlockNumbersToCheck(qf)
	for _, key := range m.cache.Keys() {
		if len(events) >= limit {
			break
		}
		if isKeyForBlock(key, "event:", blockNumbersToCheck) {
			value, ok := m.cache.Get(key)
			if ok {
				event := common.Log{}
				err := json.Unmarshal([]byte(value), &event)
				if err != nil {
					return nil, err
				}
				events = append(events, event)
			}
		}
	}
	return events, nil
}

func (m *MemoryConnector) GetMaxBlockNumber() (uint64, error) {
	maxBlockNumber := uint64(0)
	for _, key := range m.cache.Keys() {
		if strings.HasPrefix(key, "block:") {
			blockNumber, err := strconv.ParseUint(strings.Split(key, ":")[1], 10, 64)
			if err != nil {
				return 0, err
			}
			if blockNumber > maxBlockNumber {
				maxBlockNumber = blockNumber
			}
		}
	}
	return maxBlockNumber, nil
}

func isKeyForBlock(key string, prefix string, blocksFilter map[uint64]uint8) bool {
	if !strings.HasPrefix(key, prefix) {
		return false
	}
	blockNumber, err := strconv.ParseUint(strings.TrimPrefix(key, prefix), 10, 64)
	if err != nil {
		return false
	}
	if len(blocksFilter) == 0 {
		return true
	}
	_, ok := blocksFilter[blockNumber]
	return ok
}

func getLimit(qf QueryFilter) int {
	limit := qf.Limit
	if limit == 0 {
		limit = math.MaxUint16
	}
	return int(limit)
}

func getBlockNumbersToCheck(qf QueryFilter) map[uint64]uint8 {
	blockNumbersToCheck := make(map[uint64]uint8)
	for _, num := range qf.BlockNumbers {
		blockNumbersToCheck[num] = 1
	}
	return blockNumbersToCheck
}
