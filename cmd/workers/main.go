package main

import (
	log "github.com/sirupsen/logrus"
)

const (
	NoOfWorkers       = 1000
	NoOfJobsPerWorker = 1000000
)

func main() {
	ch := make(chan string)
	go CloseWorkersChannel(ch)
	for i := 0; i < NoOfWorkers; i++ {
		Wg.Add(1)
		go worker(i, NoOfJobsPerWorker, ch)
	}
	for message := range ch {
		log.Println(message)
	}
}
