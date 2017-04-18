package codegen

import (
	"os"
	//"strings"
)

//判断文件和路径是否存在
func PathExists(filename string) (bool, error){
	_, err := os.Stat(filename)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
