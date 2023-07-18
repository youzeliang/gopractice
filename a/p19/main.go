package main

import (
	"fmt"
	"os"
)

func getUsers() ([]string, error) {
	return []string{"a", "b", "c", "d"}, nil
}

func main() {
	var users = make([]string, 0)

	envUsers := os.Getenv("USERS")
	if envUsers == "" {
		fmt.Println("Get users from db")
		users, err := getUsers()
		if err != nil {
			panic("ERROR!")
		}
		fmt.Println("Users total: ", len(users))
	}

	for _, user := range users {
		fmt.Println(user)
	}
}
