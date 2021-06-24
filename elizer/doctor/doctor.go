package doctor

import (
	"log"
	"regexp"
)

func Intro() string {
	return `Hello`
}

func Echo(userInput string) string {
	return userInput
}

func Response(userInput string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	userInput = reg.ReplaceAllString(userInput, " ")

	return userInput
}
