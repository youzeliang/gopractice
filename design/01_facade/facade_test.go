package facade

import "testing"

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

	var z  = []string{"3432","432"}

	var zzz = [][]string{{"fds","fd","fds"}}


}
