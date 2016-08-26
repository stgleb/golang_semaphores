package main

import (
	"../../room_party"
	"github.com/dustinkirkland/golang-petname"
)

/*
   Problem description

   1. Any number of students can be in a room at the same time.
   2. The Dean of Students can only enter a room if there are no students in the room (to conduct a search) or
    if there are more than 50 students in the room (to break up the party).
   3. While the Dean of Students is in the room, no additional students may enter, but students may leave.
   4. The Dean of Students may not leave the room until all students have left.
   5. 	There is only one Dean of Students, so you do not have to enforce exclusion among multiple deans.
*/

func PrepareStudents(count int) []room_party.Student {
	students := make([]room_party.Student, 1, count)

	for i := 0; i < count; i++ {
		s := room_party.Student{
			Name: petname.Generate(1, ""),
		}

		students = append(students, s)
	}

	return students
}

func StudentEnter(room room_party.Room, students []room_party.Student) {
	i := 0
	j := 0

	for {
		select {
		case room.StudentIn <- students[i]:
			i++
		case room.StudentOut <- students[i]:
			j++
		}
	}
}

func DeanEnter(room room_party.Room, dean room_party.Dean) {
	for {
		select {
		case room.DeanIn <- dean:
		case room.DeanOut <- dean:
			break
		}
	}
}

func main() {
	room := room_party.NewRoom(50, 60)
	students := PrepareStudents(50)
	dean := room_party.NewDean("Mike Johnson")
	go room.Run()
	go DeanEnter(room, dean)
	go StudentEnter(room, students)
}
