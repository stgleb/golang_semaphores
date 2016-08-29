package unisex_bathroom

import "fmt"

type Person struct {
	Name string
	Sex  bool
}

func (person Person) getSex() string {
	if person.Sex {
		return "Male"
	}

	return "Female"
}

func (person Person) String() string {
	return fmt.Sprintf("[Name: %s, Sex: %s]",
		person.Name,
		person.getSex())
}
