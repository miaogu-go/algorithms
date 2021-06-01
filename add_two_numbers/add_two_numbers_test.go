package add_two_numbers

import (
	"reflect"
	"testing"
)

func Test_addTwoNumber(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				l1: []int{2, 4, 3},
				l2: []int{5, 6, 4},
			},
			want: []int{7, 0, 8},
		},
		{
			name: "2",
			args: args{
				l1: []int{3, 4, 5},
				l2: []int{4, 5, 6},
			},
			want: []int{7, 9, 1, 1},
		},
		{
			name: "3",
			args: args{
				l1: []int{7, 8, 9},
				l2: []int{3, 2},
			},
			want: []int{0, 1, 0, 1},
		},
		{
			name: "4",
			args: args{
				l1: []int{9, 9, 9, 9, 9},
				l2: []int{1},
			},
			want: []int{0, 0, 0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		l1 := createLinkList(tt.args.l1)
		l2 := createLinkList(tt.args.l2)
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumber(l1, l2)
			if got == nil {
				t.Errorf("addTwoNumber() = %v, want %v", got, tt.want)
			}
			arr := make([]int, 0)
			for got != nil {
				arr = append(arr, got.data)
				got = got.next
			}
			if !reflect.DeepEqual(arr, tt.want) {
				t.Errorf("addTwoNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createLinkList(list []int) *Node {
	var head *Node
	var current *Node
	for _, v := range list {
		if head == nil {
			head = &Node{
				data: v,
			}
			current = head
		} else {
			current.next = &Node{
				data: v,
			}
			current = current.next
		}
	}
	return head
}
