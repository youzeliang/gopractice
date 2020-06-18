package main

type geometry interface {
	area() int
	perim() int
}

type rect struct {
	width, height int
}

func area() int {
	return 1 + 2
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}
