package load

import (
	"fmt"
	"reflect"
	"strconv"
)

type Prefixable interface {
	Prefix() string
}

func ProcessField(value *reflect.Value, element reflect.Type, i int, key string, props map[string]any) {
	field := element.Field(i)
	fieldVal := value.Field(i)
	varType := field.Type.Kind()
	if varType == reflect.Struct {
		element := fieldVal.Addr().Elem()
		prefix := fieldVal.Interface().(Prefixable).Prefix()
		loadFields(&element, element.Type(), key+"."+prefix, props)
	} else {
		if newValue, ok := props[key]; ok {
			if varType == reflect.Bool {
				fieldVal.SetBool(newValue.(bool))
			} else if varType == reflect.String {
				fieldVal.SetString(newValue.(string))
			} else if varType == reflect.Float64 {
				fieldVal.SetFloat(newValue.(float64))
			} else if varType == reflect.Int64 {
				txtValue := fmt.Sprintf("%.0f", newValue.(float64))
				intValue, _ := strconv.ParseInt(txtValue, 10, 64)
				fieldVal.SetInt(intValue)
			}
		}
	}
}

func loadFields(valueOf *reflect.Value, typeOf reflect.Type, prefix string, props map[string]any) {
	if !valueOf.CanAddr() {
		fmt.Println("cannot assign to the item passed, item must be a pointer in order to assign")
		return
	}
	finalType := typeOf
	if typeOf.Kind() != reflect.Struct {
		finalType = typeOf.Elem()
	}

	if finalType.Kind() == reflect.Struct {
		for i := 0; i < finalType.NumField(); i++ {
			field := finalType.Field(i)
			key := prefix + "." + field.Tag.Get("config")
			ProcessField(valueOf, finalType, i, key, props)
		}
	}
}

func Properties(container interface{}, props map[string]any) {
	prefix := container.(Prefixable).Prefix()
	valueOf := reflect.ValueOf(container).Elem()
	typeOf := reflect.TypeOf(container)
	loadFields(&valueOf, typeOf, prefix, props)
}
