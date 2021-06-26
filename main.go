package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	logger "myapp/mylogger"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)

	go logger.ListenToLog(ch)

	fmt.Println("Please enter something")
	for i := 0; i < 5; i++ {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second)
	}
}
