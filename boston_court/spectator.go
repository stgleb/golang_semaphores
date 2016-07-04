package boston_court

import "fmt"

type Spectator struct {
	Name string
	Court Court
}

func NewSpectator(name string, court Court) {
	return Spectator{
		Name: name,
		Court: court,
	}
}

func (spectator Spectator) String() string {
	return fmt.Sprintf("< Name: %s>", spectator.Name)
}

func (spectator Spectator) Run() {
	spectator.Court.Entrance <- spectator
	randSleep(10, 20)
	spectator.Court.SpectatorOut <- spectator
}