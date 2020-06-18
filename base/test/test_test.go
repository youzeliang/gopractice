package test

import (
	"fmt"
	"testing"
)

func TestTf(t *testing.T) {
	t.Cleanup(func() {
		fmt.Println(32131)
	})
}
