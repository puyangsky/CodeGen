package codegen

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
)

func GenCode(configPath string, gentype string, output string) error {

	//---------------------generate by yaml-------------------//

	if gentype == "yaml" {
		class := Class{}

		data, err := ReadFile(configPath)
		if err != nil {
			return err
		}

		fmt.Printf("-----\n[DEBUG] Config content:\n%s\n", string(data))

		err = yaml.Unmarshal(data, &class)
		if err != nil {
			return err
		}
	}

	return nil
}

func ReadFile(configPath string) ([]byte, error) {

	//TODO 考虑在一个配置文件中同时配置多个类的生成，用"---"分割

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	return data, nil
}