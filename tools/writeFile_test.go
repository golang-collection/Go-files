package tools

import (
	"log"
	"testing"
)

/**
* @Author: super
* @Date: 2020-08-20 20:44
* @Description:
**/

func TestWriteFile(t *testing.T) {
	err := WriteFile("../test.txt")
	if err != nil {
		log.Println(err)
	}
}


func TestWriteFileWithOS(t *testing.T) {
	err := WriteFileWithOS("../test1.txt")
	if err != nil {
		log.Println(err)
	}
}

func TestWriteFileWithBufio(t *testing.T) {
	err := WriteFileWithBufio("../test2.txt")
	if err != nil {
		log.Println(err)
	}
}