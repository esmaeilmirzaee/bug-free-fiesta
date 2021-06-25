package main

import (
	"log"

	"github.com/eiannone/keyboard"
)

var pressedKeyChan chan rune

func main() {
	pressedKeyChan = make(chan rune)
	go printPressedKey()
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		keyboard.Close()
	}()

	for {
		log.Println("Press a key to continue.")
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		pressedKeyChan <- char
	}
}

func printPressedKey() {
	for {
		key := <-pressedKeyChan
		log.Println(string(key), "is pressed.")
	}
}
