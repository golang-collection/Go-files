package tools

import (
	"fmt"
	"testing"
)

/**
* @Author: super
* @Date: 2020-08-20 11:49
* @Description:
**/

func TestReadFile(t *testing.T) {
	strings, err := ReadFile("../hello.txt")
	if err != nil{
		t.Log(err)
	}
	fmt.Println(strings)
	fmt.Println(len(strings))
}