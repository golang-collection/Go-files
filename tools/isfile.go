package tools

import "os"

/**
* @Author: super
* @Date: 2020-08-20 11:15
* @Description:
**/

// 判断所给路径是否为文件
func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}