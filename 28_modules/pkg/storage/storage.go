package storage

import (
	"fmt"
	st "modules/pkg/student"
)

type StudentsData map[string]*st.Student

//Put - writes Student's data to the map StudentsData
func (m StudentsData) Put() {
	//reading student's data into the slice of Student
	s := st.GetStudents()
	for i := range s {
		m[s[i].Name] = &s[i]
	}
}

//Get - prints the contents of StudentsData
func (m StudentsData) Get() {
	for i := range m {
		fmt.Println(m[i].Name, m[i].Age, m[i].Grade)
	}
}
