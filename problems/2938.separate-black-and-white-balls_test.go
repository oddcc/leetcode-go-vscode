package problems

import "testing"

func Test_minimumSteps(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"1", args{"101"}, 1},
		{"2", args{"100"}, 2},
		{"3", args{"0111"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSteps(tt.args.s); got != tt.want {
				t.Errorf("minimumSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
