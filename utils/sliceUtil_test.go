package utils

import (
	"testing"
)

func TestFind(t *testing.T) {
	type args struct {
		slice []string
		val   string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{"Find first value", args{slice: []string{"one", "two"}, val: "one"}, 0, true},
		{"Find second value", args{slice: []string{"one", "two"}, val: "two"}, 1, true},
		{"Find not found", args{slice: []string{"one", "two"}, val: "three"}, -1, false},
		{"Find not found in empty", args{slice: []string{}, val: "three"}, -1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Find(tt.args.slice, tt.args.val)
			if got != tt.want {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Find() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIsInSlice(t *testing.T) {
	type args struct {
		slice []string
		val   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Find first value", args{slice: []string{"one", "two"}, val: "one"}, true},
		{"Find second value", args{slice: []string{"one", "two"}, val: "two"}, true},
		{"Find not found", args{slice: []string{"one", "two"}, val: "three"}, false},
		{"Find not found in empty", args{slice: []string{}, val: "three"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInSlice(tt.args.slice, tt.args.val); got != tt.want {
				t.Errorf("IsInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintArr(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"PrintArr basic", args{[]string{"one", "two"}}, "[one, two]"},
		{"PrintArr single element", args{[]string{"one"}}, "[one]"},
		{"PrintArr empty", args{[]string{}}, "[]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintArr(tt.args.arr); got != tt.want {
				t.Errorf("PrintArr() = %v, want %v", got, tt.want)
			}
		})
	}
}
