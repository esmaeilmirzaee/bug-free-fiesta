package main

import (
	"log"
	"math/rand"
	"myapp/auth"
)

func main() {
	u1 := auth.User{Id: rand.Intn(1000), Name: "Esmaeil MIRZAEE", Email: "em@gmail.com", Password: "123", Admin: false}

	u2 := auth.User{Id: rand.Intn(1000), Name: "Samuel Russell", Email: "samrus@g.com", Password: "123"}

	var users []auth.User
	users = append(users, u1, u2)
	groupOne := auth.Group{Users: users, Enabled: true}

	log.Println(groupOne.Describe())
	log.Println(u1.Describe())
	auth.PromoteUser(u1)
	log.Println(u1)
	auth.UpdateEmail(&u1, "me@gmail.com")
	log.Println(u1)
}
