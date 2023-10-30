package getdefaults

import (
	"testing"
	"time"

	"github.com/maeck70/getdefaults"
)

func TestInStr(t *testing.T) {
	equals(t, true, getdefaults.InStr("2", []string{"Fail", "2", "Foo"}))
	equals(t, false, getdefaults.InStr("1", []string{"Fail", "2", "Foo"}))
	equals(t, false, getdefaults.InStr("1", []string{"a"}))
}

func TestIn(t *testing.T) {
	equals(t, true, getdefaults.In(2, []interface{}{"1", 2, 3, "4"}))
	equals(t, false, getdefaults.In(2, []interface{}{"1", "2", "3"}))
	equals(t, false, getdefaults.In(5, []interface{}{1, 2, 3, 4}))
}

func TestValInStrOrDefault(t *testing.T) {
	equals(t, "abc", getdefaults.ValInStrOrDefault("abc", "foo", []string{"aaa", "abc"}))
	equals(t, "foo", getdefaults.ValInStrOrDefault("abc", "foo", []string{"aaa", "bbb"}))
	equals(t, "foo", getdefaults.ValInStrOrDefault("", "foo", []string{"aaa", "bbb"}))
}

func TestValInOrDefault(t *testing.T) {
	equals(t, "abc", getdefaults.ValInOrDefault("abc", "foo", []interface{}{"aaa", "abc"}))
	equals(t, "foo", getdefaults.ValInOrDefault("abc", "foo", []interface{}{"aaa", "bbb"}))
	equals(t, "foo", getdefaults.ValInOrDefault("", "foo", []interface{}{"aaa", "bbb"}))

	equals(t, 123, getdefaults.ValInOrDefault(123, 200, []interface{}{100, 123, 150}))
	equals(t, 200, getdefaults.ValInOrDefault(0.123, 200, []interface{}{"aaa", "bbb"}))
	equals(t, 0.123, getdefaults.ValInOrDefault(0.123, 200, []interface{}{0.100, 0.123, 150}))
	equals(t, 200, getdefaults.ValInOrDefault(0, 200, []interface{}{"aaa", 1, 0.123, time.Now()}))
}

func TestNotValInOrDefault(t *testing.T) {
	equals(t, "foo", getdefaults.ValNotInOrDefault("abc", "foo", []interface{}{"aaa", "abc"}))
	equals(t, "abc", getdefaults.ValNotInOrDefault("abc", "foo", []interface{}{"aaa", "bbb"}))
	equals(t, "not", getdefaults.ValNotInOrDefault("not", "foo", []interface{}{"aaa", "bbb"}))

	equals(t, 200, getdefaults.ValNotInOrDefault(123, 200, []interface{}{100, 123, 150}))
	equals(t, 0.123, getdefaults.ValNotInOrDefault(0.123, 200, []interface{}{"aaa", "bbb"}))
	equals(t, 200, getdefaults.ValNotInOrDefault(0.123, 200, []interface{}{0.100, 0.123, 150}))
	equals(t, 0, getdefaults.ValNotInOrDefault(0, 200, []interface{}{"aaa", 1, 0.123, time.Now()}))
}

func TestBoolOrDefault(t *testing.T) {
	equals(t, false, getdefaults.BoolOrDefault("", false))
	equals(t, true, getdefaults.BoolOrDefault("true", false))
	equals(t, true, getdefaults.BoolOrDefault("tRUe", false))
	equals(t, false, getdefaults.BoolOrDefault(false, false))
	equals(t, true, getdefaults.BoolOrDefault(true, false))
}
