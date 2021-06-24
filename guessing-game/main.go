package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "and press 'ENTER' when you are ready."

func main() {
	firstNumber := 2
	secondNumber := 5
	subtractionNumber := 7
	answer := 0

	reader := bufio.NewReader(os.Stdin)

	// Introduction message
	fmt.Println("Welcome to guessing game:")
	fmt.Println("-------------------------")
	fmt.Println()

	// The process
	fmt.Println("Please think of a number between 1 and 10,", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply the thought number with", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply the thought number with", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Subtract the number with the ", subtractionNumber, prompt)
	reader.ReadString('\n')

	// Show the result
	answer = firstNumber*secondNumber - subtractionNumber
	fmt.Println(answer)
}
