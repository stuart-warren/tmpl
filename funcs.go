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

func env(key string) string {
	return os.Getenv(key)
}

func envAll() map[string]string {
	env := make(map[string]string)
	for _, i := range os.Environ() {
		sep := strings.Index(i, "=")
		env[i[0:sep]] = i[sep+1:]
	}
	return env
}

func contains(items map[string]string, key string) bool {
	if _, ok := items[key]; ok {
		return true
	}
	return false
}

func filterPrefix(items map[string]string, prefix string) map[string]string {
	ret := make(map[string]string)
	for k, v := range items {
		if strings.HasPrefix(k, prefix) {
			ret[k] = v
		}
	}
	return ret
}

func filterSuffix(items map[string]string, suffix string) map[string]string {
	ret := make(map[string]string)
	for k, v := range items {
		if strings.HasSuffix(k, suffix) {
			ret[k] = v
		}
	}
	return ret
}

func parseURL(rawurl string) *url.URL {
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Fatalf("unable to parse url %s: %s", rawurl, err)
	}
	return u
}

func defaultStr(item string, defaultValue string) string {
	if item == "" {
		item = defaultValue
	}
	return item
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
