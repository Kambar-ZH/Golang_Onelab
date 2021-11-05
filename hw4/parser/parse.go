package parser

import (
	"encoding/json"
	"hw4/parser/models"
	"reflect"
	"strconv"
)

func ParseFields(obj reflect.Value, data map[string]interface{}, parseTag string) error {
	typeOf := obj.Type()
	for i := 0; i < typeOf.NumField(); i++ {
		field := obj.Field(i)
		tagName := typeOf.Field(i).Tag.Get(parseTag)
		if value, ok := data[tagName]; ok {
			newValue := reflect.ValueOf(value)
			switch field.Kind() {
			case reflect.Int, reflect.Int64, reflect.Int32:
				switch newValue.Kind() {
				case reflect.String:
					num, err := strconv.Atoi(newValue.String())
					if err != nil {
						return err
					}
					field.SetInt(int64(num))
				case reflect.Float32, reflect.Float64:
					field.SetInt(int64(newValue.Float()))
				case reflect.Int, reflect.Int32, reflect.Int64:
					field.SetInt(newValue.Int())

				}
			case reflect.Slice:
				newSlice := reflect.MakeSlice(reflect.TypeOf([]models.Author{}), newValue.Len(), newValue.Len())
				for j := 0; j < newValue.Len(); j++ {
					if err := ParseFields(newSlice.Index(j), newValue.Index(j).Interface().(map[string]interface{}), parseTag); err != nil {
						return err
					}
					field.Set(newSlice)
				}
			case reflect.Struct:
				if err := ParseFields(field, newValue.Interface().(map[string]interface{}), parseTag); err != nil {
					return err
				}
			default:
				field.Set(newValue)
			}
		}
	}
	return nil
}

func ParseUsers(rawJson []byte, parseTag string) ([]*models.User, error) {
	var rawData []map[string]interface{}
	if err := json.Unmarshal(rawJson, &rawData); err != nil {
		return nil, err
	}
	var users []*models.User
	for _, data := range rawData {
		user := &models.User{}
		obj := reflect.ValueOf(user).Elem()
		if err := ParseFields(obj, data, parseTag); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func ParsePosts(rawXml []byte, parseTag string) ([]*models.Post, error) {
	var rawData []map[string]interface{}
	if err := json.Unmarshal(rawXml, &rawData); err != nil {
		return nil, err
	}
	var posts []*models.Post
	for _, data := range rawData {
		post := &models.Post{}
		obj := reflect.ValueOf(post).Elem()
		if err := ParseFields(obj, data, parseTag); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}