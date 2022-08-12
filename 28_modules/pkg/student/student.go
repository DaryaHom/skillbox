package student

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name       string
	Age, Grade int
}

//GetStudents - makes a slice of students with user's input data
func GetStudents() (students []Student) {
	students = make([]Student, 0)
	fmt.Println("Enter student's data:")
	for {
		data, err := readStudentData() //read input data
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF received")
				return
			}
			fmt.Println("Error:", err)
			return
		}
		s, err := createStudent(data) //create new Student
		if err != nil {
			fmt.Println("Error:", err)
		}
		emptyStruct := Student{}
		if s != emptyStruct {
			students = append(students, s) //add new Student to slice
		}
	}
}

//readStudentData - reads input students in an infinite loop
func readStudentData() (string, error) {
	var data string
	fmt.Print("-->")
	s, err := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.Trim(s, "\n")
	s = strings.Trim(s, "\r")
	if err != nil {
		return "", err
	}
	data = s
	return data, nil
}

//createStudent - creates new Student struct
func createStudent(data string) (Student, error) {
	if data == "" {
		err := fmt.Errorf("invalid Student's data. Please try again")
		return Student{}, err
	}
	arr := strings.Split(data, " ")
	if len(arr) < 3 {
		err := fmt.Errorf("invalid Student's data. Please try again")
		return Student{}, err
	}

	name := arr[0]
	age, err := strconv.Atoi(arr[1])
	if err != nil {
		err := fmt.Errorf("invalid age value. Please try again")
		return Student{}, err
	}

	grade, err := strconv.Atoi(arr[2])
	if err != nil {
		fmt.Println("Invalid grade value. Please try again")
		return Student{}, err
	}

	newStudent := Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
	return newStudent, nil
}
