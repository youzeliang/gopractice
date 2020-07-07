package main

import (
	"fmt"
)

type User struct {
	Name string
	Email string
}

func (u *User) Notify() error {
	fmt.Printf("User: Sending User Email to %s<%s>\n", u.Name, u.Email)
	return nil
}

type Notifier interface {
	Notify() error
}

func SendNotification(notify Notifier) error {
	return notify.Notify()
}

type Admin struct {
	User
	Level string
}

func main() {
	user := User{
		Name: "jane",
		Email: "jane@email.com",
	}
	admin := &Admin{
		User: user,
		Level: "super",
	}
	_ = SendNotification(admin) //User: Sending User Email To jane<jane@email.com>
}