package a

import (
	"fmt"
	"reflect"
)

var (
	DeepEquals = CheckerFunc(deepEquals)
	Equals     = CheckerFunc(equals)
	IsTrue     = CheckerFunc(istrue)
	IsFalse    = CheckerFunc(isfalse)
	IsNil      = CheckerFunc(isnil)
)

func deepEquals(args ...interface{}) (ok bool, message string) {
	params, msg, err := expectNArgs(2, "DeepEquals", args)
	if err != nil {
		return false, err.Error()
	}
	if reflect.DeepEqual(params[0], params[1]) {
		return true, ""
	}
	if msg == "" {
		return false, fmt.Sprintf("a.DeepEquals: expected %#v, but got %#v", params[1], params[0])
	}
	return false, msg
}

func equals(args ...interface{}) (ok bool, message string) {
	params, msg, err := expectNArgs(2, "Equals", args)
	if err != nil {
		return false, err.Error()
	}
	defer func() {
		if err := recover(); err != nil {
			ok = false
			message = fmt.Sprint(err)
		}
	}()
	if params[0] == params[1] {
		return true, ""
	}
	if msg == "" {
		return false, fmt.Sprintf("a.Equals: expected %#v, but got %#v", params[1], params[0])
	}
	return false, msg
}

func istrue(args ...interface{}) (ok bool, message string) {
	params, msg, err := expectNArgs(1, "IsTrue", args)
	if err != nil {
		return false, err.Error()
	}
	defer func() {
		if err := recover(); err != nil {
			ok = false
			message = fmt.Sprint(err)
		}
	}()
	if params[0] == true {
		return true, ""
	}
	if msg == "" {
		return false, fmt.Sprintf("a.IsTrue: expected true, but got %#v", params[0])
	}
	return false, msg
}

func isfalse(args ...interface{}) (ok bool, message string) {
	params, msg, err := expectNArgs(1, "IsFalse", args)
	if err != nil {
		return false, err.Error()
	}
	defer func() {
		if err := recover(); err != nil {
			ok = false
			message = fmt.Sprint(err)
		}
	}()
	if params[0] == false {
		return true, ""
	}
	if msg == "" {
		return false, fmt.Sprintf("a.IsFalse: expected false, but got %#v", params[0])
	}
	return false, msg
}

func isnil(args ...interface{}) (ok bool, message string) {
	params, msg, err := expectNArgs(1, "IsNil", args)
	if err != nil {
		return false, err.Error()
	}
	defer func() {
		if err := recover(); err != nil {
			ok = false
			message = fmt.Sprint(err)
		}
	}()
	if params[0] == nil {
		return true, ""
	}
	if msg == "" {
		return false, fmt.Sprintf("a.IsNil: expected nil, but got %#v", params[0])
	}
	return false, msg
}
