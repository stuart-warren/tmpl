package main

import (
	"log"
	"net/url"
	"os"
	"reflect"
	"strings"
)

type context struct {
}

func env() map[string]string {
	env := make(map[string]string)
	for _, i := range os.Environ() {
		sep := strings.Index(i, "=")
		env[i[0:sep]] = i[sep+1:]
	}
	return env
}

func contains(item map[string]string, key string) bool {
	if _, ok := item[key]; ok {
		return true
	}
	return false
}

func defaultValue(args ...interface{}) string {
	if len(args) == 0 {
		log.Fatalf("default called with no values")
		return ""
	}

	if len(args) > 0 {
		if args[0] != nil {
			return args[0].(string)
		}
	}

	if len(args) > 1 {
		if args[1] == nil {
			log.Fatalf("default called with nil default value")
			return ""
		}

		if _, ok := args[1].(string); !ok {
			log.Fatalf("default is not a string value. hint: surround it w/ double quotes")
			return ""
		}

		return args[1].(string)
	}

	log.Fatalf("default called with no default value")
	return ""
}

func parseURL(rawurl string) *url.URL {
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Fatalf("unable to parse url %s: %s", rawurl, err)
	}
	return u
}

func eq(x, y interface{}) bool {
	normalize := func(v interface{}) interface{} {
		vv := reflect.ValueOf(v)
		switch vv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return vv.Int()
		case reflect.Float32, reflect.Float64:
			return vv.Float()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return vv.Uint()
		default:
			return v
		}
	}
	x = normalize(x)
	y = normalize(y)
	return reflect.DeepEqual(x, y)
}

func ne(x, y interface{}) bool {
	return !eq(x, y)
}
