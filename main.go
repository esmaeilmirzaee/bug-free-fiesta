package main

import "log"

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

	
}