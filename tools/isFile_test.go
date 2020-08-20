package tools

import "testing"

/**
* @Author: super
* @Date: 2020-08-20 11:29
* @Description:
**/

func TestIsFile(t *testing.T) {
	var filesTest = []struct{
		path string
		result bool
	}{
		{"../hello.txt", true},
		{"../tools", false},
		{"", false},
	}

	for _, tt := range filesTest{
		actual := IsFile(tt.path)
		if actual != tt.result{
			t.Errorf("path: %s, exist: %v, actual: %v", tt.path, tt.result, actual)
		}
	}
}