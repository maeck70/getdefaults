package getdefaults

import (
	"log"
	"strings"
)

func In(value interface{}, opt []interface{}) bool {
	for _, i := range opt {
		if i == value {
			return true
		}
	}
	return false
}

func InStr(value string, opt []string) bool {
	for _, i := range opt {
		if i == value {
			return true
		}
	}
	return false
}

func ValInStrOrDefault(value string, default_value string, opt []string) string {
	if InStr(value, opt) {
		return value
	}
	return default_value
}

func ValInOrDefault(value interface{}, default_value interface{}, opt []interface{}) interface{} {
	if In(value, opt) {
		return value
	}
	return default_value
}

func ValNotInOrDefault(value interface{}, default_value interface{}, opt []interface{}) interface{} {
	if In(value, opt) {
		return default_value
	}
	return value
}

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
		log.Fatalf("BoolOrDefault: Cannot evaluate %T for boolean", v)
	}
	return default_value
}
