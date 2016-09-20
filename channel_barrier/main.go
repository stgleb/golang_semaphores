package main

import (
	"fmt"
)

/*
	Simple program illustrates stoping producing goroutines by
	setting chanel to nil, that blocks all goroutines writing to
	that channel and then unblocking this goroutine by restoring
	channel value.
*/

const N = 100
const DELTA = 10

func Producer0(ch chan<- int) {
	for i := 0; i < N; i++ {
		ch <- 0
	}
}

func Producer1(ch chan<- int) {
	for i := 0; i < N; i++ {
		ch <- 1
	}
}

func Consumer(zeroChan, oneChan chan int) {
	var zeroCount int
	var oneCount int
	var tmp chan int

	for i := 0; i < N; i++ {
		select {
		case <-zeroChan:
			zeroCount++
			fmt.Printf("zero: %d one: %d\n", zeroCount, oneCount)

			if oneChan == nil {
				oneChan = tmp
				fmt.Println("Restore one chan")
			}

			if zeroCount > oneCount+DELTA {
				tmp = zeroChan
				zeroChan = nil
				fmt.Println("Block zero chan")
			}
		case <-oneChan:
			oneCount++
			fmt.Printf("zero: %d one: %d\n", zeroCount, oneCount)

			if zeroChan == nil {
				zeroChan = tmp
				fmt.Println("Restore zero chan")
			}

			if oneCount > zeroCount+DELTA {
				tmp = oneChan
				oneChan = nil
				fmt.Println("Block one chan")
			}
		}
	}
}

func main() {
	zeroChan := make(chan int)
	oneChan := make(chan int)
	go Producer0(zeroChan)
	go Producer1(oneChan)
	Consumer(zeroChan, oneChan)
}
