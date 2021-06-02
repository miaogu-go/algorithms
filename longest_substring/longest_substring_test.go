package longest_substring

import "testing"

func Test_substringLen(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{s: "abcabcbb"},
			want: 3,
		},
		{
			name: "2",
			args: args{s: "bbbbb"},
			want: 1,
		},
		{
			name: "3",
			args: args{s: "pwwkew"},
			want: 3,
		},
		{
			name: "4",
			args: args{s: ""},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := substringLen(tt.args.s); got != tt.want {
				t.Errorf("substringLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
