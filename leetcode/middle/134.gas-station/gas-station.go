package _34_gas_station

func canCompleteCircuit(gas []int, cost []int) int {

	start, cur, total, n := 0, 0, 0, len(gas)

	for i := 0; i < n; i++ {

		cur += gas[i] - cost[i]

		if cur < 0 {
			start = i + 1
			total += cur
			cur = 0
		}
	}

	if total+cur >= 0 {
		return start
	}

	return -1
}
