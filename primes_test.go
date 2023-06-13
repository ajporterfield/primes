package primes

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestFirst(t *testing.T) {
	type test struct {
		input int
		want  []int
	}

	tests := []test{
		{input: -1, want: make([]int, 0)},
		{input: 0, want: make([]int, 0)},
		{input: 1, want: []int{2}},
		{input: 3, want: []int{2, 3, 5}},
		{input: 10, want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
	}

	for _, tc := range tests {
		got := First(tc.input)
		if !slices.Equal(got, tc.want) {
			t.Errorf("First(%d) = %v; want %v", tc.input, got, tc.want)
		}
	}
}

func TestList(t *testing.T) {
	type test struct {
		count int
		start int
		want  []int
	}

	tests := []test{
		{count: -1, start: -2, want: make([]int, 0)},
		{count: -2, start: -1, want: make([]int, 0)},
		{count: 1, start: 0, want: []int{2}},
		{count: 1, start: 2, want: []int{2}},
		{count: 10, start: 100, want: []int{101, 103, 107, 109, 113, 127, 131, 137, 139, 149}},
	}

	for _, tc := range tests {
		got := List(tc.count, tc.start)
		if !slices.Equal(got, tc.want) {
			t.Errorf("List(%d, %d) = %v; want %v", tc.count, tc.start, got, tc.want)
		}
	}
}

func TestBetween(t *testing.T) {
	type test struct {
		start int
		end   int
		want  []int
	}

	tests := []test{
		{start: 3, end: 7, want: []int{3, 5, 7}},
		{start: 7, end: 3, want: []int{3, 5, 7}},
	}

	for _, tc := range tests {
		got := Between(tc.start, tc.end)

		if !slices.Equal(got, tc.want) {
			t.Errorf("Between(%d, %d) = %v; want %v", tc.start, tc.end, got, tc.want)
		}
	}
}

func TestUpTo(t *testing.T) {
	type test struct {
		input int
		want  []int
	}

	tests := []test{
		{input: 10, want: []int{2, 3, 5, 7}},
		{input: -1, want: []int{}},
	}

	for _, tc := range tests {
		got := UpTo(tc.input)

		if !slices.Equal(got, tc.want) {
			t.Errorf("UpTo(%d) = %v; want %v", tc.input, got, tc.want)
		}
	}
}

func TestNext(t *testing.T) {
	type test struct {
		input int
		want  int
	}

	tests := []test{
		{input: 10, want: 11},
		{input: -1, want: 2},
	}

	for _, tc := range tests {
		got := Next(tc.input)

		if got != tc.want {
			t.Errorf("Next(%d) = %v; want %v", tc.input, got, tc.want)
		}
	}
}

func TestPrev(t *testing.T) {
	type test struct {
		input int
		want  int
	}

	tests := []test{
		{input: 11, want: 7},
		{input: 2, want: 0},
		{input: -1, want: 0},
	}

	for _, tc := range tests {
		got, err := Prev(tc.input)

		if got != tc.want {
			t.Errorf("Next(%d) = %v; want %v", tc.input, got, tc.want)
		}

		if tc.input > 2 && err != nil {
			t.Errorf("Next(%d) error should be nil", tc.input)
		}

		if tc.input <= 2 && err == nil {
			t.Errorf("Next(%d) error shouldn't be nil", tc.input)
		}
	}
}

func TestIsPrime(t *testing.T) {
	type test struct {
		input int
		want  bool
	}

	tests := []test{
		{input: -1, want: false},
		{input: 0, want: false},
		{input: 1, want: false},
		{input: 2, want: true},
		{input: 3, want: true},
		{input: 4, want: false},
		{input: 5, want: true},
	}

	for _, tc := range tests {
		got := IsPrime(tc.input)
		if got != tc.want {
			t.Errorf("IsPrime(%d) = %t; want %t", tc.input, got, tc.want)
		}
	}
}
