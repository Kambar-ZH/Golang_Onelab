package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

var rawJson = []byte(`[
  {
    "id": 1,
    "address": {
      "city_id": 5,
      "street": "Satbayev"
    },
    "age": 20
  },
  {
    "id": 1,
    "address": {
      "city_id": "6",
      "street": "Al-Farabi"
    },
    "age": "32"
  }
]`)

func Parse(obj reflect.Value, data map[string]interface{}, parseTag string) error {
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
			case reflect.Struct:
				if err := Parse(field, newValue.Interface().(map[string]interface{}), parseTag); err != nil {
					return err
				}
			default:
				field.Set(newValue)
			}
		}
	}
	return nil
}

func Solve() interface{} {
	var rawData []map[string]interface{}
	if err := json.Unmarshal(rawJson, &rawData); err != nil {
		panic(err)
	}
	var users []*User
	for _, data := range rawData {
		user := &User{}
		obj := reflect.ValueOf(user).Elem()
		Parse(obj, data, "json")
		users = append(users, user)
	}
	return users
}

type User struct {
	ID      int64   `json:"id"`
	Address Address `json:"address"`
	Age     int     `json:"age"`
}

func (u *User) String() string {
	return fmt.Sprintf("ID: %d, Address: %v, Age: %d.", u.ID, u.Address, u.Age)
}

type Address struct {
	CityID int64  `json:"city_id"`
	Street string `json:"street"`
}

func (a Address) String() string {
	return fmt.Sprintf("CityID: %d, Street: %s", a.CityID, a.Street)
}
