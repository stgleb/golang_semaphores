package child_care_center

import "fmt"

type Adult struct {
	Name string
}

func (adult *Adult) String() string {
	return fmt.Sprintf("Adult: %s", adult.Name)
}
