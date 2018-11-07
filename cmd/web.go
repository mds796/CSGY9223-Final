package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

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
}

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Runs the web server for our Twitter clone.",
	Long:  `Runs a web server process to serve the static and dynamic assets for our twitter clone.`,
	Run: func(cmd *cobra.Command, args []string) {
		writePidFile()
		startServer()
	},
}

var startWebCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the web server for our Twitter clone.",
	Long:  `Starts a web server process to serve the static and dynamic assets for our twitter clone.`,
	Run: func(cmd *cobra.Command, args []string) {
		writePidFile()
		startServer()
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

	process, err := os.FindProcess(pid)
	panicOnError(err)
	err = process.Kill()

	if err != nil {
		select {
		case <-time.After(5 * time.Second):
			process.Signal(syscall.SIGKILL)
		}
	}

	os.Remove(pidFile)
}

func readPid() int {
	bytes, err := ioutil.ReadFile(pidFile)
	panicOnError(err)

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

func startServer() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			if _, err := io.Copy(w, r.Body); err != nil {
				log.Printf("Encountered an error while echoing the body: %v\n.", err)
			}
		}
	})
	http.Handle("/", http.FileServer(http.Dir("static/build/default")))
	log.Printf("Now listening on port %v.\n", port)
	log.Fatal(http.ListenAndServe(host+":"+strconv.Itoa(int(port)), nil))
}

func writePidFile() {
	bytes := []byte(strconv.Itoa(os.Getpid()) + "\n")
	err := ioutil.WriteFile(pidFile, bytes, 400)
	if err != nil {
		panic(err)
	}
}
