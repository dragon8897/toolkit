package toolkit

import "testing"

func TestVersion(t *testing.T) {
	Version()
}

func TestIf(t *testing.T) {
	if If(true, 1, 2).(int) != 1 {
		t.Error("wrong ans")
	}
	if If(false, 1, 2).(int) != 2 {
		t.Error("wrong ans")
	}
}
