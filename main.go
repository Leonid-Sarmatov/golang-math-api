package main

import (
	"log"
	"os"
	"os/signal"
	"server_3/pkg"
)

func main() {
	api := pkg.API{
		AppName: "Math service",
		Port:    "8082",
		Executors: []pkg.Executor{
			pkg.NewMVEMathExpressionExecutor(),
		},
	}

	api.APIRun()

	osSignalsChan := make(chan os.Signal, 1)
	signal.Notify(osSignalsChan, os.Interrupt)

	<-osSignalsChan
	log.Println("Math service was stoped!")
}
