package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/logger"
	"github.com/jberny/monitoring-demo-api/api/controller"
)

func handleSignals() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		<- sigc
		logger.Infoln("\n\nTime to go!")
		os.Exit(0)
	}()
}

func usage() {
    flag.PrintDefaults()
    os.Exit(2)
}

var verbose = flag.Bool("v", false, "print info level logs to stdout")

func init() {
	flag.Usage = usage
    
    flag.Parse()
}

func main() {
	handleSignals()
	logger.Init("StdoutLogger", *verbose, false, os.Stdout)
	logger.Infoln("Starting http server on port 8080")
	router := controller.NewRouter()
	err := http.ListenAndServe(":8080", router)
	if err != nil { 
		logger.Fatalln(err)
	}
}

