package array

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Gender    string
}

func ArrayTest() {
	testMalePersons()
}

func testMalePersons() {
	var persons [2]Person
	persons[0] = Person{
		FirstName: "F1",
		LastName:  "L1",
		Gender:    "M",
	}
	persons[1] = Person{
		FirstName: "F2",
		LastName:  "L2",
		Gender:    "F",
	}

	makePersons := getMalePersons(persons)

	for _, p := range makePersons {
		fmt.Println("FirstName: ", p.FirstName)
	}
}

func getMalePersons(persons [2]Person) (malePersons []Person) {
	for _, p := range persons {
		if p.Gender == "M" {
			malePersons = append(malePersons, p)
		}
	}
	return
}
