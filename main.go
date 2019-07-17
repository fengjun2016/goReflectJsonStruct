package main

import (
	"reflect"
	"fmt"
	"errors"
	"log"
)

type User struct {
	Name string `field:"name" type:"varchar(50)"`
	Age int `field:"age" type:"int"`
}

type Job struct {
	User   //匿名字段
	Job   string `field:"job" type:"string"`
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
		f := structDataType.Field(i)
		if f.Anonymous {
			for j := 0; j < f.Type.NumField(); j++ {
				anonymousVar := structDataType.FieldByIndex([]int{i, j})
				log.Println(anonymousVar)
				log.Printf("%v v:", structDataValue.Field(i).Interface())
				data[structDataType.FieldByIndex([]int{i, j}).Tag.Get("field")] = structDataValue.Field(i).Interface()
			}
		} else {
			data[structDataType.Field(i).Tag.Get("field")] = structDataValue.Field(i).Interface()
		}
	}

	fmt.Println(data)
	return data, nil
}

func main() {
	var u User
	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%s: %s %s\n", f.Name, f.Tag.Get("field"), f.Tag.Get("type"))
	}

	StructToMap(u)

	var j Job
	j.Age = 24
	j.Name = "fengjun"
	j.Job = "softwareEngineer"
	StructToMap(j)
}