package option

import (
	"fmt"
	"reflect"
)

func Merge[T any](values ...T) T {
	if len(values) == 0 {
		// zero value
		var ret T
		return ret
	} else {
		ret := values[0]
		for _, value := range values[1:] {
			update(&ret, value)
		}

		return ret
	}
}

func update[T any](one *T, two T) {
	ty := reflect.TypeOf(two)
	fields := reflect.VisibleFields(ty)

	for _, field := range fields {
		oneFieldValue := reflect.ValueOf(one).Elem().FieldByIndex(field.Index)
		twoFieldValue := reflect.ValueOf(two).FieldByIndex(field.Index)

		if !twoFieldValue.IsZero() && oneFieldValue.CanSet() {
			fmt.Println("setting field", field.Name, oneFieldValue.CanSet())
			oneFieldValue.Set(twoFieldValue)
		} else {
			fmt.Println("not setting field", field.Name, oneFieldValue.CanSet())
		}
	}
}
