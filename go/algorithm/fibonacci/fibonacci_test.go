package fibonacci

import (
	"testing"
)

func Test_fibo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				n: 10,
			},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibo(tt.args.n); got != tt.want {
				t.Errorf("fibo() = %v, want %v", got, tt.want)
			}
		})
	}
}
