package mid

import "testing"

func Test_bin(t *testing.T) {
	type args struct {
		arr []int
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				arr: []int{1, 3, 5, 7, 9},
				val: 3,
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				arr: []int{1, 3, 5, 7, 9},
				val: 9,
			},
			want: 4,
		},
		{
			name: "3",
			args: args{
				arr: []int{1, 3, 5, 7, 9},
				val: 10,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bin(tt.args.arr, tt.args.val); got != tt.want {
				t.Errorf("bin() = %v, want %v", got, tt.want)
			}
		})
	}
}
