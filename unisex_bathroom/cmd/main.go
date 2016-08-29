package main

import (
	"../../unisex_bathroom"
	"github.com/dustinkirkland/golang-petname"
	"math/rand"
	"sync"
	"time"
)

/*
   A friend of the author took a job at Xerox. She was working in a cubicle
   in the basement of a concrete monolith, and the nearest women’s bathroom
   was two floors up. She proposed to the Uberboss that they convert the men’s
   bathroom on her floor to a unisex bathroom, sort of like on Ally McBeal.

   The Uberboss agreed, provided that the following synchronization constraints can be maintained:

   • There cannot be men and women in the bathroom at the same time.
   • There should never be more than three employees squandering company time in the bathroom.
*/

const N = 20

func CreatePersons(count int) []unisex_bathroom.Person {
	var p unisex_bathroom.Person
	result := make([]unisex_bathroom.Person, 0, count)

	for i := 0; i < count; i++ {
		if rand.Int()%2 == 0 {
			p = unisex_bathroom.Person{
				Name: petname.Generate(1, ""),
				Sex:  true,
			}
		} else {
			p = unisex_bathroom.Person{
				Name: petname.Generate(1, ""),
				Sex:  false,
			}
		}
		result = append(result, p)
	}

	return result
}

func main() {
	var wg sync.WaitGroup
	wg.Add(N)
	bath := unisex_bathroom.NewBath(3)
	persons := CreatePersons(N)
	go bath.Run()

	for i := 0; i < len(persons); i++ {
		go func(p unisex_bathroom.Person) {
			if p.Sex {
				bath.MenIn <- p
				time.Sleep(time.Millisecond * 10)
				bath.MenOut <- p
			} else {
				bath.WomenIn <- p
				time.Sleep(time.Millisecond * 10)
				bath.WomenOut <- p
			}
			wg.Done()
		}(persons[i])
	}

	wg.Wait()
	bath.Close <- struct{}{}
}
