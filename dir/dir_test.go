package dir

import (
	"os"
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
	fs = List("..", WithReg(`.*\.go`))
	for _, f := range fs {
		if !strings.HasSuffix(f, ".go") {
			t.Errorf("%s not with ext .*\\.go", f)
		}
	}
}

func TestExist(t *testing.T) {
	var tests = []struct {
		d string
		e bool
	}{
		{
			".",
			true,
		},
		{
			"../file",
			true,
		},
		{
			"../test",
			false,
		},
	}
	for _, test := range tests {
		if Exist(test.d) != test.e {
			t.Errorf("Exist : dir %s ret %v", test.d, test.e)
		}
	}
}

func TestEnsureDir(t *testing.T) {
	var tests = []struct {
		d string
		c bool
	}{
		{
			"../dir",
			false,
		},
		{
			"../file",
			false,
		},
		{
			"../test",
			true,
		},
	}
	for _, test := range tests {
		c := EnsureDir(test.d)
		if c {
			os.Remove(test.d)
		}
		if c != test.c {
			t.Errorf("EnsureDir dir : %s created : %v", test.d, test.c)
		}
	}
}
