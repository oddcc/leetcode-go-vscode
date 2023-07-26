package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name     string
		args     args
		want     int
		expected []int
	}{
		{"1", args{[]int{1, 1, 2}}, 2, []int{1, 2}},
		{"2", args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, 5, []int{0, 1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}

			assert.ElementsMatch(t, tt.expected, tt.args.nums[:tt.want])
		})
	}
}
