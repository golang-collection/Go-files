package tools

import (
	"bufio"
	"errors"
	"io"
	"os"
)

/**
* @Author: super
* @Date: 2020-08-20 11:49
* @Description:
**/

func ReadFile(path string) ([]string, error){
	if !IsFile(path){
		return nil, errors.New("this path isn't a file")
	}
	inputFile, inputError := os.Open(path)
	if inputError != nil {
		return nil, inputError
	}
	defer inputFile.Close()

	result := make([]string, 0)

	//默认缓冲区大小4096
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		//fmt.Printf("%s", inputString)
		result = append(result, inputString)
		if readerError == io.EOF {
			return result, nil
		}
	}
}
