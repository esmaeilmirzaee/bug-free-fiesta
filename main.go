package main

import (
	"log"
	"sort"
)

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animal) Says() {
	log.Println(a.Name, "says", a.Sound)
}

func main() {
	dog := Animal{
		Name:         "dog",
		Sound:        "Woof",
		NumberOfLegs: 4,
	}
	dog.Says()

	var cat Animal
	cat.Name = "Cat"
	cat.Sound = "Meow"
	cat.NumberOfLegs = 4
	cat.Says()

	// Arrays
	var names [3]string
	names[0] = "Esmaeil"
	names[1] = "Sam"
	names[2] = "MIRZAEE"

	log.Println(names)

	// Slice
	var courses []string
	courses = append(courses, "Math")
	courses = append(courses, "Physics")
	courses = append(courses, "English")
	courses = append(courses, "Literature")

	log.Println("----------------------------------")
	log.Println(courses)

	for _, val := range courses {
		log.Println(val)
	}

	for i, val := range names {
		log.Println(i, val)
	}

	log.Println(names[0:2])
	log.Println(courses[0:2])
	log.Println(len(courses))

	log.Println(sort.StringsAreSorted(courses))
	sort.Strings(courses)
	log.Println(courses)

	// Maps
	intMap := make(map[string]string)
	intMap["A"] = "a"
	intMap["B"] = "b"
	intMap["C"] = "c"

	// No one should sort a map variable
	for key, value := range intMap {
		log.Println(key, value)
	}

	delete(intMap, "A")
	for key, value := range intMap {
		log.Println(key, value)
	}

}
