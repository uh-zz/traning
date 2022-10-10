package datastructure

import (
	"reflect"
	"testing"
)

func TestLinkedList_Traverse(t *testing.T) {
	type fields struct {
		root *Node
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "test traverse",
			fields: fields{
				root: &Node{
					data: 1,
					next: &Node{
						data: 2,
						next: &Node{
							data: 3,
							next: nil,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				root: tt.fields.root,
			}
			l.Traverse()
		})
	}
}
func TestLinkedList_Insert(t *testing.T) {
	type fields struct {
		root *Node
		len  int
	}
	type args struct {
		current *Node
		data    int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				root: tt.fields.root,
				len:  tt.fields.len,
			}
			if got := l.Insert(tt.args.current, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHoge(t *testing.T) {
	l := NewLinkedList()
	current := NewNode()
	current = l.Insert(current, 1)
	current = l.Insert(current, 2)
	current = l.Insert(current, 3)
	l.Traverse()
}
