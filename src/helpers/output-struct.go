package helpers

import (
	"fmt"
	"reflect"
)

func OutputStruct(s interface{}) {
	val := reflect.ValueOf(s)
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Name
		fieldValue := val.Field(i).Interface()
		fmt.Printf("%s: %v\n", fieldName, fieldValue)
	}
}
