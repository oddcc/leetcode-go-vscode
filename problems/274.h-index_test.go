package problems

import "testing"

func Test_hIndex(t *testing.T) {
	type args struct {
		citations []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 0, 6, 1, 5}}, 3},
		{"2", args{[]int{1, 3, 1}}, 1},
		{"3", args{[]int{5, 3, 1, 0}}, 2},
		{"4", args{[]int{0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex(tt.args.citations); got != tt.want {
				t.Errorf("hIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
