package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	message string = "default message."
	sleepTime string = "1s"
)

func main() {
	for {
		// get hostname from os
		hostname, err := os.Hostname()
		if err != nil {
			log.Panicln("Error getting hostname.")
		}

		// print message
		if len(os.Getenv("MESSAGE")) != 0 {
			message = os.Getenv("MESSAGE")
		}
		fmt.Printf("hostname: %s - %s\n", hostname, message)

		// sleep time
		if len(os.Getenv("SLEEP_TIME")) != 0 {
			sleepTime = os.Getenv("SLEEP_TIME")
		}

		sleepTimeDuration, _ := time.ParseDuration(sleepTime)
		time.Sleep(sleepTimeDuration)
	}
}
