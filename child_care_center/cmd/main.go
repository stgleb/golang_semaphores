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

func RunChildren(center child_care_center.ChildCareCenter,
	children []child_care_center.Child) {
	for i := range children {

		go func(child child_care_center.Child) {
			center.ChildIn <- child
			center.ChildOut <- child
		}(children[i])
	}

	wg.Done()
}

func RunAdults(center child_care_center.ChildCareCenter,
	adults []child_care_center.Adult) {
	for i := range adults {

		go func(adult child_care_center.Adult) {
			center.AdultIn <- adult
			center.AdultOut <- adult
		}(adults[i])
	}

	wg.Done()
}

func main() {
	center := child_care_center.NewChildCenter()
	children := CreateChildren(30)
	adults := CreateAdults(10)
	wg.Add(2)
	go center.Run()
	go RunAdults(center, adults)
	go RunChildren(center, children)
	wg.Wait()
}
