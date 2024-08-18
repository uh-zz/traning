package sort

import (
	"sort"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	type args struct {
		x func() sort.Interface
	}
	tests := map[string]struct {
		args args
	}{
		"numeric": {
			args: args{
				x: func() sort.Interface {
					return sort.IntSlice([]int{1, 2, 3})
				},
			},
		},
		"string(jis)": {
			args: args{
				x: func() sort.Interface {
					return sort.StringSlice([]string{"あ", "い", "う"})
				},
			},
		},
		"string(jis with sort)": {
			args: args{
				x: func() sort.Interface {
					x := []string{"ち", "お", "ん", "ま", "こ", "さ", "ら"}
					sort.Sort(sort.StringSlice(x))
					return sort.StringSlice(x)
				},
			},
		},
	}
	for casename, tt := range tests {
		t.Run(casename, func(t *testing.T) {
			x := tt.args.x()
			for {
				if !NextPermutation(x) {
					break
				}
				t.Logf("%v", x)
			}
		})
	}
}
