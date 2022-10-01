package shuffle

import (
	"reflect"
	"testing"
)

func TestRandomIntList(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomIntList(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RandomIntList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shuffle(t *testing.T) {
	type args struct {
		array *[]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				array: &[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shuffle(tt.args.array)
		})
	}
}
