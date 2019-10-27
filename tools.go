package cqrs

import "reflect"

func typeOf(t interface{}) string {
	return reflect.TypeOf(t).Elem().Name()
}
