package main

import (
	"reflect"
	"fmt"
	"errors"
	"log"
)

type TestStruct struct {
	Name string `field:"name" type:"varchar(50)"`
	Age int `field:"age" type:"int"`
}

//将结构体转换成map 保留tag 里面的名称
func StructToMap(structData interface{}) (map[string]interface{}, error) {
	if structData == nil {
		return nil, errors.New("struct is empty.")
		log.Println("struct is empty")
	}
	data := make(map[string]interface{})
	structDataType := reflect.TypeOf(structData)
	structDataValue := reflect.ValueOf(structData)
	for i := 0; i < structDataType.NumField(); i++ {
		data[structDataType.Field(i).Tag.Get("field")] = structDataValue.Field(i).Interface()
	}
	fmt.Println(data)
	return data, nil
}

func main() {
	var u TestStruct
	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%s: %s %s\n", f.Name, f.Tag.Get("field"), f.Tag.Get("type"))
	}

	StructToMap(u)

}