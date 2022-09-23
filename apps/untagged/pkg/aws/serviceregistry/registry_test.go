package serviceregistry

import (
	"reflect"
	"testing"
)

type TestEc2 struct{}

//
func TestRegisterType(t *testing.T) {

	stringName := "TestEc2"
	// register dummy type
	RegisterType((*TestEc2)(nil))

	// get all from registry
	allRegistered := Registry()
	length := len(allRegistered)

	if length < 1 {
		t.Errorf("[RegisterType] incorrect length, expected >= 1, actual [%v]", length)
	}

	// check dummy found
	if !IsRegistered(stringName) {
		t.Error("[RegisterType] dummy type was not registered")
	}

	found := Instance(stringName)
	foundType := reflect.TypeOf(found).Name()
	if foundType != stringName {
		t.Errorf("[RegisterType] instance type name did not match, expected [%v], actual [%v]", stringName, foundType)
	}
}
