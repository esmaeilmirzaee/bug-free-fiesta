package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var scores [5]float32
	rand.Seed(time.Now().UnixNano())
	scores[0] = rand.Float32() * 100
	scores[1] = rand.Float32() * 100
	scores[2] = rand.Float32() * 100
	scores[3] = rand.Float32() * 100
	scores[4] = rand.Float32() * 100

	// log.Printf("%.2f", scores)
	for _, value := range scores {
		str := fmt.Sprintf("%.2f", value)
		log.Println("str->", str)
		if temp, err := strconv.ParseFloat(str, 32); err != nil {
			log.Println(err.Error())
		} else {
			log.Println(temp)
		}
		// tmp := strconv.FormatFloat(temp, 64, 2, 64)
		// log.Println(tmp)
	}
}
