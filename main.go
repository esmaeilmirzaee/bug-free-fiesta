package main

import (
	"log"
	"math/rand"
	"strconv"
	"fmt"
	"time"
)

func main(){
	var scores [5]float32
	rand.Seed(time.Now().UnixNano())
	scores[0] = rand.Float32() * 100
	scores[1] = rand.Float32() * 100
	scores[2] = rand.Float32() * 100
	scores[3] = rand.Float32() * 100
	scores[4] = rand.Float32() * 100
	
	log.Printf("%.2f", scores)
	for _, value := range scores {
		temp := fmt.Sprintf("%.2f", value)
		x := 3.95
		log.Println(x)
		
		if tmp, err := strconv.ParseFloat(temp, 32); err == nil {
			log.Println(tmp)
		} else {
			log.Println(err)
		}	
	}
}