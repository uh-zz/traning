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
		word         string
		text         string
		patternTable []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				word:         "ab",
				text:         "ababacaab",
				patternTable: patternMatchTable("ababacaab"),
			},
			want: []int{0, 2, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kmp(tt.args.word, tt.args.text, tt.args.patternTable); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kmp() = %v, want %v", got, tt.want)
			}
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
