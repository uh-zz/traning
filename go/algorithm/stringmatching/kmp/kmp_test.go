package kmp

import (
	"reflect"
	"testing"
)

func Test_patternMatchTable(t *testing.T) {
	type args struct {
		w string
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
			if got := patternMatchTable(tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("patternMatchTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kmp(t *testing.T) {
	type args struct {
		word string
		text string
		t    []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				word: "ababaca",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kmp(tt.args.word, tt.args.text, tt.args.t)
		})
	}
}

func Test_nonkmp(t *testing.T) {
	type args struct {
		word string
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nonkmp(tt.args.word, tt.args.text)
		})
	}
}
