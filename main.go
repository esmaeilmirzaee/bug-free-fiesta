package main

import "log"

func main() {

	var coffees = make(map[int]string)
	coffees[0] = "Latte"

	value, ok := coffees[0]
	log.Println(ok, value)

	if _, ok := coffees[0]; ok {
		log.Println(coffees[0])
	}
}
