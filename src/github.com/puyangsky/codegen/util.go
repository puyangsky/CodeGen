package codegen

import (
	"os"
	//"strings"
)

//判断文件是否存在
func PathExists(filename string) (bool, error){
	//strings.Contains(filename, "/")
	_, err := os.Stat(filename)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
