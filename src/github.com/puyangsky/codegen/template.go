package codegen

import "fmt"

//类
type Class struct {
	ClassName 	string			`yaml:"name"`
	Type		string			`yaml:"type"`				//类型: Controller/Dao/Service/Model
	ClassType	string			`yaml:"ctype"`				//类的类型: class/interface
	AccessMode	string			`yaml:"mode"`				//类的访问类型: public/private/protected/static/final
	Methods 	[]Method		`yaml:"methods"`
	Variables	[]Variable		`yaml:"variables"`
}

//成员变量
type Variable struct {
	VariableName 	string		`yaml:"vname"`
	VariableType 	string		`yaml:"vtype"`				//成员变量的类型: String/int/double等
	AccessMode		string		`yaml:"vmode"`				//成员变量的访问类型: public/private/protected/static/final
	VariableValue	interface{}	`yaml:"vvalue"`
}

//参数
type Parameter struct {
	ParameterName	string		`yaml:"pname"`				//参数名称
	ParameterType 	string		`yaml:"ptype"`				//参数类型：String/int/double等
}


//方法
type Method struct {
	MethodName 	string			`yaml:"mname"`
	AccessMode	string			`yaml:"mmode"`			//方法的访问类型: public/private/protected/static/final/abstract
	Parameters 	[]Parameter		`yaml:"mparameters"`		//方法的参数数组
	Return		Parameter		`yaml:"mreturn"`		//方法的返回值
}

func (c *Class) toString() string {
	classStr := fmt.Sprintf("ClassName: %s\nType: %s\nClassType: %s\nAccessMode: %s\nMethods len: %d\n" +
		"Variables len: %d", c.ClassName, c.Type, c.ClassType, c.AccessMode, len(c.Methods), len(c.Variables))
	return classStr
}