package main

import "log"

type Vehicle struct {
	NumberOfWheels     int
	NumberOfPassengers int
}

func (v *Vehicle) showDetails() {
	log.Println("Number of wheels is", v.NumberOfWheels)
	log.Println("Number of passengers is", v.NumberOfPassengers)
}

type Car struct {
	Make       string
	Model      string
	Year       int
	isElectric bool
	isHybrid   bool
	Vehicle    Vehicle
}

func (c *Car) show() {
	log.Println("Make:", c.Make)
	log.Println("Model:", c.Model)
	log.Println("Year:", c.Year)
	log.Println("is it elctrical:", c.isElectric)
	log.Println("Is it hybrid:", c.isHybrid)
	c.Vehicle.showDetails()
}

func main() {
	suv := Vehicle{
		NumberOfWheels:     4,
		NumberOfPassengers: 6,
	}

	myCar := Car{
		Make:       "Peugeot",
		Model:      "206",
		Year:       2005,
		isElectric: false,
		isHybrid:   false,
		Vehicle:    suv,
	}
	myCar.show()
}
