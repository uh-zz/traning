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
			args: args{input: &[]int{7, 2, 3, 4, 8, 1, 5, 6, 10}, left: 0, right: 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.input, tt.args.left, tt.args.right)
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(&[]int{7, 2, 3, 4, 8, 1, 5, 6, 10}, 0, 9)
	}
}
