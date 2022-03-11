package LongestCommonPrefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: "2",
			args: args{strs: []string{"dog", "racecar", "car"}},
			want: "",
		},
		{
			name: "3",
			args: args{strs: []string{"leets", "leetcode", "leetc", "leeds"}},
			want: "lee",
		},
		{
			name: "4",
			args: args{strs: nil},
			want: "",
		},
		{
			name: "5",
			args: args{strs: []string{"flow", "flower", "flight"}},
			want: "fl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonPrefix2(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{strs: []string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: "2",
			args: args{strs: []string{"dog", "racecar", "car"}},
			want: "",
		},
		{
			name: "3",
			args: args{strs: []string{"leets", "leetcode", "leetc", "leeds"}},
			want: "lee",
		},
		{
			name: "4",
			args: args{strs: nil},
			want: "",
		},
		{
			name: "5",
			args: args{strs: []string{"flow", "flower", "flight"}},
			want: "fl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix2(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix2() = %v, want %v", got, tt.want)
			}
		})
	}
}
