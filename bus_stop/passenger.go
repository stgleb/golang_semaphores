package bus_stop

import "fmt"

type Passenger struct {
	FirstName string
	LastName  string
}

func (p Passenger) String() string {
	return fmt.Sprintf("< %s %s >", p.FirstName, p.LastName)
}

func NewPassenger(firstName, lastName string) Passenger {
	return Passenger{
		FirstName: firstName,
		LastName:  lastName,
	}
}
