package a

import (
	"fmt"
	"reflect"
)

// Slices:
// - Contains (slice contains values)
// Maps:
// - Contains (contains key/val pairs)
// - ContainsKeys (map contains keys)
// Numeric
// - Lt
// - Leq
// - Gt
// - Geq
// - Approx (n within 0.1% of x)
// - ApproxDelta (n equal to x +/- d)
// Strings
// - Contains (s contains substring)
// - Matches (s matches regex)


// DeepEquals is a checker which compares two values using reflect.DeepEquals.
func DeepEquals(args ...interface{}) (ok bool, message string) {
	params, msg, err := expectNArgs(2, args)
	if err != "" {
		return false, err
	}
	if reflect.DeepEqual(params[0], params[1]) {
		return true, ""
	}
	if msg == "" {
		return false, fnPrefix("expected %#v, but got %#v", params[1], params[0])
	}
	return false, msg
}

func Equals(args ...interface{}) (ok bool, message string) {
	params, msg, err := expectNArgs(2, args)
	if err != "" {
		return false, err
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
		return false, fnPrefix("expected %#v, but got %#v", params[1], params[0])
	}
	return false, msg
}

func IsTrue(args ...interface{}) (ok bool, message string) {
	params, msg, err := expectNArgs(1, args)
	if err != "" {
		return false, err
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
		return false, fnPrefix("expected true, but got %#v", params[0])
	}
	return false, msg
}

func IsFalse(args ...interface{}) (ok bool, message string) {
	params, msg, err := expectNArgs(1, args)
	if err != "" {
		return false, err
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
		return false, fnPrefix("expected false, but got %#v", params[0])
	}
	return false, msg
}

func IsNil(args ...interface{}) (ok bool, message string) {
	params, msg, err := expectNArgs(1, args)
	if err != "" {
		return false, err
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
		return false, fnPrefix("expected nil, but got %#v", params[0])
	}
	return false, msg
}
