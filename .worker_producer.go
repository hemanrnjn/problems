package main

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)

type job struct {
	id int
	time int
}

func main() {
	var wg sync.WaitGroup
	jobChan := make(chan job)
	for i := 0; i < 4; i++ {
		go worker(jobChan, &wg)
		wg.Add(1)
	}
	startTime := time.Now()
	go assigner(jobChan, &wg)
	wg.Wait()
	// endTime := time.Now()
	fmt.Println("Total time taken: ", time.Since(startTime))
}

func assigner(jobChan chan job, wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		job := job{
			id: i,
			time: rand.Intn(5),
		}
		jobChan <- job
	}
	close(jobChan)
	wg.Done()
}

func worker(jobChan chan job, wg *sync.WaitGroup) {
	for j := range jobChan {
		fmt.Println("Working for job id: ", j.id)
		time.Sleep(time.Second * time.Duration(j.time))
		fmt.Println("Finished job: ", j.id)
	}
	wg.Done()
}
