// Package a provides some simple assertions for tests using package testing.
package a

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func Assert(t *testing.T, args ...interface{}) {
	if ok, message := assert("Assert", args...); !ok {
		t.Fatal(message)
	}
}

func Check(t *testing.T, args ...interface{}) {
	if ok, message := assert("Check", args...); !ok {
		t.Error(message)
	}
}

type CheckerFunc func(args ...interface{}) (ok bool, message string)

func assert(fn string, args ...interface{}) (ok bool, message string) {
	if len(args) < 2 {
		return false, fmt.Sprintf(format("a.%s: too few arguments"), fn)
	}
	checker, ok := args[1].(CheckerFunc)
	if !ok {
		// Third argument because the first one is the *testing.T, not passed to assert.
		return false, fmt.Sprintf(format("a.%s: third argument not an a.CheckerFunc"), fn)
	}
	if ok, message := checker(append(args[:1], args[2:]...)...); !ok {
		return false, format(message)
	}
	return true, ""
}

func expectNArgs(n int, fn string, args []interface{}) (params []interface{}, message string, err error) {
	if len(args) < n {
		return nil, "", fmt.Errorf("a.%s: not enough arguments", fn)
	}
	if len(args) > n+1 {
		return nil, "", fmt.Errorf("a.%s: too many arguments", fn)
	}
	if len(args) == n+1 {
		s, ok := args[n].(string)
		if !ok {
			stringer, ok := args[n].(fmt.Stringer)
			if !ok {
				return nil, "", fmt.Errorf("a.%s: last argument passed not a string or fmt.Stringer: %v", fn, args[n])
			}
			s = stringer.String()
		}
		message = s
	}
	return args[:n], message, nil
}

func format(err string) string {
	return fmt.Sprintf("\r\t%s: %s", caller(), err)
}

func caller() string {
	for i := 0; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			return ""
		}
		dir, filename := filepath.Split(file)
		if strings.HasSuffix(dir, "github.com/cespare/a/") && filename != "a_test.go" {
			continue
		}
		return fmt.Sprintf("%s:%d", filename, line)
	}
}
