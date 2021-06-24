package main

import "fmt"

func main() {
	var whatToSay string = "Hello, World"
	toWhom := "Esmaeil"
	printMessage(whatToSay, toWhom)
}

func printMessage(whatToSay string, toWhom string) {
	fmt.Print(whatToSay, " ", toWhom)
}
