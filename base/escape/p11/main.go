package main

//go:noinline
func makeBuffer() []byte {
	return make([]byte, 1)
}

func main() {
	buf := makeBuffer()
	for i := range buf {
		buf[i] = buf[i] + 1
	}
}
