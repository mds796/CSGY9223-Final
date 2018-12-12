package cmd

import (
	"github.com/mds796/CSGY9223-Final/storage/raftkv"
	"github.com/spf13/cobra"
)

var raftkvConfig raftkv.Config
var raftkvPidFile string

func init() {
	rootCmd.AddCommand(raftkvCmd)

	raftkvCmd.PersistentFlags().StringVarP(&raftkvPidFile, "pidFile", "p", ".raftkv.pid", "The name of the process ID file.")
	raftkvCmd.AddCommand(startRaftKVCmd)
	raftkvCmd.AddCommand(stopRaftKVCmd)
	raftkvCmd.AddCommand(restartRaftKVCmd)

	raftkvSetStartArgs(startRaftKVCmd)
	raftkvSetStartArgs(restartRaftKVCmd)
	raftkvSetStopArgs(stopRaftKVCmd)
}

func raftkvSetStartArgs(command *cobra.Command) {
	command.Flags().StringVarP(&raftkvConfig.NodeID, "nodeID", "N", "node0", "The unique NodeID identifying this RaftKV node.")
	command.Flags().StringVarP(&raftkvConfig.Host, "host", "H", "localhost", "The host interface to listen on.")
	command.Flags().Uint16VarP(&raftkvConfig.Port, "port", "P", 2379, "The TCP port to listen on.")
	command.Flags().StringVar(&raftkvConfig.JoinHost, "joinHost", "", "The host interface the current leader is listening on.")
	command.Flags().Uint16Var(&raftkvConfig.JoinPort, "joinPort", 0, "The TCP port the current leader is listening on.")
}

func raftkvSetStopArgs(command *cobra.Command) {
	command.Flags().StringVarP(&raftkvConfig.NodeID, "nodeID", "N", "node0", "The unique NodeID identifying this RaftKV node.")
}

var raftkvCmd = &cobra.Command{
	Use:   "raftkv",
	Short: "Runs raftkv.",
	Long:  `Runs a raftkv node process.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var startRaftKVCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts raftkv.",
	Long:  `Starts a raftkv node process.`,
	Run: func(cmd *cobra.Command, args []string) {
		writePidFile("." + raftkvConfig.NodeID + raftkvPidFile)
		err := raftkv.New(&raftkvConfig).Start()

		if err != nil {
			panic(err)
		}
	},
}

var stopRaftKVCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops raftkv.",
	Long:  `Stops a raftkv node process.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopServer("." + raftkvConfig.NodeID + raftkvPidFile)
	},
}

var restartRaftKVCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts raftkv.",
	Long:  `Restarts a raftkv node process.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopRaftKVCmd.Run(cmd, args)
		startRaftKVCmd.Run(cmd, args)
	},
}
