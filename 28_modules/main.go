package main

import (
	str "modules/pkg/storage"
)

func main() {
	//creating new map[string]*Student
	m := str.StudentsData{}
	//filling map with student's data
	m.Put()
	//after-EOF code
	m.Get()
}
