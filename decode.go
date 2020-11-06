package valedictory

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func getType(v interface{}) (reflect.Value, error) {
	val := reflect.ValueOf(v)
	for val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return reflect.Value{}, fmt.Errorf("Value of %s should be struct, got %s", val, val.Kind())
	}
	return val, nil
}

func parseTag(tag string) (string, string) {
	s := strings.Split(tag, ",")
	options := map[string]string{}
	for _, t := range s[1:] {
		kv := strings.Split(t, ":")
		switch len(kv) {
		case 1:
			options[kv[0]] = ""
		case 2:
			options[kv[0]] = kv[1]
		}
	}
	return s[0], options["default"]
}

func Decode(m interface{}, v url.Values) error {
	val, err := getType(m)
	if err != nil {
		return err
	}

	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		sf := typ.Field(i)
		sv := val.Field(i)
		if !sv.CanSet() {
			continue
		}

		tag := sf.Tag.Get("valedictory")
		if tag == "-" || tag == "" {
			continue
		}

		name, dft := parseTag(tag)

		switch sf.Type.Kind() {
		case reflect.String:
			var stringv = dft
			if s, ok := v[name]; ok {
				stringv = s[0]
			}
			sv.SetString(stringv)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			var ints string = dft
			if s, ok := v[name]; ok {
				ints = s[0]
			}
			if i, err := strconv.Atoi(ints); err != nil {
				continue
			} else {
				sv.SetInt(int64(i))
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			var ints string = dft
			if s, ok := v[name]; ok {
				ints = s[0]
			}
			if i, err := strconv.Atoi(ints); err != nil {
				continue
			} else {
				sv.SetUint(uint64(i))
			}
		case reflect.Bool:
			var boolv = dft
			if s, ok := v[name]; ok {
				boolv = s[0]
			}
			if boolv == "true" {
				sv.SetBool(true)
			} else {
				sv.SetBool(false)
			}
		}
	}
	return nil
}
