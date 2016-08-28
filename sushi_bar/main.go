package main

import (
	"fmt"
	"sync"
	"github.com/dustinkirkland/golang-petname"
	"log"
)


/*
	Imagine a sushi bar with 5 seats. If you arrive while there is an empty seat,
	you can take a seat immediately. But if you arrive when all 5 seats are full,
	that means that all of them are dining together,
	and you will have to wait for the entire party to leave before you sit down.
*/

type Client struct {
	Name string
}

func (client Client) String() string {
	return fmt.Sprintf("Client %s", client.Name)
}

type Bar struct {
	In chan Client
	Out chan Client
	count int
	capacity int
}

var wg sync.WaitGroup

func NewBar(capacity int) Bar {
	return Bar{
		In: make(chan Client),
		Out: make(chan Client),
		count: 0,
		capacity: capacity,
	}
}

func (bar *Bar) Run() {
	var in chan Client
	log.Println("Start running sushi bar")

	for {
		select {
		case client := <- bar.In:
			log.Printf("Enter %v", client)
			bar.count++
			// Lock enter to sushi bar
			if bar.count == bar.capacity {
				in = bar.In
			}
		case client := <- bar.Out:
			log.Printf("Exit %v", client)
			// Unlock enter to sushi bar
			if bar.count == bar.capacity {
				bar.In = in
			}

			bar.count--
		}
	}
}

func GenerateClients(count int) []Client {
	clients := make([]Client, 0, count)

	for i := 0;i < count; i++ {
		clients = append(clients, Client{
			Name: petname.Generate(1, ""),
		})
	}

	return clients
}

func RunClient(client Client, bar Bar) {
	bar.In <- client
	bar.Out <- client
	wg.Done()
}

func main() {
	bar := NewBar(5)
	clients := GenerateClients(10)
	go bar.Run()
	wg.Add(len(clients))

	for index := range clients {
		client := clients[index]
		go RunClient(client, bar)
	}

	wg.Wait()
}