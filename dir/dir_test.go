package dir

import "testing"

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
