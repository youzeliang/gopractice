package facade

import (
	"fmt"
	"testing"
)

var expect = "A module running\nB module running"

// TestFacadeAPI ...
func TestFacadeAPI(t *testing.T) {
	api := NewAPI()
	ret := api.Test()
	if ret != expect {
		t.Fatalf("expect %s, return %s", expect, ret)
	}

	var n []string
	n = append(n, "fd")

	var z = []string{"3432", "432"}

	fmt.Println(z)

	var zzz = [][]string{{"fds", "fd", "fds"}}

	fmt.Println(zzz)
}
