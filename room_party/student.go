package room_party

import "fmt"

type Student struct {
	Name string
}

func (student *Student) String() string {
	return fmt.Sprintf("Student: %s", student.Name)
}
