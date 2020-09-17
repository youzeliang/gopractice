package main

import (
	"reflect"
	"testing"
)

func Test_filter(t *testing.T) {
	type args struct {
		s []student
		f func(student) bool
	}
	tests := []struct {
		name string
		args args
		want []student
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filter(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
