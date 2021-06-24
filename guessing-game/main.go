package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// package level variable
const prompt = "and press 'ENTER' when you are ready."

func main() {
	// block level variables
	firstNumber := 2
	secondNumber := 5
	subtractionNumber := 7
	answer := 0

	reader := bufio.NewReader(os.Stdin)

	Guess(firstNumber, secondNumber, subtractionNumber, answer, reader)
}

func Guess(firstNumber, secondNumber, subtractionNumber, answer int, reader *bufio.Reader) {
	// shadowing variable should be like
	prompt := "please do not share it and press 'ENTER' when you are ready."
	// Introduction message
	fmt.Println("Welcome to guessing game:")
	fmt.Println("-------------------------")
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(8), time.Now(), time.Now().Unix(), time.Now().UnixNano())

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
