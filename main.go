package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var build = "develop"

func main() {
	log.Println("Start service", build)
	defer log.Panicln("Service ended")

	schutdown := make(chan os.Signal, 1)
	signal.Notify(schutdown, syscall.SIGINT, syscall.SIGTERM)
	<-schutdown

	log.Println("Stopping service")
}
