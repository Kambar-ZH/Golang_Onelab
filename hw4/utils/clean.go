package utils

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
)

func RemoveCyrillicFromString(obj reflect.Value) {
	value := obj.Interface().(string)
	re, err := regexp.Compile(`[а-яА-Я]`)
	if err != nil {
		log.Fatal(err)
	}
	value = re.ReplaceAllString(value, "")
	obj.SetString(value)
}

func TraverseStruct(obj reflect.Value, funcOnFields func(string, reflect.Value)) {
	switch obj.Kind() {
	case reflect.Ptr:
		objValue := obj.Elem()
		TraverseStruct(objValue, funcOnFields)
	case reflect.Struct:
		typeOfT := obj.Type()
		for i := 0; i < obj.NumField(); i++ {
			field := obj.Field(i)
			if (field.Kind() != reflect.Struct) {
				funcOnFields(typeOfT.Field(i).Name, field)
			}
			TraverseStruct(field, funcOnFields)
		}
	case reflect.String:
		RemoveCyrillicFromString(obj)
	}
}

func RemoveCyrillicFromStruct(v interface{}) {
	funcOnFields := func(string, reflect.Value) {}
	obj := reflect.ValueOf(v)
	TraverseStruct(obj, funcOnFields)
}

func ShowStruct(v interface{}) string {
	var sb strings.Builder
	funcOnFields := func(fieldName string, obj reflect.Value) {
		info := fmt.Sprintf("_____START_____\nField: %s\nType:  %s\nValue: %v\n______END_______\n\n", fieldName, obj.Type(), obj.Interface())
		sb.WriteString(info)
	}
	obj := reflect.ValueOf(v)
	TraverseStruct(obj, funcOnFields)
	return sb.String()
}
