package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		keyboard.Close()
	}()

	fmt.Println("Press any key you like to get the ascii code, to quit press ESC.")

	for {
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if key != 0 {
			fmt.Println("You pressed", char)
		} else {
			fmt.Println("You pressed", char, key)
		}

		if key == keyboard.KeyEsc {
			break
		}
	}
	log.Fatal("You pressed ESC. Quiting...")

}
