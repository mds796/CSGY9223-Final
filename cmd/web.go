package cmd

import (
	"github.com/mds796/CSGY9223-Final/web"
	"github.com/spf13/cobra"
)

var webConfig web.Config
var webPidFile string

func init() {
	rootCmd.AddCommand(webCmd)

	webCmd.PersistentFlags().StringVarP(&webPidFile, "pidFile", "p", ".web.pid", "The name of the process ID file.")
	webCmd.AddCommand(startWebCmd)
	webCmd.AddCommand(stopWebCmd)
	webCmd.AddCommand(restartWebCmd)

	webSetStartArgs(startWebCmd)
	webSetStartArgs(restartWebCmd)
}

func webSetStartArgs(command *cobra.Command) {
	command.Flags().StringVarP(&webConfig.Host, "host", "H", "localhost", "The host interface to listen on.")
	command.Flags().Uint16VarP(&webConfig.Port, "port", "P", 8080, "The TCP port to listen on.")

	command.Flags().StringVar(&webConfig.UserHost, "userHost", "localhost", "The hostname user service listens on.")
	command.Flags().Uint16Var(&webConfig.UserPort, "userPort", 8081, "The TCP port user service listens on.")

	command.Flags().StringVar(&webConfig.AuthHost, "authHost", "localhost", "The hostname auth service listens on.")
	command.Flags().Uint16Var(&webConfig.AuthPort, "authPort", 8082, "The TCP port auth service listens on.")

	command.Flags().StringVar(&webConfig.PostHost, "postHost", "localhost", "The hostname post service listens on.")
	command.Flags().Uint16Var(&webConfig.PostPort, "postPort", 8083, "The TCP port post service listens on.")

	command.Flags().StringVar(&webConfig.FollowHost, "followHost", "localhost", "The hostname follow service listens on.")
	command.Flags().Uint16Var(&webConfig.FollowPort, "followPort", 8084, "The TCP port follow service listens on.")

	command.Flags().StringVar(&webConfig.FeedHost, "feedHost", "localhost", "The hostname feed service listens on.")
	command.Flags().Uint16Var(&webConfig.FeedPort, "feedPort", 8085, "The TCP port feed service listens on.")

	command.Flags().StringVarP(&webConfig.StaticPath, "staticPath", "S", "static/build/default", "The file path of the static assets directory.")
	command.Flags().StringVarP(&webConfig.StaticUrl, "staticUrl", "U", "http://localhost:8000", "The URL path of the static assets server.")
}

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Runs the web server for our Twitter clone.",
	Long:  `Runs a web server process to serve the static and dynamic assets for our twitter clone.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var startWebCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the web server for our Twitter clone.",
	Long:  `Starts a web server process to serve the static and dynamic assets for our twitter clone.`,
	Run: func(cmd *cobra.Command, args []string) {
		writePidFile(webPidFile)
		web.New(&webConfig).Start()
	},
}

var stopWebCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the web server for our Twitter clone.",
	Long:  `Stops a web server process to serve the static and dynamic assets for our twitter clone.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopServer(webPidFile)
	},
}

var restartWebCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts the web server for our Twitter clone.",
	Long:  `Restarts a web server process to serve the static and dynamic assets for our twitter clone.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopWebCmd.Run(cmd, args)
		startWebCmd.Run(cmd, args)
	},
}
