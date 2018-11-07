package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mds796/CSGY9223-Final/web"
	"github.com/spf13/cobra"
)

var staticPath string
var pidFile string
var host string
var port uint16

func init() {
	rootCmd.AddCommand(webCmd)

	webCmd.PersistentFlags().StringVarP(&pidFile, "pidFile", "p", ".web.pid", "The name of the process ID file.")
	webCmd.AddCommand(startWebCmd)
	webCmd.AddCommand(stopWebCmd)
	webCmd.AddCommand(restartWebCmd)

	setStartArgs(startWebCmd)
	setStartArgs(restartWebCmd)
}

func setStartArgs(command *cobra.Command) {
	command.Flags().StringVarP(&host, "host", "H", "localhost", "The host interface to listen on.")
	command.Flags().Uint16VarP(&port, "port", "P", 8080, "The TCP port to listen on.")
	command.Flags().StringVarP(&staticPath, "staticPath", "S", "static/build/default", "The file path of the static assets directory.")
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
		writePidFile()
		web.Start(host, port, staticPath)
	},
}

var stopWebCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the web server for our Twitter clone.",
	Long:  `Stops a web server process to serve the static and dynamic assets for our twitter clone.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopServer()
	},
}

func stopServer() {
	pid := readPid()
	if pid == -1 {
		return
	}

	process, err := os.FindProcess(pid)
	panicOnError(err)

	err = process.Kill()
	if err != nil {
		log.Print(err)
	}

	err = process.Release()
	panicOnError(err)

	os.Remove(pidFile)
}

func readPid() int {
	bytes, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return -1
	}

	pid, err := strconv.Atoi(strings.TrimSpace(string(bytes)))
	panicOnError(err)

	return pid
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

var restartWebCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts the web server for our Twitter clone.",
	Long:  `Restarts a web server process to serve the static and dynamic assets for our twitter clone.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopWebCmd.Run(stopWebCmd, args)
		startWebCmd.Run(cmd, args)
	},
}

func writePidFile() {
	bytes := []byte(strconv.Itoa(os.Getpid()) + "\n")
	err := ioutil.WriteFile(pidFile, bytes, 400)

	if err != nil {
		panic(err)
	}
}
