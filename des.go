package lbs_des

import (
	"github.com/robertkrimen/otto"
	"io/ioutil"
)

type DesLbs struct {
}

func NewDesLbs() *DesLbs {
	return &DesLbs{}
}

func (lbs *DesLbs) StrEnc(data, key1, key2, key3 string) (result string) {

	filePath := "./js/des.js"
	//先读入文件内容
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	vm := otto.New()

	_, err = vm.Run(string(bytes))
	if err != nil {
		panic(err)
	}

	value, err := vm.Call("strEnc", nil, data, key1, key2, key3)
	if err != nil {
		panic(err)
	}

	result = value.String()
	return
}

func (lbs *DesLbs) StrDec(data, key1, key2, key3 string) (result string) {
	filePath := "./js/des.js"
	//先读入文件内容
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	vm := otto.New()

	_, err = vm.Run(string(bytes))
	if err != nil {
		panic(err)
	}

	value, err := vm.Call("strDec", nil, data, key1, key2, key3)
	if err != nil {
		panic(err)
	}

	result = value.String()
	return
}
