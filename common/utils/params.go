package utils

import (
	"reflect"

	"github.com/1046102779/base/pkg/errors"
	"github.com/1046102779/base/pkg/log"
)

func CheckAndAssignParams(args, dest map[string]interface{}) (err error) {
	if args == nil {
		return
	}
	var (
		argValue interface{}
		ok       bool
	)
	for key, destValue := range dest {
		if argValue, ok = args[key]; !ok {
			log.ErrorRaw("[CheckAndAssignParams] param `%s` empty. ", key)
			return errors.FGEInvalidRequestParam
		}
		value := reflect.ValueOf(destValue).Elem()
		switch value.Kind() {
		case reflect.Int, reflect.Int16:
			if value.CanSet() {
				tmp := reflect.ValueOf(argValue).Int()
				if tmp <= 0 {
					log.ErrorRaw("[CheckAndAssignParams] param `%s` empty.", key)
					return errors.FGEInvalidRequestParam
				}
				value.SetInt(tmp)
			}
		case reflect.Float32, reflect.Float64:
			if value.CanSet() {
				tmp := reflect.ValueOf(argValue).Float()
				if int(tmp) <= 0 {
					log.ErrorRaw("[CheckAndAssignParams] param `%s` empty.", key)
					return errors.FGEInvalidRequestParam
				}
				value.SetFloat(tmp)
			}
		case reflect.String:
			if value.CanSet() {
				tmp := reflect.ValueOf(argValue).String()
				if tmp == "" {
					log.ErrorRaw("[CheckAndAssignParams] param `%s` empty.", key)
					return errors.FGEInvalidRequestParam
				}
				value.SetString(tmp)
			}
		}
	}
	return
}
