package cmd

import (
	"github.com/mds796/CSGY9223-Final/user"
	"github.com/spf13/cobra"
)

var userConfig user.Config
var userPidFile string

func init() {
	rootCmd.AddCommand(userCmd)

	userCmd.PersistentFlags().StringVarP(&userPidFile, "pidFile", "p", ".user.pid", "The name of the process ID file.")
	userCmd.AddCommand(startUserCmd)
	userCmd.AddCommand(stopUserCmd)
	userCmd.AddCommand(restartUserCmd)

	userSetStartArgs(startUserCmd)
	userSetStartArgs(restartUserCmd)
}

func userSetStartArgs(command *cobra.Command) {
	command.Flags().StringVarP(&userConfig.Host, "host", "H", "localhost", "The host interface to listen on.")
	command.Flags().Uint16VarP(&userConfig.Port, "port", "P", 8081, "The TCP port to listen on.")
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Runs the user server for our Twitter clone.",
	Long:  `Runs a user server process.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var startUserCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the user server for our Twitter clone.",
	Long:  `Starts a user server process.`,
	Run: func(cmd *cobra.Command, args []string) {
		writePidFile(userPidFile)
		err := user.New(&userConfig).Start()

		if err != nil {
			panic(err)
		}
	},
}

var stopUserCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the user server for our Twitter clone.",
	Long:  `Stops a user server process.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopServer(userPidFile)
	},
}

var restartUserCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts the user server for our Twitter clone.",
	Long:  `Restarts a user server process.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopUserCmd.Run(cmd, args)
		startUserCmd.Run(cmd, args)
	},
}
