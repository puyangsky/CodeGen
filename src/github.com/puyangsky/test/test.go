package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
)

//var data = `
//a: Easy!
//b:
//  c: 2
//  d: [3, 4]
//`

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func getFiled(i interface{}) {
	t := reflect.TypeOf(i)

	for {
		if t.Kind() == reflect.Struct {
			for i := 0; i < t.NumField(); i++ {
				println(t.Field(i).Name)
			}
		}

		fmt.Printf("\n%-8v %v 个方法:\n", t, t.NumMethod())

		break
	}
}

func main() {
	t := T{}

	data, err := ioutil.ReadFile("/Users/imac/GoglandProjects/CodeGen/src/github.com/puyangsky/test/test.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	//err = yaml.Unmarshal([]byte(data), &t)
	err = yaml.Unmarshal(data, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))
	//
	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal(data, &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)
	//
	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))

	getFiled(t)

	//m1 := make(map[int]string)
	//
	//m1[1] = "a"
	//
	//
	////m1["a"] = "b"
	//fmt.Printf("%v\n",m1)
	//for k, v := range m1 {
	//	println(k, v)
	//	println("type:", reflect.TypeOf(k))
	//}
}