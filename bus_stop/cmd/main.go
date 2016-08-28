package main

import (
	"../../bus_stop"
	"log"
	"os"
	"time"
)

/*
    Riders come to a bus stop and wait for a bus. When the bus arrives, all the waiting
    riders invoke boardBus, but anyone who arrives while the bus is boarding has to wait
    for the next bus.The capacity of the bus is 50 people;
    if there are more than 50 people waiting, some will have to wait for the next bus.
    When all the waiting riders have boarded, the bus can invoke depart. If the bus arrives
    when there are no riders, it should depart immediately.
*/

var logger = log.New(os.Stdout,
	"MAIN: ",
	log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	logger.Println("Begin")
	p1 := bus_stop.NewPassenger("Sean", "Pen")
	p2 := bus_stop.NewPassenger("Will", "Smith")
	p3 := bus_stop.NewPassenger("Bruce", "Willis")
	passengers := make([]bus_stop.Passenger, 0, 3)
	passengers = append(passengers, p1)
	passengers = append(passengers, p2)
	passengers = append(passengers, p3)

	NewYorkStation := bus_stop.NewBusStop("New York")
	SanFranciscoStation := bus_stop.NewBusStop("San-Francisco")

	bus := bus_stop.NewBus(NewYorkStation, SanFranciscoStation)

	go NewYorkStation.Run()
	go SanFranciscoStation.Run()

	// Put all passengers on New York
	for _, p := range passengers {
		logger.Printf("Board passenger %s to station %s", p, NewYorkStation)
		NewYorkStation.PassengerArriveChannel <- p
	}

	// Run bus and depart it from San Francisco station
	go bus.Run()
	bus.BusDeparture <- SanFranciscoStation

	select {
	case <-time.After(time.Second * 1):
	}

}
