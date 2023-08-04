package problems

import "testing"

func Test_canCompleteCircuit(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}}, 3},
		{"2", args{[]int{2, 3, 4}, []int{3, 4, 3}}, -1},
		{"3", args{[]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canCompleteCircuit(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("canCompleteCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}
