package main

import (
	"../../child_care_center"
	"github.com/dustinkirkland/golang-petname"
	"sync"
)

var wg sync.WaitGroup

func CreateChildren(count int) []child_care_center.Child {
	children := make([]child_care_center.Child, 0, count)

	for i := 0; i < count; i++ {
		children = append(children, child_care_center.Child{
			Name: petname.Generate(1, ""),
		})
	}

	return children
}

func CreateAdults(count int) []child_care_center.Adult {
	adults := make([]child_care_center.Adult, 0, count)

	for i := 0; i < count; i++ {
		adults = append(adults, child_care_center.Adult{
			Name: petname.Generate(1, ""),
		})
	}

	return adults
}

func main() {
	center := child_care_center.NewChildCenter()
	children := CreateChildren(30)
	adults := CreateAdults(10)
	wg.Add(len(adults))
	go center.Run()

	for i := range adults {
		go func(adult child_care_center.Adult) {
			center.AdultIn <- adult
			center.AdultOut <- adult
			wg.Done()
		}(adults[i])
	}

	for i := range children {

		go func(child child_care_center.Child) {
			center.ChildIn <- child
			center.ChildOut <- child
			wg.Done()
		}(children[i])
	}

	wg.Wait()
}
