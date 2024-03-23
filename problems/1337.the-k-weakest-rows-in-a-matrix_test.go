package problems

import (
	"reflect"
	"testing"
)

func Test_kWeakestRows(t *testing.T) {
	type args struct {
		mat [][]int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[][]int{
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 0},
			{1, 0, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 1}}, 3}, []int{2, 0, 3}},
		{"2", args{[][]int{
			{1, 0, 0, 0},
			{1, 1, 1, 1},
			{1, 0, 0, 0},
			{1, 0, 0, 0}}, 2}, []int{0, 2}},
		{"3", args{[][]int{
			{1, 0}, 
			{1, 0}, 
			{1, 0}, 
			{1, 1}}, 4}, []int{0, 1, 2, 3}},
		{"4", args{[][]int{
			{1, 1}, 
			{1, 1}, 
			{1, 1}, 
			{1, 1}}, 1}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kWeakestRows(tt.args.mat, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kWeakestRows() = %v, want %v", got, tt.want)
			}
		})
	}
}
