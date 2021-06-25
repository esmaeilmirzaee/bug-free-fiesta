package main

import (
	"fmt"
	"log"
	"strconv"

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

		if char == 'q' || char == 'Q' {
			break
		}

		coffees := make(map[int]string)
		coffees[1] = "Latte"
		coffees[2] = "Americano"

		i, _ := strconv.Atoi(string(char))
		t := fmt.Sprintf("Press a key to choose %q", coffees[i])

		if key != 0 {
			fmt.Println("You pressed", t)
		} else {
			fmt.Println("You pressed", t, char)
		}

	}
	log.Fatal("You pressed ESC. Quiting...")

}
