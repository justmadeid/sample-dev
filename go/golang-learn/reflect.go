package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type ContohLAgi struct {
	Name        string
	Description string
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf((data))
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "ture" {
			// versi ringkas
			// return reflect.ValueOf(data).Field(i).Interface() != ""
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"Made"}
	// var sampleType reflect.Type = reflect.TypeOf(sample)
	sampleType := reflect.TypeOf(sample)
	// structField := sampleType.Field(0)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	sample.Name = ""
	fmt.Println(IsValid(sample))

	fmt.Println("contoh lagi")
	contoh := ContohLAgi{"mama", "papa"}
	fmt.Println(IsValid(contoh))

}
