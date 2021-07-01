package auth

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
