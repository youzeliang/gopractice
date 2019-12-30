package _551_student_attendance_record_i

import "strings"

func checkRecord(s string) bool {
	absent := 0
	if strings.Contains(s, "LLL") {
		return false
	}
	for _, v := range s {
		if v == 'A' {
			absent++
		}
		if absent > 1 {
			return false
		}
	}
	if absent > 1 {
		return false
	}
	return true
}

func checkRecord2(s string) bool {
	return !(strings.Count(s, "A") > 1 || strings.Contains(s, "LLL"))
}
