package codegen

//类
type Class struct {
	ClassName 	string		`yaml:"cname"`
	Type		string		`yaml:"type"`				//类型: Controller/Dao/Service/Model
	ClassType	string		`yaml:"ctype"`				//类的类型: class/interface
	AccessMode	string		`yaml:"mode"`				//类的访问类型: public/private/protected/static/final
	Methods 	[]Method	`yaml:"method"`
	Variables	[]Variable	`yaml:"v"`
}

//成员变量
type Variable struct {
	VariableName 	string		`yaml:"vname"`
	VariableType 	string		`yaml:"vtype"`
	AccessMode		string		`yaml:"vmode"`
	VariableValue	interface{}	`yaml:"vvalue"`
}

//方法
type Method struct {
	MethodName 	string
	MethodType	string
	AccessMode	string
	Return		Variable
}
