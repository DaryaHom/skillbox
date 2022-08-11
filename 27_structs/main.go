package main

import (
	"skillbox/27_structs/strt"
)

func main() {
	//reading student's data into the slice of Student
	s := strt.GetStudents()
	//creating new map[string]*Student
	m := strt.StudentsData{}
	//filling map with student's data
	for i := range s {
		m.Put(&s[i])
	}
	//after-EOF code
	m.Get()
}
