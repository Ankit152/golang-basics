package main

import "fmt"

// struct student
type Student struct {
	name, department string
	year             int
}

// struct teacher
type Teacher struct {
	name, subject string
	exp           int
	stud          []Student
}

func main() {
	// nested structure code

	result := []Teacher{
		{
			name:    "James Gosling",
			subject: "Java",
			exp:     40,
			stud:    []Student{{"Ankit", "CSE", 4}},
		},
		{
			name:    "Denis Ritchie",
			subject: "C Programming",
			exp:     40,
			stud: []Student{
				{
					name:       "Ankit",
					department: "CSE",
					year:       4,
				},
				{
					name:       "Tony Stark",
					department: "CSE",
					year:       4,
				},
			},
		},
		{
			name:    "Guido van Rossum",
			subject: "Python",
			exp:     30,
			stud: []Student{
				{
					name:       "Bruce Wayne",
					department: "CSE",
					year:       4,
				},
				{
					name:       "Tony Stark",
					department: "CSE",
					year:       4,
				},
				{
					name:       "Ankit",
					department: "CSE",
					year:       4,
				},
			},
		},
	}

	// some operations on the nested structure
	fmt.Printf("Total teachers: %d\n", len(result))

	// list of teachers with more than 1 students
	fmt.Printf("List of teachers with more than 1 students: \n")
	for _, teacher := range result {
		if len(teacher.stud) >= 2 {
			fmt.Printf("Teacher name: %v and total students: %v\n", teacher.name, len(teacher.stud))
		}
	}

	// iterating over the students of each teacher
	fmt.Println("List of teachers and students:")
	for _, teacher := range result {
		fmt.Printf("Teacher name: %v\n", teacher.name)
		fmt.Printf("Student names: ")
		for _, student := range teacher.stud {
			fmt.Printf("%v - ", student.name)
		}
		fmt.Println("End")
	}
}
