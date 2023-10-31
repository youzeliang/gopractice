package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	assert.True(t, GetInstance() == GetInstance())
	assert.False(t, GetInstance() == GetLazyInstance())
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstance() != GetInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
