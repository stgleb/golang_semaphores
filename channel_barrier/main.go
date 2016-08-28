package main

import (
	"fmt"
	"time"
)

/*
	Simple program illustrates stoping producing goroutines by
	setting chanel to nil, that blocks all gorout	ines and then
	restoring chanel value.
*/


func foo(id int, ch chan int) {
	fmt.Printf("Hello, i am %d\n",id)
	ch <- id
}

func main() {
	ch := make(chan int)
	var ch2 chan int

	for i:= 0; i < 10;i++ {
		go foo(i, ch)
	}

	for i := 0;i < 5; i++ {
		<- ch
	}
	ch2 = ch
	ch = nil
	fmt.Println("Stop receiving and sleep")
	time.Sleep(1)
	ch = ch2

	for i := 0;i < 5; i++ {
		<- ch
	}
}
