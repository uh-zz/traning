package sort

import "testing"

func TestOtherMergeSort(t *testing.T) {
	type args struct {
		arr   []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{arr: []int{5, 2, 4, 3, 7, 6}, left: 0, right: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OtherMergeSort(tt.args.arr, tt.args.left, tt.args.right)
		})
	}
}

func BenchmarkOtherMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OtherMergeSort([]int{5, 4, 3, 7, 6}, 0, 4)
	}
}
