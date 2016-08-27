package child_care_center

import "fmt"

type Child struct {
	Name string
}

func (child *Child) String() string {
	return fmt.Sprintf("Child: %s", child.Name)
}
