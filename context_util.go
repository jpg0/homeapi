package main

import (
	"reflect"
	"fmt"
	"strconv"
	"github.com/Sirupsen/logrus"
	"strings"
	"golang.org/x/net/html"
)

func Hydrate(ctx map[string]string, target interface{}) {
	v := reflect.ValueOf(target)

	logrus.Debugf("Hydrating %v: %+v", v.Kind(), target)

	for { //deref
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
			logrus.Debugf("Unwrapped %v: %+v", v.Kind(), v)
		} else {
			break
		}
	}

	if v.Kind() != reflect.Struct {
		panic(fmt.Sprintf("Cannot hydrate a %+v (%+v) (%v not Struct)", v, target, v.Kind()))
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)

		tag := field.Tag.Get("ctx")

		val, present := ctx[tag]

		if present {
			v.FieldByName(field.Name).Set(coerce(val, field.Type))
		}
	}
}

func coerce(val string, t reflect.Type) reflect.Value {

	switch t.Kind() {
	case reflect.String:
		return reflect.ValueOf(val).Convert(t)
	case reflect.Int64, reflect.Int:
		rv, e := strconv.ParseInt(val, 10, 0)

		if e != nil {
			logrus.Errorf("Failed to coerce %v to int64, setting to 0", val)
		}

		return reflect.ValueOf(rv).Convert(t)
	case reflect.Slice:
		unescaped := strings.Split(val, "|")

		for i := range unescaped {
			unescaped[i] = html.UnescapeString(unescaped[i])
		}

		return reflect.ValueOf(strings.Split(val, "&"))
	default:
		panic(fmt.Sprintf("Unknown kind: %v", t.Kind()))
	}
}

func Dehydrate(from interface{}) map[string]string {
	ctx := make(map[string]string)

	t := reflect.TypeOf(from).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := field.Tag.Get("ctx")
		val := reflect.ValueOf(from).Elem().FieldByName(field.Name).Interface()

		fmt.Printf("%v = %+v\n", field.Name, val)

		writeValue(tag, val, ctx)
	}

	return ctx
}

func writeValue(key string, value interface{}, m map[string]string) {
	var val string

	if isZero(reflect.ValueOf(value)) {
		val = DEFAULT_CONTEXT_VALUE
	} else if reflect.TypeOf(value).Kind() == reflect.Slice {
		asSlice := value.([]string)

		escaped := make([]string, len(asSlice))

		for i := range asSlice {
			escaped[i] = html.EscapeString(asSlice[i])
		}

		val = strings.Join(escaped, "&")

	} else {
		val = fmt.Sprintf("%v", value)
	}

	m[key] = val
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && isZero(v.Index(i))
		}
		return z
	case reflect.Struct:
		z := true
		for i := 0; i < v.NumField(); i++ {
			z = z && isZero(v.Field(i))
		}
		return z
	}
	// Compare other types directly:
	z := reflect.Zero(v.Type())
	return v.Interface() == z.Interface()
}