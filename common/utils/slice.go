package utils

import (
	"fmt"
	"reflect"
)

func ArrayToString(items interface{}, sep string) (dest string) {
	value := reflect.ValueOf(items)
	if value.Type().Kind() != reflect.Slice {
		return
	}
	for index := 0; index < value.Len(); index++ {
		if index == 0 {
			dest = fmt.Sprintf("%v", value.Index(index))
		} else {
			dest = fmt.Sprintf("%s%s%v", dest, sep, value.Index(index))
		}
	}
	return
}
