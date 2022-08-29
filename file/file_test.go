package file

import "testing"

func TestExist(t *testing.T) {
	var tests = []struct {
		f string
		e bool
	}{
		{"../README.md", true},
		{"../README.md333", false},
	}
	for _, tt := range tests {
		if Exist(tt.f) != tt.e {
			t.Errorf("test failed: %v, %v", tt.f, tt.e)
		}
	}
}

func TestMD5(t *testing.T) {
	_, err := MD5("../README.md")
	if err != nil {
		t.Errorf("md5 err: %+v", err)
	}
}

func TestNameNoExt(t *testing.T) {
	var tests = []struct {
		f string
		n string
	}{
		{"../README.md", "../README"},
		{"../README.md333", "../README"},
		{"../README.txt", "../README"},
		{"../README.pdf", "../README"},
		{"../README", "../README"},
	}
	for _, tt := range tests {
		if NameNoExt(tt.f) != tt.n {
			t.Errorf("name no ext failed: %v, %v", tt.f, tt.n)
		}
	}
}
