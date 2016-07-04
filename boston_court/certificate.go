package boston_court

import (
	"fmt"
)

type Certificate struct {
	CourtName string
	UUID int
}

func (cert Certificate) String() string {
	return fmt.Sprintf("<Judge: %s, UUID: %d >", cert.CourtName, cert.UUID)
}

func NewCertificate(court Court) Certificate {
	return Certificate{
		CourtName: court.CourtName,
		UUID: court.GetUUID(),
	}
}