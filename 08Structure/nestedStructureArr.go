package main

import "fmt"

func main() {
	type Marks struct {
		physics   int
		chemistry int
		maths     int
	}

	nestedStructArr := []struct {
		name  string
		roll  int
		marks Marks
	}{
		{
			name: "Ankit Kurmi",
			roll: 1,
			marks: Marks{
				physics:   92,
				chemistry: 88,
				maths:     97,
			},
		},
		{
			name: "Test-2",
			roll: 2,
			marks: Marks{
				physics:   60,
				chemistry: 60,
				maths:     88,
			},
		},
	}

	for i, val := range nestedStructArr {
		fmt.Printf("%v --- %v\n", i, val)
	}
}
