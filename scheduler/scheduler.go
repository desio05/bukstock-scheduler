package scheduler

import (
	"log"
	"time"
)

// Schedule ticker updating
func Schedule() {
	scheduleTimer()
}

type timerMessage struct {
	config *Config
}

func scheduleTimer() {
	var config Config
	for {
		config.Parse()
		execute(config.Tickers)
		time.Sleep(time.Duration(config.TimeoutMillis) * time.Millisecond)
	}
}

func execute(tickers []string) {
	log.Print("execute.in")
	log.Print("execute.out")
}
