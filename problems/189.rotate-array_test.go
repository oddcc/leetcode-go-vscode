package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums     []int
		k        int
		expected []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}}},
		{"2", args{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}}},
		{"3", args{[]int{1, 2, 3, 4, 5, 6, 7}, 10, []int{5, 6, 7, 1, 2, 3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
		})

		assert.Equal(t, tt.args.expected, tt.args.nums)
	}
}
