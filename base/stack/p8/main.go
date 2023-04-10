package main

func t0() {
	num := 1
	t1(&num)
}

type user struct {
	name *int
}

func t1(n *int) *user {
	u := &user{}
	u.name = n
	return u
}
func main() {
	t0()
}
