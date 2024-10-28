package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	configs "github.com/thirdweb-dev/indexer/configs"
	customLogger "github.com/thirdweb-dev/indexer/internal/log"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "indexer",
		Short: "TBD",
		Long:  "TBD",
		Run: func(cmd *cobra.Command, args []string) {
			go func() {
				RunOrchestrator(cmd, args)
			}()
			RunApi(cmd, args)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./configs/config.yml)")
	rootCmd.PersistentFlags().String("rpc-url", "", "RPC Url to use for the indexer")
	rootCmd.PersistentFlags().Int("rpc-blocks-blocksPerRequest", 0, "How many blocks to fetch per request")
	rootCmd.PersistentFlags().Int("rpc-blocks-batchDelay", 0, "Milliseconds to wait between batches of blocks when fetching from the RPC")
	rootCmd.PersistentFlags().Int("rpc-logs-blocksPerRequest", 0, "How many blocks to fetch logs per request")
	rootCmd.PersistentFlags().Int("rpc-logs-batchDelay", 0, "Milliseconds to wait between batches of logs when fetching from the RPC")
	rootCmd.PersistentFlags().Bool("rpc-blockReceipts-enabled", false, "Whether to enable fetching block receipts from the RPC")
	rootCmd.PersistentFlags().Int("rpc-blockReceipts-blocksPerRequest", 0, "How many blocks to fetch receipts for per request")
	rootCmd.PersistentFlags().Int("rpc-blockReceipts-batchDelay", 0, "Milliseconds to wait between batches of receipts when fetching from the RPC")
	rootCmd.PersistentFlags().Bool("rpc-traces-enabled", true, "Whether to enable fetching traces from the RPC")
	rootCmd.PersistentFlags().Int("rpc-traces-blocksPerRequest", 0, "How many blocks to fetch traces per request")
	rootCmd.PersistentFlags().Int("rpc-traces-batchDelay", 0, "Milliseconds to wait between batches of traces when fetching from the RPC")
	rootCmd.PersistentFlags().String("log-level", "", "Log level to use for the application")
	rootCmd.PersistentFlags().Bool("log-prettify", false, "Whether to prettify the log output")
	rootCmd.PersistentFlags().Bool("poller-enabled", true, "Toggle poller")
	rootCmd.PersistentFlags().Bool("poller-interval", true, "Poller interval")
	rootCmd.PersistentFlags().Int("poller-blocks-per-poll", 10, "How many blocks to poll each interval")
	rootCmd.PersistentFlags().Int("poller-from-block", 0, "From which block to start polling")
	rootCmd.PersistentFlags().Bool("poller-force-from-block", false, "Force the poller to start from the block specified in `poller-from-block`")
	rootCmd.PersistentFlags().Int("poller-until-block", 0, "Until which block to poll")
	rootCmd.PersistentFlags().Int("poller-parallel-pollers", 5, "Maximum number of parallel pollers")
	rootCmd.PersistentFlags().Bool("committer-enabled", true, "Toggle committer")
	rootCmd.PersistentFlags().Int("committer-blocks-per-commit", 10, "How many blocks to commit each interval")
	rootCmd.PersistentFlags().Int("committer-interval", 1000, "How often to commit blocks in milliseconds")
	rootCmd.PersistentFlags().Int("committer-from-block", 0, "From which block to start committing")
	rootCmd.PersistentFlags().Bool("reorgHandler-enabled", true, "Toggle reorg handler")
	rootCmd.PersistentFlags().Int("reorgHandler-interval", 1000, "How often to run reorg handler in milliseconds")
	rootCmd.PersistentFlags().Int("reorgHandler-blocks-per-scan", 100, "How many blocks to scan for reorgs")
	rootCmd.PersistentFlags().Int("reorgHandler-from-block", 0, "From which block to start scanning for reorgs")
	rootCmd.PersistentFlags().Bool("reorgHandler-force-from-block", false, "Force the reorg handler to start from the block specified in `reorgHandler-from-block`")
	rootCmd.PersistentFlags().Bool("failure-recoverer-enabled", true, "Toggle failure recoverer")
	rootCmd.PersistentFlags().Int("failure-recoverer-blocks-per-run", 10, "How many blocks to run failure recoverer for")
	rootCmd.PersistentFlags().Int("failure-recoverer-interval", 1000, "How often to run failure recoverer in milliseconds")
	rootCmd.PersistentFlags().String("storage-staging-clickhouse-database", "", "Clickhouse database for staging storage")
	rootCmd.PersistentFlags().Int("storage-staging-clickhouse-port", 0, "Clickhouse port for staging storage")
	rootCmd.PersistentFlags().String("storage-main-clickhouse-database", "", "Clickhouse database for main storage")
	rootCmd.PersistentFlags().Int("storage-main-clickhouse-port", 0, "Clickhouse port for main storage")
	rootCmd.PersistentFlags().String("storage-orchestrator-clickhouse-database", "", "Clickhouse database for orchestrator storage")
	rootCmd.PersistentFlags().Int("storage-orchestrator-clickhouse-port", 0, "Clickhouse port for orchestrator storage")
	rootCmd.PersistentFlags().String("storage-orchestrator-clickhouse-host", "", "Clickhouse host for orchestrator storage")
	rootCmd.PersistentFlags().String("storage-orchestrator-clickhouse-username", "", "Clickhouse username for orchestrator storage")
	rootCmd.PersistentFlags().String("storage-orchestrator-clickhouse-password", "", "Clickhouse password for orchestrator storage")
	rootCmd.PersistentFlags().Bool("storage-orchestrator-clickhouse-asyncInsert", false, "Clickhouse async insert for orchestrator storage")
	rootCmd.PersistentFlags().Int("storage-orchestrator-clickhouse-maxRowsPerInsert", 100000, "Clickhouse max rows per insert for orchestrator storage")
	rootCmd.PersistentFlags().Bool("storage-orchestrator-clickhouse-disableTLS", false, "Clickhouse disableTLS for orchestrator storage")
	rootCmd.PersistentFlags().Int("storage-orchestrator-memory-maxItems", 0, "Max items for orchestrator memory storage")
	rootCmd.PersistentFlags().Int("storage-orchestrator-redis-poolSize", 0, "Redis pool size for orchestrator storage")
	rootCmd.PersistentFlags().String("storage-orchestrator-redis-addr", "", "Redis address for orchestrator storage")
	rootCmd.PersistentFlags().String("storage-orchestrator-redis-password", "", "Redis password for orchestrator storage")
	rootCmd.PersistentFlags().Int("storage-orchestrator-redis-db", 0, "Redis db for orchestrator storage")
	rootCmd.PersistentFlags().String("storage-staging-clickhouse-host", "", "Clickhouse host for staging storage")
	rootCmd.PersistentFlags().String("storage-main-clickhouse-host", "", "Clickhouse host for main storage")
	rootCmd.PersistentFlags().String("storage-main-clickhouse-username", "", "Clickhouse username for main storage")
	rootCmd.PersistentFlags().String("storage-main-clickhouse-password", "", "Clickhouse password for main storage")
	rootCmd.PersistentFlags().Bool("storage-main-clickhouse-asyncInsert", false, "Clickhouse async insert for main storage")
	rootCmd.PersistentFlags().Int("storage-main-clickhouse-maxRowsPerInsert", 100000, "Clickhouse max rows per insert for main storage")
	rootCmd.PersistentFlags().Bool("storage-main-clickhouse-disableTLS", false, "Clickhouse disableTLS for main storage")
	rootCmd.PersistentFlags().String("storage-staging-clickhouse-username", "", "Clickhouse username for staging storage")
	rootCmd.PersistentFlags().String("storage-staging-clickhouse-password", "", "Clickhouse password for staging storage")
	rootCmd.PersistentFlags().Bool("storage-staging-clickhouse-asyncInsert", false, "Clickhouse async insert for staging storage")
	rootCmd.PersistentFlags().Int("storage-staging-clickhouse-maxRowsPerInsert", 100000, "Clickhouse max rows per insert for staging storage")
	rootCmd.PersistentFlags().Bool("storage-staging-clickhouse-disableTLS", false, "Clickhouse disableTLS for staging storage")
	rootCmd.PersistentFlags().String("api-host", "localhost:3000", "API host")
	rootCmd.PersistentFlags().String("api-basicAuth-username", "", "API basic auth username")
	rootCmd.PersistentFlags().String("api-basicAuth-password", "", "API basic auth password")
	viper.BindPFlag("rpc.url", rootCmd.PersistentFlags().Lookup("rpc-url"))
	viper.BindPFlag("rpc.blocks.blocksPerRequest", rootCmd.PersistentFlags().Lookup("rpc-blocks-blocksPerRequest"))
	viper.BindPFlag("rpc.blocks.batchDelay", rootCmd.PersistentFlags().Lookup("rpc-blocks-batchDelay"))
	viper.BindPFlag("rpc.logs.blocksPerRequest", rootCmd.PersistentFlags().Lookup("rpc-logs-blocksPerRequest"))
	viper.BindPFlag("rpc.logs.batchDelay", rootCmd.PersistentFlags().Lookup("rpc-logs-batchDelay"))
	viper.BindPFlag("rpc.blockReceipts.enabled", rootCmd.PersistentFlags().Lookup("rpc-blockReceipts-enabled"))
	viper.BindPFlag("rpc.blockReceipts.blocksPerRequest", rootCmd.PersistentFlags().Lookup("rpc-blockReceipts-blocksPerRequest"))
	viper.BindPFlag("rpc.blockReceipts.batchDelay", rootCmd.PersistentFlags().Lookup("rpc-blockReceipts-batchDelay"))
	viper.BindPFlag("rpc.traces.enabled", rootCmd.PersistentFlags().Lookup("rpc-traces-enabled"))
	viper.BindPFlag("rpc.traces.blocksPerRequest", rootCmd.PersistentFlags().Lookup("rpc-traces-blocksPerRequest"))
	viper.BindPFlag("rpc.traces.batchDelay", rootCmd.PersistentFlags().Lookup("rpc-traces-batchDelay"))
	viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log-level"))
	viper.BindPFlag("log.prettify", rootCmd.PersistentFlags().Lookup("log-prettify"))
	viper.BindPFlag("poller.enabled", rootCmd.PersistentFlags().Lookup("poller-enabled"))
	viper.BindPFlag("poller.interval", rootCmd.PersistentFlags().Lookup("poller-interval"))
	viper.BindPFlag("poller.blocksPerPoll", rootCmd.PersistentFlags().Lookup("poller-blocks-per-poll"))
	viper.BindPFlag("poller.fromBlock", rootCmd.PersistentFlags().Lookup("poller-from-block"))
	viper.BindPFlag("poller.forceFromBlock", rootCmd.PersistentFlags().Lookup("poller-force-from-block"))
	viper.BindPFlag("poller.untilBlock", rootCmd.PersistentFlags().Lookup("poller-until-block"))
	viper.BindPFlag("poller.parallelPollers", rootCmd.PersistentFlags().Lookup("poller-parallel-pollers"))
	viper.BindPFlag("committer.enabled", rootCmd.PersistentFlags().Lookup("committer-enabled"))
	viper.BindPFlag("committer.blocksPerCommit", rootCmd.PersistentFlags().Lookup("committer-blocks-per-commit"))
	viper.BindPFlag("committer.interval", rootCmd.PersistentFlags().Lookup("committer-interval"))
	viper.BindPFlag("committer.fromBlock", rootCmd.PersistentFlags().Lookup("committer-from-block"))
	viper.BindPFlag("reorgHandler.enabled", rootCmd.PersistentFlags().Lookup("reorgHandler-enabled"))
	viper.BindPFlag("reorgHandler.interval", rootCmd.PersistentFlags().Lookup("reorgHandler-interval"))
	viper.BindPFlag("reorgHandler.blocksPerScan", rootCmd.PersistentFlags().Lookup("reorgHandler-blocks-per-scan"))
	viper.BindPFlag("reorgHandler.fromBlock", rootCmd.PersistentFlags().Lookup("reorgHandler-from-block"))
	viper.BindPFlag("reorgHandler.forceFromBlock", rootCmd.PersistentFlags().Lookup("reorgHandler-force-from-block"))
	viper.BindPFlag("failureRecoverer.enabled", rootCmd.PersistentFlags().Lookup("failure-recoverer-enabled"))
	viper.BindPFlag("failureRecoverer.blocksPerRun", rootCmd.PersistentFlags().Lookup("failure-recoverer-blocks-per-run"))
	viper.BindPFlag("failureRecoverer.interval", rootCmd.PersistentFlags().Lookup("failure-recoverer-interval"))
	viper.BindPFlag("storage.staging.clickhouse.database", rootCmd.PersistentFlags().Lookup("storage-staging-clickhouse-database"))
	viper.BindPFlag("storage.staging.clickhouse.host", rootCmd.PersistentFlags().Lookup("storage-staging-clickhouse-host"))
	viper.BindPFlag("storage.staging.clickhouse.port", rootCmd.PersistentFlags().Lookup("storage-staging-clickhouse-port"))
	viper.BindPFlag("storage.staging.clickhouse.username", rootCmd.PersistentFlags().Lookup("storage-staging-clickhouse-username"))
	viper.BindPFlag("storage.staging.clickhouse.password", rootCmd.PersistentFlags().Lookup("storage-staging-clickhouse-password"))
	viper.BindPFlag("storage.staging.clickhouse.asyncInsert", rootCmd.PersistentFlags().Lookup("storage-staging-clickhouse-asyncInsert"))
	viper.BindPFlag("storage.staging.clickhouse.maxRowsPerInsert", rootCmd.PersistentFlags().Lookup("storage-staging-clickhouse-maxRowsPerInsert"))
	viper.BindPFlag("storage.staging.clickhouse.disableTLS", rootCmd.PersistentFlags().Lookup("storage-staging-clickhouse-disableTLS"))
	viper.BindPFlag("storage.main.clickhouse.database", rootCmd.PersistentFlags().Lookup("storage-main-clickhouse-database"))
	viper.BindPFlag("storage.main.clickhouse.host", rootCmd.PersistentFlags().Lookup("storage-main-clickhouse-host"))
	viper.BindPFlag("storage.main.clickhouse.port", rootCmd.PersistentFlags().Lookup("storage-main-clickhouse-port"))
	viper.BindPFlag("storage.main.clickhouse.username", rootCmd.PersistentFlags().Lookup("storage-main-clickhouse-username"))
	viper.BindPFlag("storage.main.clickhouse.password", rootCmd.PersistentFlags().Lookup("storage-main-clickhouse-password"))
	viper.BindPFlag("storage.main.clickhouse.asyncInsert", rootCmd.PersistentFlags().Lookup("storage-main-clickhouse-asyncInsert"))
	viper.BindPFlag("storage.main.clickhouse.maxRowsPerInsert", rootCmd.PersistentFlags().Lookup("storage-main-clickhouse-maxRowsPerInsert"))
	viper.BindPFlag("storage.main.clickhouse.disableTLS", rootCmd.PersistentFlags().Lookup("storage-main-clickhouse-disableTLS"))
	viper.BindPFlag("storage.orchestrator.clickhouse.database", rootCmd.PersistentFlags().Lookup("storage-orchestrator-clickhouse-database"))
	viper.BindPFlag("storage.orchestrator.clickhouse.host", rootCmd.PersistentFlags().Lookup("storage-orchestrator-clickhouse-host"))
	viper.BindPFlag("storage.orchestrator.clickhouse.port", rootCmd.PersistentFlags().Lookup("storage-orchestrator-clickhouse-port"))
	viper.BindPFlag("storage.orchestrator.clickhouse.username", rootCmd.PersistentFlags().Lookup("storage-orchestrator-clickhouse-username"))
	viper.BindPFlag("storage.orchestrator.clickhouse.password", rootCmd.PersistentFlags().Lookup("storage-orchestrator-clickhouse-password"))
	viper.BindPFlag("storage.orchestrator.clickhouse.asyncInsert", rootCmd.PersistentFlags().Lookup("storage-orchestrator-clickhouse-asyncInsert"))
	viper.BindPFlag("storage.orchestrator.clickhouse.maxRowsPerInsert", rootCmd.PersistentFlags().Lookup("storage-orchestrator-clickhouse-maxRowsPerInsert"))
	viper.BindPFlag("storage.orchestrator.clickhouse.disableTLS", rootCmd.PersistentFlags().Lookup("storage-orchestrator-clickhouse-disableTLS"))
	viper.BindPFlag("storage.orchestrator.memory.maxItems", rootCmd.PersistentFlags().Lookup("storage-orchestrator-memory-maxItems"))
	viper.BindPFlag("storage.orchestrator.redis.poolSize", rootCmd.PersistentFlags().Lookup("storage-orchestrator-redis-poolSize"))
	viper.BindPFlag("storage.orchestrator.redis.addr", rootCmd.PersistentFlags().Lookup("storage-orchestrator-redis-addr"))
	viper.BindPFlag("storage.orchestrator.redis.password", rootCmd.PersistentFlags().Lookup("storage-orchestrator-redis-password"))
	viper.BindPFlag("storage.orchestrator.redis.db", rootCmd.PersistentFlags().Lookup("storage-orchestrator-redis-db"))
	viper.BindPFlag("api.host", rootCmd.PersistentFlags().Lookup("api-host"))
	viper.BindPFlag("api.basicAuth.username", rootCmd.PersistentFlags().Lookup("api-basicAuth-username"))
	viper.BindPFlag("api.basicAuth.password", rootCmd.PersistentFlags().Lookup("api-basicAuth-password"))
	rootCmd.AddCommand(orchestratorCmd)
	rootCmd.AddCommand(apiCmd)
}

func initConfig() {
	configs.LoadConfig(cfgFile)
	customLogger.InitLogger()
}
