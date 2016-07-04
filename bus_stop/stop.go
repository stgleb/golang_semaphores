package bus_stop

import (
	"fmt"
	"sync"
)

const BUS_CAPACITY = 2

// Bus stop implements Station interface.
type BusStop struct {
	BusArriveChannel       chan Bus
	PassengerArriveChannel chan Passenger
	StationName            string
	Lock                   sync.Mutex
	capacity               int
	passengerBuffer        []Passenger
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}

}

func NewBusStop(stationName string) BusStop {
	return BusStop{
		BusArriveChannel:       make(chan Bus),
		PassengerArriveChannel: make(chan Passenger, BUS_CAPACITY),
		StationName:            stationName,
		capacity:               0,
		passengerBuffer:        make([]Passenger, 0, BUS_CAPACITY),
		Lock:                   sync.Mutex{},
	}
}

func (stop BusStop) String() string {
	return fmt.Sprintf("< %s >", stop.StationName)
}

// Called by Bus on arrival.
func (stop *BusStop) BusArrives(bus Bus) {
	stop.Lock.Lock()
	defer stop.Lock.Unlock()
	cnt := min(len(stop.passengerBuffer), BUS_CAPACITY)

	for i := 0; i < cnt; i++ {
		bus.PassengerBoard <- stop.passengerBuffer[i]
	}
	// Truncate boarded passengers from slice
	stop.passengerBuffer = stop.passengerBuffer[cnt:]
	// Signal to bus about departure
	bus.BusDeparture <- *stop
}

// Called to board passenger.
func (stop *BusStop) boardPassenger(passenger Passenger) {
	stop.Lock.Lock()
	defer stop.Lock.Unlock()

	stopLogger.Printf("Passenger %s has boarder on station %s ", passenger, stop)
	stop.passengerBuffer = append(stop.passengerBuffer, passenger)
}

// Main event loop for Bus Stop.
func (stop *BusStop) Run() {
	for {
		select {
		case bus := <-stop.BusArriveChannel:
			stop.BusArrives(bus)
		case passenger := <-stop.PassengerArriveChannel:
			stopLogger.Printf("Passenger %s arrives on station %s", passenger, stop)
			stop.boardPassenger(passenger)
		}
	}
}
