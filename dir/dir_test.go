package dir

import (
	"strings"
	"testing"
)

func TestWalk(t *testing.T) {
	var tests = []struct {
		d string
		f int
	}{
		{".", 2},
		{"./dir", 0},
	}
	for _, tt := range tests {
		fs := Walk(tt.d)
		actual := len(fs)
		if actual != tt.f {
			t.Errorf("test failed: %v, %v with %v : %+v", tt.d, tt.f, actual, fs)
		}
	}
}

func TestIsDir(t *testing.T) {
	var tests = []struct {
		d string
		f bool
	}{
		{".", true},
		{"../file", true},
		{"./dir", false},
	}
	for _, tt := range tests {
		actual := IsDir(tt.d)
		if actual != tt.f {
			t.Errorf("test failed: %v, %v with %v", tt.d, tt.f, actual)
		}
	}
}

func TestList(t *testing.T) {
	fs := List(".", WithExt("go"))
	for _, f := range fs {
		if !strings.HasSuffix(f, ".go") {
			t.Errorf("%s not with ext .go", f)
		}
	}
	fs = List("..", WithNoExt("go"))
	for _, f := range fs {
		if strings.HasSuffix(f, ".go") {
			t.Errorf("%s with ext .go", f)
		}
	}
}
