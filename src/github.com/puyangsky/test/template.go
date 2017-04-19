package main

import (
	"text/template"
	"os"
)

//模板引擎测试代码，嵌套struct+循环

func main()  {
	tpl := `
Name:   {{.Name}}
Methods:{{range $v := .Methods.MethodName}}
MethodName: {{$v}}{{end}}
	`
	names := []string{"a", "b"}

	data := struct {
		Name string
		Methods struct{
			MethodName []string
		}
	}{
		Name: "fuck",
		Methods: struct{ MethodName []string }{MethodName: names},
	}

	t, _ := template.New("test").Parse(tpl)
	t.Execute(os.Stdout, data)
}
