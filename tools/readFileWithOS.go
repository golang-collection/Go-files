package tools

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

/**
* @Author: super
* @Date: 2020-08-20 20:24
* @Description:
**/

func ReadFileWithOS(path string) (string, error){
	if !IsFile(path){
		return "", errors.New("this path is not a file")
	}
	f, err := os.Open(path)
	if err != nil{
		return "", err
	}
	defer f.Close()

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	//seek(偏移量，起始位置)
	o2, err := f.Seek(6, 0)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	_, err = f.Seek(0, 0)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	fmt.Printf("5 bytes: %s\n", string(b4))

	return "", nil

}