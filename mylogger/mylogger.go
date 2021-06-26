package mylogger

import "log"

func ListenToLog(ch chan string) {
	for {
		input := <-ch
		log.Println(input)
	}
}
