package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var client = &http.Client{}
var Wg sync.WaitGroup

func worker(id int, itr int, ch chan<- string) {
	defer Wg.Done()
	for i := 0; i < itr; i++ {
		randomWaitTime := time.Duration(rand.Int63n(int64(time.Second)))

		time.Sleep(randomWaitTime)
		req, err := http.NewRequest("GET", "http://localhost:8080/health", nil)
		if err != nil {
			ch <- "Error creating request"
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			ch <- "Error sending request"
			return
		}
		if resp.StatusCode != 200 {
			ch <- "Error sending request"
			return
		}
		ch <- fmt.Sprintf("Id %d: is on %d iteration resp %d", id, i, resp.Body)
	}
}

func CloseWorkersChannel(ch chan string) {
	Wg.Wait()
	close(ch)
}
