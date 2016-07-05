package main

import (
	"../../boston_court"
	"time"
	"log"
)

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
	go court.Run()

	for _, spectator := range spectators {
		go spectator.Run()
	}

	for  _, immigrant := range immigrants {
		go immigrant.Run()
	}

	go judge.Run()
	// Wait until all immigrants take certificate.
	time.Sleep(time.Millisecond * 2000)
	log.Println(court.WG)
}
