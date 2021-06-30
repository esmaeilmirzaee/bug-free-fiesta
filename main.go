package main

import (
	"log"
	"fmt"
	"reflect"
)

func main() {
	// Array
	var scores [5]float64
	names := [3]string{"Esmaeil", "Sam", "John"}
	for _, score := range scores {
		log.Println(score)
	}

	for _, name:= range names {
		log.Println(name)
	}

	// Slice
	fruits := []string{"apple", "orange", "pear", "kumquat"}
	sliceOfFruits := fruits[1:3]
	fmt.Printf("%T\n", sliceOfFruits)
	log.Println(cap(sliceOfFruits))
	log.Println("->", reflect.TypeOf(fruits))

	fruits = append(fruits, "cantelope", "cherries")
	log.Println(fruits, len(fruits), cap(fruits))

	log.Println(ages())
}

func ages() (ageOfBob int, ageOfSally int) {
	ageOfBob = 21
	ageOfSally = 34
	return 
}