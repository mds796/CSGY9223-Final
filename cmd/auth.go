package cmd

import (
	"github.com/mds796/CSGY9223-Final/auth"
	"github.com/spf13/cobra"
)

var authConfig auth.Config
var authPidFile string

func init() {
	rootCmd.AddCommand(authCmd)

	authCmd.PersistentFlags().StringVarP(&authPidFile, "pidFile", "p", ".auth.pid", "The name of the process ID file.")
	authCmd.AddCommand(startAuthCmd)
	authCmd.AddCommand(stopAuthCmd)
	authCmd.AddCommand(restartAuthCmd)

	authSetStartArgs(startAuthCmd)
	authSetStartArgs(restartAuthCmd)
}

func authSetStartArgs(command *cobra.Command) {
	command.Flags().StringVarP(&authConfig.Host, "host", "H", "localhost", "The host interface to listen on.")
	command.Flags().Uint16VarP(&authConfig.Port, "port", "P", 8089, "The TCP port to listen on.")

	command.Flags().StringVar(&authConfig.UserHost, "userHost", "localhost", "The hostname user service listens on.")
	command.Flags().Uint16Var(&authConfig.UserPort, "userPort", 8081, "The TCP port user service listens on.")

	command.Flags().StringSliceVar(&authConfig.StorageHosts, "storageHosts", []string{"localhost:7000", "localhost:7010", "localhost:7020"}, "The hostnames and TCP ports storage nodes listen on.")

}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Runs the auth server for our Twitter clone.",
	Long:  `Runs a auth server process.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var startAuthCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the auth server for our Twitter clone.",
	Long:  `Starts a auth server process.`,
	Run: func(cmd *cobra.Command, args []string) {
		writePidFile(authPidFile)
		err := auth.New(&authConfig).Start()

		if err != nil {
			panic(err)
		}
	},
}

var stopAuthCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the auth server for our Twitter clone.",
	Long:  `Stops a auth server process.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopServer(authPidFile)
	},
}

var restartAuthCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts the auth server for our Twitter clone.",
	Long:  `Restarts a auth server process.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopAuthCmd.Run(cmd, args)
		startAuthCmd.Run(cmd, args)
	},
}
