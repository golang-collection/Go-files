package tools

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

/**
* @Author: super
* @Date: 2020-08-20 20:43
* @Description:
**/

func WriteFile(path string) error {
	d1 := []byte("hello\ngo\n")
	//没有文件自动创建
	err := ioutil.WriteFile(path, d1, 0644)
	return err
}

func WriteFileWithOS(path string) error {
	//通过os的常量确定打开文件的模式
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()


	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// Sync commits the current contents of the file to stable storage.
	// Typically, this means flushing the file system's in-memory copy
	// of recently written data to disk.
	f.Sync()

	return nil
}

func WriteFileWithBufio(path string) error{
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
	return nil
}
