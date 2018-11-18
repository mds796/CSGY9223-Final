package cmd

import (
	"github.com/mds796/CSGY9223-Final/feed"
	"github.com/spf13/cobra"
)

var feedConfig feed.Config
var feedPidFile string

func init() {
	rootCmd.AddCommand(feedCmd)

	feedCmd.PersistentFlags().StringVarP(&feedPidFile, "pidFile", "p", ".feed.pid", "The name of the process ID file.")
	feedCmd.AddCommand(startFeedCmd)
	feedCmd.AddCommand(stopFeedCmd)
	feedCmd.AddCommand(restartFeedCmd)

	feedSetStartArgs(startFeedCmd)
	feedSetStartArgs(restartFeedCmd)
}

func feedSetStartArgs(command *cobra.Command) {
	command.Flags().StringVarP(&feedConfig.Host, "host", "H", "localhost", "The host interface to listen on.")
	command.Flags().Uint16VarP(&feedConfig.Port, "port", "P", 8085, "The TCP port to listen on.")

	command.Flags().StringVar(&feedConfig.UserHost, "userHost", "localhost", "The hostname user service listens on.")
	command.Flags().Uint16Var(&feedConfig.UserPort, "userPort", 8081, "The TCP port user service listens on.")

	command.Flags().StringVar(&feedConfig.PostHost, "postHost", "localhost", "The hostname post service listens on.")
	command.Flags().Uint16Var(&feedConfig.PostPort, "postPort", 8083, "The TCP port post service listens on.")

	command.Flags().StringVar(&feedConfig.FollowHost, "followHost", "localhost", "The hostname follow service listens on.")
	command.Flags().Uint16Var(&feedConfig.FollowPort, "followPort", 8084, "The TCP port follow service listens on.")
}

var feedCmd = &cobra.Command{
	Use:   "feed",
	Short: "Runs the feed server for our Twitter clone.",
	Long:  `Runs a feed server process.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var startFeedCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the feed server for our Twitter clone.",
	Long:  `Starts a feed server process.`,
	Run: func(cmd *cobra.Command, args []string) {
		writePidFile(feedPidFile)
		err := feed.New(&feedConfig).Start()

		if err != nil {
			panic(err)
		}
	},
}

var stopFeedCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the feed server for our Twitter clone.",
	Long:  `Stops a feed server process.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopServer(feedPidFile)
	},
}

var restartFeedCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts the feed server for our Twitter clone.",
	Long:  `Restarts a feed server process.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopFeedCmd.Run(cmd, args)
		startFeedCmd.Run(cmd, args)
	},
}
