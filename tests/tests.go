// Package tests contains supporting code for running tests.
package tests

import (
	"fmt"
	"reflect"
	"runtime"
)

// Success and failure markers.
var (
	success = "\u2713"
	failed  = "\u2717"
	reset   = "\033[0m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	purple  = "\033[35m"
	cyan    = "\033[36m"
	gray    = "\033[37m"
	white   = "\033[97m"
)

func init() {
	if runtime.GOOS == "windows" {
		reset = ""
		red = ""
		green = ""
		yellow = ""
		blue = ""
		purple = ""
		cyan = ""
		gray = ""
		white = ""
	}
}

// InArray checks if a given value exists in a specific array and returns the index or -1 if not found
func InArray(val interface{}, array interface{}) int {
	values := reflect.ValueOf(array)

	if reflect.TypeOf(array).Kind() == reflect.Slice || values.Len() > 0 {
		for i := 0; i < values.Len(); i++ {
			if reflect.DeepEqual(val, values.Index(i).Interface()) {
				return i
			}
		}
	}

	return -1
}

// Success prints in green the given string
func Success(before string, s string) string {
	return before + fmt.Sprintf("%s%s\t%s%s", green, success, s, reset)
}

// Failed prints in red the given string
func Failed(before string, s string) string {
	return before + fmt.Sprintf("%s%s\\%s%s", red, failed, s, reset)
}

// Given just prints the "Given" statements on tests in purple
func Given(s string) string {
	return fmt.Sprintf("%s%s%s", purple, s, reset)
}

// Reset returns the code to reset color
func Reset() string {
	return reset
}

// Red returns the code to red color
func Red() string {
	return red
}

// Green returns the code to green color
func Green() string {
	return green
}

// Yellow returns the code to yellow color
func Yellow() string {
	return yellow
}

// Blue returns the code to blue color
func Blue() string {
	return blue
}

// Purple returns the code to purple color
func Purple() string {
	return purple
}

// Cyan returns the code to cyan color
func Cyan() string {
	return cyan
}

// Gray returns the code to gray color
func Gray() string {
	return gray
}

// White returns the code to white color
func White() string {
	return white
}
