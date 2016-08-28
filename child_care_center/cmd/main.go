package main

import (
	"../../child_care_center"
	"github.com/dustinkirkland/golang-petname"
	"time"
)

/*
State licensing rules require a child-care center to have no more than three
infants present for each adult
*/


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
	go center.Run()


	for i := range children {
		go func(child child_care_center.Child) {
			center.ChildIn <- child
			time.Sleep(time.Millisecond * 1)
			center.ChildOut <- child
		}(children[i])
	}

	for i := range adults {
		go func(adult child_care_center.Adult) {
			center.AdultIn <- adult
			time.Sleep(time.Millisecond * 1)
			center.AdultOut <- adult
		}(adults[i])
	}

	<- time.After(time.Millisecond * 100)
}
