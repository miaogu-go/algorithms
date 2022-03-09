package FindDiagonalOrder

import (
	"reflect"
	"testing"
)

func Test_findDiagonalOrder(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				mat: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{
			name: "2",
			args: args{
				mat: [][]int{
					{1, 2},
					{3, 4},
				},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "3",
			args: args{
				mat: [][]int{
					{1, 2, 3, 4, 5},
					{6, 7, 8, 9, 10},
					{11, 12, 13, 14, 15},
					{16, 17, 18, 19, 20},
				},
			},
			want: []int{1, 2, 6, 11, 7, 3, 4, 8, 12, 16, 17, 13, 9, 5, 10, 14, 18, 19, 15, 20},
		},
		{
			name: "4",
			args: args{
				mat: [][]int{
					{3},
					{2},
				},
			},
			want: []int{3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDiagonalOrder(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
