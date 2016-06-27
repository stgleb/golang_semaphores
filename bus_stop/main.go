package main

import (
	"log"
	"sync"
)

/*
Riders come to a bus stop and wait for a bus. When the bus arrives, all the waiting riders invoke boardBus,
but anyone who arrives while the bus is boarding has to wait for the next bus. The capacity of the bus is 50 people;
if there are more than 50 people waiting, some will have to wait for the next bus.

When all the waiting riders have boarded, the bus can invoke depart. If the bus arrives when there are no riders,
it should depart immediately.
*/
const BUS_CAPACITY = 50

// General interface for stop
type Station interface {
	BusArrives()
	PassengerArrive(passenger Passenger)
}

// Bus stop implements Station interface.
type BusStop struct {
	busArrive       chan struct{}
	passengerArrive chan Passenger
	capacity        int
	passengerBuffer []Passenger
	Lock            sync.Mutex
}

type Passenger struct{}

func min(a, b int) {
	if a < b {
		return a
	} else {
		return b
	}
}

func NewBusStop() BusStop {
	return BusStop{
		busArrive:       make(chan struct{}),
		passengerArrive: make(chan Passenger, BUS_CAPACITY),
		capacity:        0,
		passengerBuffer: make([]Passenger, BUS_CAPACITY),
		Lock:            sync.Mutex{},
	}
}

// Called by Bus on arrival.
// TODO: add Bus type to argument, add passengerBuffer truncation.
func (stop *BusStop) BusArrives() {
	stop.Lock.Lock()
	defer stop.Lock.Unlock()

	for i := 0; i < min(len(stop.passengerBuffer), BUS_CAPACITY); i++ {
		log.Printf("Passenger %v has been boarded", stop.passengerBuffer[i])
	}
}

// Called to board passenger.
func (stop *BusStop) boardPassenger(passenger Passenger) {
	stop.Lock.Lock()
	defer stop.Lock.Unlock()

	log.Printf("Passenger %v has boarder", passenger)
	stop.passengerBuffer = append(stop.passengerBuffer, passenger)
}

// Called by Passenger on arrival.
func (stop *BusStop) PassengerArrive(passenger Passenger) {
	log.Printf("Passenger %v has arrived", passenger)
	stop.passengerArrive <- passenger
}

// Main event loop for Bus Stop.
func (stop *BusStop) Run() {
	for {
		select {
		case <-stop.busArrive:
			stop.BusArrives()
		case passenger := <-stop.passengerArrive:
			log.Printf("Passenger %v arrives", passenger)
			stop.boardPassenger(passenger)
		}
	}
}

func main() {
	log.Println("Begin")
}
