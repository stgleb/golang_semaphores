package boston_court

import "fmt"

type Immigrant struct {
	Name string
	Court Court
	CertificateChannel chan Certificate
}

func NewImmigrant(name string, court Court) Immigrant {
	return Immigrant{
		Name: name,
		Court: court,
		CertificateChannel: make(chan Certificate),
	}
}

func (immigrant Immigrant) String() string {
	return fmt.Sprintf("< Name: %s >", immigrant.Name)
}

func (immigrant *Immigrant) Run() {
	immigrant.Court.Entrance <- immigrant
	immigrant.CertificateChannel
	immigrant.Court.ImmigrantOut <- immigrant
}


