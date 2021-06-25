package game

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func GetYesOrNo(msg string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		keyboard.Close()
	}()

	fmt.Println(msg)
	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'n' || char == 'N' {
			return false
		}
		return true
	}

}
