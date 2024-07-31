package main

import (
	"reflect"
)

func walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)

	for i := range val.NumField() {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
