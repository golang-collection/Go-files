package tools

import (
	"errors"
	"io/ioutil"
	"os"
)

/**
* @Author: super
* @Date: 2020-08-20 12:02
* @Description:
**/

func ReadMinFile(path string)(string, error){
	if !IsFile(path){
		return "", errors.New("this path isn't a file")
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func ReadMinFileAll(path string)(string, error){
	if !IsFile(path){
		return "", errors.New("this path isn't a file")
	}
	file, err := os.Open(path)
	if err !=nil{
		return "", err
	}
	bytes, err := ioutil.ReadAll(file)
	if err !=nil{
		return "", err
	}
	return string(bytes), nil
}