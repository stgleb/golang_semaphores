package boston_court

import "fmt"

type Immigrant struct {
	Name string
	Court *Court
	CertificateChannel chan Certificate
	cert Certificate
}

func NewImmigrant(name string, court *Court) Immigrant {
	return Immigrant{
		Name: name,
		Court: court,
		CertificateChannel: make(chan Certificate),
	}
}

func (immigrant Immigrant) String() string {
	return fmt.Sprintf("< Name: %s >", immigrant.Name)
}

func (immigrant Immigrant) Run() {
	Info.Printf("Immigrant %s has come to country", immigrant)
	// Proof that immigrant has successfully entered to court
	select {
		case immigrant.Court.Entrance <- immigrant:
			Info.Printf("Immigrant %s has entered", immigrant)
	}
	immigrant.cert = <- immigrant.CertificateChannel
	Info.Printf("Immigrant %s has obtained certificate %s", immigrant, immigrant.cert)
	immigrant.Court.ImmigrantOut <- immigrant
}


