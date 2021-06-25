package main

import "log"

// interface
type Animal interface {
	Says() string
	HowManyLegs() int
}

// dog type
type Dog struct {
	Name         string
	Call         string
	NumberOfLegs int
}

func (d *Dog) Says() string {
	return d.Call
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

type Cat struct {
	Name         string
	Call         string
	NumberOfLegs int
	HasTail      bool
}

func (c *Cat) Says() string {
	return c.Call
}

func (c *Cat) HowManyLegs() int {
	return c.NumberOfLegs
}

func main() {
	dog := Dog{
		Name:         "Dog",
		Call:         "Woef",
		NumberOfLegs: 4,
	}

	riddle(&dog)

	var cat Cat
	cat.Name = "Cat"
	cat.Call = "Meow"
	cat.NumberOfLegs = 4
	cat.HasTail = true

	riddle(&cat)
}

func riddle(a Animal) {
	log.Println("An animal has", a.HowManyLegs(), "legs and ", a.Says(), "call. What is it?")
}
