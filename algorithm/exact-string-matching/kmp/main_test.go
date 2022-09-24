package main

import (
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_patternMatchTable(t *testing.T) {
	type args struct {
		w string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test",
			args: args{
				w: "aabaaba",
			},
		},
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
		// TODO: Add test cases.
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

func Benchmark_kmp(b *testing.B) {
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
				word: "tested",
				text: "Testers have tested many programs.",
				t:    []int{-1, 0, 0, 0, 1, 2},
			},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				kmp(tt.args.word, tt.args.text, tt.args.t)
			}
		})
	}
}

func Benchmark_nonkmp(b *testing.B) {
	type args struct {
		word string
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				word: "tested",
				text: "Testers have tested many programs.",
			},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				nonkmp(tt.args.word, tt.args.text)
			}
		})
	}
}