package main

import (
	"log"
	"math"
	"math/rand"
	"time"
)

func main() {
	var scores [5]float64
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		scores[i] = rand.Float64() * 100
	}

	for idx, value := range scores {
		scores[idx] = toFixed(value, 2)
	}

	log.Println(average(scores))
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func average(scores [5]float64) float64 {
	var total float64
	for _, score := range scores {
		total += score
	}
	return total / float64(len(scores))
}
