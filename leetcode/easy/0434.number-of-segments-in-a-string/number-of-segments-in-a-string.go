package _434_number_of_segments_in_a_string

import (
	"strings"
)

func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}

	strs := strings.Fields(s)
	return len(strs)
}
