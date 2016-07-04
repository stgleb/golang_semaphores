package bus_stop

import (
	"time"
)

type Bus struct {
	PassengerBoard   chan Passenger
	BusDeparture     chan BusStop
	BusArrive        chan BusStop
	passengers       []Passenger
	busStops         []BusStop
	currentStopIndex int
}

func NewBus(stops ...BusStop) *Bus {
	return &Bus{
		PassengerBoard:   make(chan Passenger),
		BusDeparture:     make(chan BusStop),
		BusArrive:        make(chan BusStop),
		passengers:       make([]Passenger, 0, BUS_CAPACITY),
		busStops:         stops,
		currentStopIndex: 0,
	}
}

func (bus *Bus) Run() {
	for {
		select {
		case passenger := <-bus.PassengerBoard:
			busLogger.Printf("Passenger %s  has boarded on the Bus", passenger)
			bus.passengers = append(bus.passengers, passenger)
		case station := <-bus.BusDeparture:
			busLogger.Printf("Bus has been departed from station %s", station)
			bus.currentStopIndex = (bus.currentStopIndex + 1) % len(bus.busStops)
			nextStation := bus.busStops[bus.currentStopIndex]
			delay()

			// Drop passengers on next station
			for _, passenger := range bus.passengers {
				busLogger.Printf("Passenger %s has exited on station %s", passenger, nextStation)
				nextStation.PassengerArriveChannel <- passenger
			}
			// Notify next station about bus arrival.
			nextStation.BusArriveChannel <- *bus
		}
	}
}

func delay() {
	select {
	case <-time.After(time.Millisecond * 100):
	}
}
