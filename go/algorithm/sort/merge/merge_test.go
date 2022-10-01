package sort

import "testing"

func TestMergeSort(t *testing.T) {
	type args struct {
		input *[]int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{input: &[]int{7, 2, 3, 4, 8, 1, 5, 6}, left: 0, right: 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.input, tt.args.left, tt.args.right)
		})
	}
}
