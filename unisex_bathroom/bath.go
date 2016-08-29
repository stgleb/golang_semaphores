package unisex_bathroom

import (
	"fmt"
	"log"
)

type Bath struct {
	MenIn    chan Person
	MenOut   chan Person
	WomenIn  chan Person
	WomenOut chan Person
	Close    chan struct{}

	maleCount   int
	femaleCount int
	capacity    int
}

func NewBath(capacity int) Bath {
	return Bath{
		MenIn:    make(chan Person),
		MenOut:   make(chan Person),
		WomenIn:  make(chan Person),
		WomenOut: make(chan Person),
		Close:    make(chan struct{}),

		maleCount:   0,
		femaleCount: 0,
		capacity:    capacity,
	}
}

func (bath Bath) String() string {
	return fmt.Sprintf("[FemaleCount: %d, MaleCount: %d]",
		bath.femaleCount,
		bath.maleCount)
}

// Main event loop
func (bath Bath) Run() {
	var wIn chan Person
	var mIn chan Person

	for {
		select {
		case p := <-bath.MenIn:
			bath.maleCount++
			log.Printf("Enter: %s Bath: %s", p, bath)
			// Close bath for women
			if bath.WomenIn != nil {
				wIn = bath.WomenIn
				bath.WomenIn = nil
			}
			// Close bath for another men
			if bath.maleCount == bath.capacity {
				mIn = bath.MenIn
				bath.MenIn = nil
			}
		case p := <-bath.MenOut:
			bath.maleCount--
			log.Printf("Exit: %s Bath: %s", p, bath)
			// Open bath for women
			if bath.maleCount == 0 {
				bath.WomenIn = wIn
			}
			// Open bath for another men
			if bath.maleCount == bath.capacity-1 {
				bath.MenIn = mIn
			}
		case p := <-bath.WomenIn:
			bath.femaleCount++
			log.Printf("Enter: %s Bath: %s", p, bath)
			// Close bath for men
			if bath.MenIn != nil {
				mIn = bath.MenIn
				bath.MenIn = nil
			}
			// Close bath for another women
			if bath.femaleCount == bath.capacity {
				wIn = bath.WomenIn
				bath.WomenIn = nil
			}
		case p := <-bath.WomenOut:
			bath.femaleCount--
			log.Printf("Exit: %s Bath: %s", p, bath)

			// Open bath for men
			if bath.femaleCount == 0 {
				bath.MenIn = mIn
			}
			// Open bath for another women
			if bath.femaleCount == bath.capacity-1 {
				bath.WomenIn = wIn
			}
		case <-bath.Close:
			return
		}
	}
}
