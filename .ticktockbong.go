package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var totalTime int
	fmt.Print("Enter time in seconds:")
	fmt.Scanln(&totalTime)
	var wg sync.WaitGroup
	timeCounter := 1
	secChan := make(chan string)
	minChan := make(chan string)
	hrChan := make(chan string)
	doneChan := make(chan bool)

	go generator(secChan, minChan, hrChan, doneChan, timeCounter, totalTime, &wg)
	go printer(secChan, minChan, hrChan, doneChan, &wg)
	wg.Add(2)
	wg.Wait()
}

func generator(secChan, minChan, hrChan chan string, doneChan chan bool, timeCounter, totalTime int, wg *sync.WaitGroup) {
outer:
	for {
		switch {
		case timeCounter == totalTime:
			break outer
		case timeCounter%60 == 0:
			minChan <- "tock"
		case timeCounter%3600 == 0:
			hrChan <- "bong"
		default:
			secChan <- "tick"
		}
		timeCounter++
		time.Sleep(time.Second * 1)
	}
	doneChan <- true
	wg.Done()
}

func printer(secChan, minChan, hrChan chan string, doneChan chan bool, wg *sync.WaitGroup) {
outer:
	for {
		select {
		case s := <-secChan:
			fmt.Println(s)
		case m := <-minChan:
			fmt.Println(m)
		case h := <-hrChan:
			fmt.Println(h)
		case <-doneChan:
			break outer
		default:
			continue
		}
	}
	wg.Done()
}
