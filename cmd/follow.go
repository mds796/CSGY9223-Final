package cmd

import (
	"github.com/mds796/CSGY9223-Final/follow"
	"github.com/spf13/cobra"
)

var followConfig follow.Config
var followPidFile string

func init() {
	rootCmd.AddCommand(followCmd)

	followCmd.PersistentFlags().StringVarP(&followPidFile, "pidFile", "p", ".follow.pid", "The name of the process ID file.")
	followCmd.AddCommand(startFollowCmd)
	followCmd.AddCommand(stopFollowCmd)
	followCmd.AddCommand(restartFollowCmd)

	followSetStartArgs(startFollowCmd)
	followSetStartArgs(restartFollowCmd)
}

func followSetStartArgs(command *cobra.Command) {
	command.Flags().StringVarP(&followConfig.Host, "host", "H", "localhost", "The host interface to listen on.")
	command.Flags().Uint16VarP(&followConfig.Port, "port", "P", 8084, "The TCP port to listen on.")

	command.Flags().StringVar(&followConfig.UserHost, "userHost", "localhost", "The hostname user service listens on.")
	command.Flags().Uint16Var(&followConfig.UserPort, "userPort", 8081, "The TCP port user service listens on.")

	command.Flags().StringSliceVar(&followConfig.StorageHosts, "storageHosts", []string{"localhost:7000", "localhost:7010", "localhost:7020"}, "The hostnames and TCP ports storage nodes listen on.")
}

var followCmd = &cobra.Command{
	Use:   "follow",
	Short: "Runs the follow server for our Twitter clone.",
	Long:  `Runs a follow server process.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var startFollowCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the follow server for our Twitter clone.",
	Long:  `Starts a follow server process.`,
	Run: func(cmd *cobra.Command, args []string) {
		writePidFile(followPidFile)
		err := follow.New(&followConfig).Start()

		if err != nil {
			panic(err)
		}
	},
}

var stopFollowCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the follow server for our Twitter clone.",
	Long:  `Stops a follow server process.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopServer(followPidFile)
	},
}

var restartFollowCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts the follow server for our Twitter clone.",
	Long:  `Restarts a follow server process.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopFollowCmd.Run(cmd, args)
		startFollowCmd.Run(cmd, args)
	},
}
