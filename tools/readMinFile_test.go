package tools

import (
	"fmt"
	"testing"
)

/**
* @Author: super
* @Date: 2020-08-20 12:04
* @Description:
**/

func TestReadMinFile(t *testing.T) {
	str, err := ReadMinFile("../hello.txt")
	if err != nil{
		t.Log(err)
	}
	fmt.Println(str)
	fmt.Println(len(str))
}

func TestReadMinFileAll(t *testing.T) {
	str, err := ReadMinFileAll("../hello.txt")
	if err != nil{
		t.Log(err)
	}
	fmt.Println(str)
	fmt.Println(len(str))
}