package main

import (
	"bytes"
	"fmt"
)

type User struct {
	UserId   int
	UserName string
}
type UserTag struct {
	UserId   int    `json:"user_id" bson:"user_id"`
	UserName string `json:"user_name" bson:"user_name"`
}

type MyInt1 int
type MyInt2 = int

func hello() []string {
	return nil
}

type Person struct {
	age int
}

const (
	a = iota
	b = iota
)
const (
	name     = "name"
	fdfsdfsd = "fafa"
	c        = iota
	d        = iota
)

func main() {

	//	sb := []byte("da")
	////	st := []byte("aa")
	////
	////println(	bytes.IndexByte(sb, sb[1]))
	////println(	bytes.IndexByte(st, st[1]))

	fmt.Println(1)

}

func isIsomorphic(s string, t string) bool {
	sb := []byte(s)
	st := []byte(t)

	if len(sb) == 0 && len(st) == 0 {
		return true
	}

	if len(sb) != len(st) {
		return false
	}

	for i := 0; i < len(sb); i++ {
		if bytes.IndexByte(sb, sb[i]) != bytes.IndexByte(st, st[i]) {
			return false
		}
	}

	return true
}

func change(s ...int) {
	s = append(s, 3)
}

type S struct {
	m string
}

func f() *S {
	return &S{"foo"} //A
}

func tes() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
