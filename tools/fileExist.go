package tools

import "os"

/**
* @Author: super
* @Date: 2020-08-20 11:13
* @Description:
**/

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}