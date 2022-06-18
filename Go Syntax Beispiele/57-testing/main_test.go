package main

import (
	"fmt"
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

func TestIntMinTableDrive(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{
			a:    0,
			b:    1,
			want: 0,
		},
		{
			a:    1,
			b:    0,
			want: 0,
		},
		{
			a:    2,
			b:    -2,
			want: -2,
		},
		{
			a:    0,
			b:    -1,
			want: -1,
		},
		{
			a:    -1,
			b:    0,
			want: -1,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

