package room_party

import "fmt"

type Dean struct {
	Name string
}

func (dean *Dean) String() string {
	return fmt.Sprintf("Dean: %s", dean.Name)
}

func NewDean(name string) Dean {
	return Dean{
		Name: name,
	}
}
