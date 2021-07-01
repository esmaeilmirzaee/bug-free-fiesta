package auth

import "fmt"

type User struct {
	Id                    int
	Name, Email, Password string
	Admin                 bool
}

type Group struct {
	Users   []User
	Enabled bool
}

func PromoteUser(u User) bool {
	u.Admin = true
	return true
}

func UpdateEmail(u *User, e string) {
	u.Email = e
}

func (u User) Describe() string {
	return fmt.Sprintf("User %s owns %s address. Is s/he admin? %t", u.Name, u.Email, u.Admin)
}

func (g Group) Describe() string {
	return fmt.Sprintf("The current group has about %d users and it is enabled: %t", len(g.Users), g.Enabled)
}
