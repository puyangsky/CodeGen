package codegen

import (
	"text/template"
	"os"
	"strings"
	"errors"
)

/*
注入数据，从模板生成代码
 */
func GenFromTempalte(dst string, class Class) error{

	tpl := class.Type + ".tpl"
	t, err := template.ParseFiles(tpl)
	CheckErr(err)

	var separator string
	if os.IsPathSeparator('\\') {  //前边的判断是否是系统的分隔符
		separator = "\\"
	} else {
		separator = "/"
	}

	// 确保dst不以'/'结尾
	path := dst + separator + strings.ToLower(class.Type)
	println("path:", path)
	// 判断文件夹是否存在，不存在则创建文件夹
	if exist, _ := PathExists(path); !exist {
		os.Mkdir(path, os.ModePerm)
	}
	// 判断文件是否存在，不存在则创建文件
	filePath := path + separator + class.ClassName + ".java"
	if exist, _ := PathExists(filePath); !exist {
		file, err := os.Create(filePath)
		err = t.Execute(file, class)
		CheckErr(err)
	}

	return errors.New("File already exists!")
}