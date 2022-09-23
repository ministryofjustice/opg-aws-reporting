package serviceregistry

import (
	"reflect"
)

// method to register the structs we're creating to see if
// an aws service can be handled
var typeRegistry = make(map[string]reflect.Type)

// RegisterType is helper method to track each type of struct
// mapped to aws services
func RegisterType(typed interface{}) {
	t := reflect.TypeOf(typed).Elem()
	typeRegistry[t.Name()] = t
}

func RegisterTypes(typedList []interface{}) {
	for _, typed := range typedList {
		RegisterType(typed)
	}
}

// IsRegistered simply checks if the typeName is in the
// registry
func IsRegistered(typeName string) bool {
	_, exists := typeRegistry[typeName]
	return exists
}

// Registry simply returns the map
func Registry() map[string]reflect.Type {
	return typeRegistry
}

func Instance(typeName string) interface{} {
	return reflect.New(typeRegistry[typeName]).Elem().Interface()
}
