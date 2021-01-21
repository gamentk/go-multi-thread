package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	data = make([]int, 0)
	m    = sync.Mutex{}
)

func sleepTime() int {
	min, max := 1, 4
	sleep := (rand.Intn(max - min) + min)

	return sleep
}

func producer(n int) {
	min := 1
	max := 20

	for {
		if len(data) < 100 {
			mydata := rand.Intn(max-min) + min
			m.Lock()
			data = append(data, mydata)
			m.Unlock()

			fmt.Printf("Producer %d => %d\n", n, mydata)
		} else {
			fmt.Printf("Buffer is full\n")
		}

		min := 1
		max := 4

		sleep := (rand.Intn(max-min) + min)
		time.Sleep(time.Duration(sleep) * time.Second)
	}
}

func consumer(n int) {
	for {
		if len(data) != 0 {
			mydata := data[0]
			m.Lock()
			data = data[1:]
			m.Unlock()

			fmt.Printf("Consumer %d => %d, Fibonacci = %d\n", n, mydata, fib(mydata))
		} else {
			fmt.Printf("No data\n")
		}

		min := 1
		max := 4

		sleep := (rand.Intn(max-min) + min)
		time.Sleep(time.Duration(sleep) * time.Second)
	}
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	go producer(1)
	go producer(2)
	go consumer(1)
	go consumer(2)
	go consumer(3)
	select {}
}
