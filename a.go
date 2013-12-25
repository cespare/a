// Package a provides some simple assertions for tests using package testing.
package a

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"unicode"
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
	checker, ok := args[1].(func(...interface{}) (bool, string))
	if !ok {
		checker, ok = args[1].(CheckerFunc)
		if !ok {
			// Third argument because the first one is the *testing.T, not passed to assert.
			return false, fmt.Sprintf(format("a.%s: third argument not an a.CheckerFunc"), fn)
		}
	}
	if ok, message := checker(append(args[:1], args[2:]...)...); !ok {
		return false, format(message)
	}
	return true, ""
}

func expectNArgs(n int, args []interface{}) (params []interface{}, message string, err string) {
	if len(args) < n {
		return nil, "", fnPrefix("not enough arguments")
	}
	if len(args) > n+1 {
		return nil, "", fnPrefix("too many arguments")
	}
	if len(args) == n+1 {
		s, ok := args[n].(string)
		if !ok {
			stringer, ok := args[n].(fmt.Stringer)
			if !ok {
				return nil, "", fnPrefix(fmt.Sprintf("last argument passed not a string or fmt.Stringer: %v", args[n]))
			}
			s = stringer.String()
		}
		message = s
	}
	return args[:n], message, ""
}

func format(err string) string {
	return fmt.Sprintf("\r\t%s: %s", caller(), err)
}

func fnPrefix(format string, args ...interface{}) string {
	return fmt.Sprintf(publicFn()+": "+format, args...)
}

func isInternal(path string) bool {
	dir, filename := filepath.Split(path)
	return strings.HasSuffix(dir, "github.com/cespare/a/") && filename != "a_test.go"
}

// caller returns the file:lineno of the first caller up the stack that's not in package a.
func caller() string {
	for i := 0; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			return ""
		}
		if isInternal(file) {
			continue
		}
		return fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}
}

// publicFn traverses up the stack until it finds an exported function from this package and returns its name.
// This is useful to formulate error messages which have the right public function name in them from inside
// helper functions.
func publicFn() string {
	for i := 0; ; i++ {
		pc, file, _, ok := runtime.Caller(i)
		if !ok {
			return ""
		}
		if !isInternal(file) {
			return ""
		}
		name := filepath.Base(runtime.FuncForPC(pc).Name())
		parts := strings.Split(name, ".")
		if len(parts) != 2 || parts[1] == "" {
			return ""
		}
		if unicode.IsUpper([]rune(parts[1])[0]) {
			// Exported
			return name
		}
	}
}
