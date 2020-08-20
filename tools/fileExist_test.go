package tools

import "testing"

/**
* @Author: super
* @Date: 2020-08-20 11:16
* @Description:
**/

func TestFileExists(t *testing.T) {
	var filesTest = []struct{
		path string
		exist bool
	}{
		{"../hello.txt", true},
		{"hello.txt", false},
	}

	for _, tt := range filesTest{
		actual := FileExists(tt.path)
		if actual != tt.exist{
			t.Errorf("path: %s, exist: %v, actual: %v", tt.path, tt.exist, actual)
		}
	}
}