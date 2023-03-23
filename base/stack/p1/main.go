package main

type user struct {
	name string
}

func getUser() *user {

	user := user{}
	user.name = "Rolle"
	return &user
}
func main() {

}
