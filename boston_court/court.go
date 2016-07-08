package boston_court


type Court struct {
	Entrance chan interface{}
	ImmigrantOut chan Immigrant
	SpectatorOut chan Spectator
	Exit chan Immigrant
	JudgeOut chan Judge
	CertificateChan chan Certificate

	ImmigrantCount int
	spectatorCount int

	immigrantBench map[*Immigrant]bool
	spectatorBench map[*Spectator]bool

	CourtName string
	seq int
}

func NewCourt(courtName string) Court {
	return Court{
		Entrance: make(chan interface{}),
		ImmigrantOut: make(chan Immigrant),
		SpectatorOut: make(chan Spectator),
		Exit: make(chan Immigrant),
		JudgeOut: make(chan Judge),
		CertificateChan: make(chan Certificate),
		ImmigrantCount: 0,
		spectatorCount: 0,
		immigrantBench: make(map[*Immigrant]bool),
		spectatorBench: make(map[*Spectator]bool),
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
				immigrant, ok := person.(Immigrant)

				if !ok {
					Error.Println("Cannot convert to type Immigrant")
				}
				court.immigrantBench[&immigrant] = true
				court.ImmigrantCount++
				Info.Printf("Immigrant %s has entered", person)
			case Spectator:
				spectator, ok := person.(Spectator)

				if !ok {
					Error.Println("Cannot convert to type Spectator")
				}
				court.spectatorBench[&spectator] = true
				court.spectatorCount++
				Info.Printf("Spectator %s has entered", person)
			case Judge:
				// Block entrance channel
				// NOTE: Bad for reusing channel
				court.Entrance = nil
				Info.Printf("Judge %s has entered", person)
			}
		case spectator := <- court.SpectatorOut:
			Info.Printf("Spectator %s has quit", spectator)
			delete(court.spectatorBench, &spectator)
		case judge := <- court.JudgeOut:
			court.GrantCertificates(judge)
			Info.Printf("Judge %s has quit", judge)
			// Restore entrance channel
			court.Entrance = make(chan interface{})
		case immigrant := <- court.ImmigrantOut:
			Info.Printf("Immigrant %s has quit", immigrant)
			delete(court.immigrantBench, &immigrant)
			court.Exit <- immigrant
		}
	}
}

func (court Court) GrantCertificates(judge Judge) {
	Info.Printf("Judge %s grant ceritifactes for all Immigrants", judge)
	for immigrant, _ := range court.immigrantBench {
		Info.Printf("Grant certificate for %s", immigrant)
		immigrant.CertificateChannel <- NewCertificate(court)
	}
}

func (court Court) GetUUID() int {
	court.seq += 1
	return court.seq
}