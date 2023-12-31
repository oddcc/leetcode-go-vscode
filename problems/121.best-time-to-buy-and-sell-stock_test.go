package problems

import "testing"

func Test_maxProfit121(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{7, 1, 5, 3, 6, 4}}, 5},
		{"2", args{[]int{7, 6, 4, 3, 1}}, 0},
		{"3", args{[]int{2, 7, 1, 5, 4}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit_deleteToSubmit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
