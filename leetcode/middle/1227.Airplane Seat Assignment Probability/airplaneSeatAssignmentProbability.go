package _227_Airplane_Seat_Assignment_Probability

func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1
	}
	return 0.5
}
