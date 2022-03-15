package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1",
			args: args{
				arr:    []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "2",
			args: args{
				arr:    []int{2, 7, 11, 15},
				target: 18,
			},
			want: []int{1, 2},
		},
		{
			name: "3",
			args: args{
				arr:    []int{2, 7, 11, 15},
				target: 13,
			},
			want: []int{0, 2},
		},
		{
			name: "4",
			args: args{
				arr:    []int{1, 2, 3, 4, 5},
				target: 5,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.arr, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				numbers: []int{2, 7, 11, 15},
				target:  9,
			},
			want: []int{1, 2},
		},
		{
			name: "2",
			args: args{
				numbers: []int{2, 3, 4},
				target:  6,
			},
			want: []int{1, 3},
		},
		{
			name: "3",
			args: args{
				numbers: []int{-1, 0},
				target:  -1,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
