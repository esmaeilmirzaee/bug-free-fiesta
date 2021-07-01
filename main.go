package main

import (
	"log"
	"math/rand"
	"myapp/auth"
)

func main() {
	u1 := auth.User{Id: rand.Intn(1000), Name: "Esmaeil MIRZAEE", Email: "em@gmail.com", Password: "123", Admin: false}

	auth.PromoteUser(u1)
	log.Println(u1)
}
