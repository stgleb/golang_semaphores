package main

import (
	"../../boston_court"
)

/*
There are three kinds of threads: immigrants, spectators, and a one judge.
Immigrants must wait in line, check in, and then sit down.
At some point, the judge enters the building. When the judge is in the building, no one may enter,
and the immigrants may not leave. Spectators may leave. Once all immigrants check in,
the judge can confirm the naturalization. After the confirmation, the immigrants pick up their
certificates of U.S. Citizenship. The judge leaves at some point after the confirmation.
Spectators may now enter as before. After immigrants get their certificates, they may leave.
*/


func makeSpectators(court *boston_court.Court) [] boston_court.Spectator {
	return []boston_court.Spectator{
		boston_court.NewSpectator("Sansa Stark", court),
		boston_court.NewSpectator("Arya Stark", court),
		boston_court.NewSpectator("Sandor Clegane", court),
		boston_court.NewSpectator("Bronn", court),
	}
}

func makeImmigrants(court *boston_court.Court) []boston_court.Immigrant {
	immigrants := []boston_court.Immigrant{
		boston_court.NewImmigrant("John Snow", court),
		boston_court.NewImmigrant("Peter Baelish", court),
		boston_court.NewImmigrant("Jaime Lannister", court),
		boston_court.NewImmigrant("Cersei Lannister", court),
	}

	return immigrants
}


func main() {
	boston_court.Info.Println("Court is opened")
	court := boston_court.NewCourt("Boston court #1")
	judge := boston_court.NewJudge("Tywin Lannister", court)
	spectators := makeSpectators(&court)
	immigrants := makeImmigrants(&court)

	for _, spectator := range spectators {
		go spectator.Run()
	}

	for  _, immigrant := range immigrants {
		go immigrant.Run()
	}

	go judge.Run()
	// Start main loop
	go court.Run()

	// Deadlock prone solution as far as some immigrants
	// can be "lost" in channel
	//for i := 0;i < len(immigrants);i++{
	//	<- court.Exit
	//}

	select {
	case <- court.Exit:
		break
	}
}
