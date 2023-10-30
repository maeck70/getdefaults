package getdefaults

import (
	"log"
	"strings"
)

// Boolean function returning whether the value of any type is in the list of options.
func In(value interface{}, opt []interface{}) bool {
	for _, i := range opt {
		if i == value {
			return true
		}
	}
	return false
}

// Boolean function returning whether the string value is in the list of options.
func InStr(value string, opt []string) bool {
	for _, i := range opt {
		if i == value {
			return true
		}
	}
	return false
}

// Return the string value if it is in the list of options (strings), otherwise return the default value.
func ValInStrOrDefault(value string, default_value string, opt []string) string {
	if InStr(value, opt) {
		return value
	}
	return default_value
}

// Return the value if it is in the list of options (any type), otherwise return the default value.
func ValInOrDefault(value interface{}, default_value interface{}, opt []interface{}) interface{} {
	if In(value, opt) {
		return value
	}
	return default_value
}

// Return the value if it is not in the list of options (any type), otherwise return the default value.
func ValNotInOrDefault(value interface{}, default_value interface{}, opt []interface{}) interface{} {
	if In(value, opt) {
		return default_value
	}
	return value
}

// Return the boolean value if the value represents a boolean value, otherwise return the default boolean value.
// The value is considered a boolean value if it is a string containing one of the following values (any type):
// true, t, yes, y, != 0
// false, f, no, n, 0
func BoolOrDefault(value interface{}, default_value bool) bool {
	switch v := value.(type) {
	case string:
		v = strings.ToLower(v)
		if InStr(v, []string{"true", "t", "yes", "y"}) {
			return true
		}
		if InStr(v, []string{"false", "f", "no", "n"}) {
			return false
		}
	case bool:
		return v
	case int:
		return v != 0
	default:
		log.Printf("BoolOrDefault: Cannot evaluate %T for boolean\n", v)
	}
	return default_value
}
