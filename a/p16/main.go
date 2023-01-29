package main

func main() {

	fillOne(1000000, 1000000)
}

func fillOne(m, n int) {

	res := [1000][1000]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[j][i] = 0
		}
	}

}
