package cmd

import (
	"github.com/mds796/CSGY9223-Final/post"
	"github.com/spf13/cobra"
)

var postConfig post.Config
var postPidFile string

func init() {
	rootCmd.AddCommand(postCmd)

	postCmd.PersistentFlags().StringVarP(&postPidFile, "pidFile", "p", ".post.pid", "The name of the process ID file.")
	postCmd.AddCommand(startPostCmd)
	postCmd.AddCommand(stopPostCmd)
	postCmd.AddCommand(restartPostCmd)

	postSetStartArgs(startPostCmd)
	postSetStartArgs(restartPostCmd)
}

func postSetStartArgs(command *cobra.Command) {
	command.Flags().StringVarP(&postConfig.Host, "host", "H", "localhost", "The host interface to listen on.")
	command.Flags().Uint16VarP(&postConfig.Port, "port", "P", 8083, "The TCP port to listen on.")
	command.Flags().StringSliceVar(&postConfig.StorageHosts, "storageHosts", []string{"localhost:7000", "localhost:7010", "localhost:7020"}, "The hostnames and TCP ports storage nodes listen on.")
}

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Runs the post server for our Twitter clone.",
	Long:  `Runs a post server process.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var startPostCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the post server for our Twitter clone.",
	Long:  `Starts a post server process.`,
	Run: func(cmd *cobra.Command, args []string) {
		writePidFile(postPidFile)
		err := post.New(&postConfig).Start()

		if err != nil {
			panic(err)
		}
	},
}

var stopPostCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the post server for our Twitter clone.",
	Long:  `Stops a post server process.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopServer(postPidFile)
	},
}

var restartPostCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts the post server for our Twitter clone.",
	Long:  `Restarts a post server process.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopPostCmd.Run(cmd, args)
		startPostCmd.Run(cmd, args)
	},
}
