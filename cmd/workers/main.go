package main

import (
	log "github.com/sirupsen/logrus"
)

const (
	NoOfWorkers       = 100
	NoOfJobsPerWorker = 100000
)

func main() {
	ch := make(chan string)
	Wg.Add(NoOfWorkers)
	go CloseWorkersChannel(ch)

	for i := 0; i < NoOfWorkers; i++ {
		go worker(i, NoOfJobsPerWorker, ch)
	}
	for message := range ch {
		log.Println(message)
	}
}
