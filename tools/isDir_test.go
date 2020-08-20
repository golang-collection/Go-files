package tools

import "testing"

/**
* @Author: super
* @Date: 2020-08-20 11:27
* @Description:
**/

func TestIsDir(t *testing.T) {
	var filesTest = []struct{
		path string
		result bool
	}{
		{"../hello.txt", false},
		{"../tools", true},
		{"", false},
	}

	for _, tt := range filesTest{
		actual := IsDir(tt.path)
		if actual != tt.result{
			t.Errorf("path: %s, exist: %v, actual: %v", tt.path, tt.result, actual)
		}
	}
}