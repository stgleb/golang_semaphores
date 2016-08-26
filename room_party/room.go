package room_party

import (
	"log"
	"time"
)

type Room struct {
	StudentIn  chan Student
	StudentOut chan Student
	DeanIn     chan Dean
	DeanOut    chan Dean

	studentCount  int
	roomCapacity  int
	deanInRoom    bool
	partyDuration time.Duration
}

func NewRoom(roomCapacity int, timeout time.Duration) Room {
	return Room{
		StudentIn:     make(chan Student),
		StudentOut:    make(chan Student),
		DeanIn:        make(chan Dean),
		DeanOut:       make(chan Dean),
		studentCount:  0,
		roomCapacity:  roomCapacity,
		deanInRoom:    false,
		partyDuration: timeout,
	}
}

func (room *Room) Run() {
	var (
		sIn chan Student
		dIn chan Dean
	)

	for {
		select {
		case student := <-room.StudentIn:
			log.Printf("In %s", student)

			// Dean can not enter while there are students in the room.
			if room.studentCount == 0 {
				dIn = room.DeanIn
			}

			// Dean may enter to break up the room.
			if room.studentCount > room.roomCapacity {
				room.DeanIn = dIn
			}
			room.studentCount++
		case student := <-room.StudentOut:
			log.Printf("Out %s", student)
			room.studentCount--
		case dean := <-room.DeanIn:
			log.Printf("In %s", dean)
			room.deanInRoom = true
			// When dean enters no more students may enter.
			sIn = room.StudentIn
			room.StudentIn = nil
		case dean := <-room.DeanOut:
			log.Printf("Out %s", dean)
			// When dean leaves new students may enter.
			room.StudentIn = sIn
			room.deanInRoom = false
		case <-time.After(room.partyDuration):
			log.Printf("Party timeout has been expired")
		}
	}
}
