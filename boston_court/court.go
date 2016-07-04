package boston_court

type Court struct {
	Entrance chan struct{}
	ImmigrantOut chan Immigrant
	SpectatorOut chan Spectator
	JudgeOut chan Judge
	CertificateChan chan Certificate

	immigrantCount int
	spectatorCount int

	immigrantBench []Immigrant
	spectatorBench []Spectator

	CourtName string
	seq int
}

func NewCourt(courtName string) Court {
	return Court{
		Entrance: make(chan struct{}),
		ImmigrantOut: make(chan Immigrant),
		SpectatorOut: make(chan Spectator),
		JudgeOut: make(chan Judge),
		CertificateChan: make(chan Certificate),
		immigrantCount: 0,
		spectatorCount: 0,
		immigrantBench: make([]Immigrant, 0, 10),
		spectatorBench: make([]Spectator, 0, 10),
		CourtName: courtName,
		seq: 0,
	}
}

func (court Court) Run() {
	for {
		select {
		case person := <- court.Entrance:
			switch person.(type) {
			case Immigrant:
				court.immigrantBench = append(court.immigrantBench, person)
				court.immigrantCount++
			case Spectator:
				court.spectatorBench = append(court.spectatorBench, person)
				court.spectatorCount++
			case Judge:
				// Block entrance channel
				court.Entrance = nil
			}
		case spectator := court.SpectatorOut:
			
		}
	}
}

func (court Court) AssignCertificates() {
	for immigrant := range court.immigrantBench {
		immigrant.CertificateChannel <- NewCertificate(court)
	}
}

func (court Court) GetUUID() int {
	court.seq++
	return court.seq
}