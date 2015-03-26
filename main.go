package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	var (
		bytesIn []byte
		err     error
	)
	if len(os.Args) > 1 {
		filename := os.Args[1]
		bytesIn, err = ioutil.ReadFile(filename)
	} else {
		bytesIn, err = ioutil.ReadAll(os.Stdin)
	}
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf(string(bytesIn))
	tmpl := template.Must(template.New("stdin").Funcs(template.FuncMap{
		"contains": contains,        // contains(item map[string]string, key string) bool
		"split":    strings.Split,   // split(s, sep string) []string
		"replace":  strings.Replace, // replace(s, old, new string, n int) string
		"default":  defaultValue,    // default(args ...interface{}) string
		"parseURL": parseURL,        // parseURL(rawurl string) *url.URL
		"eq":       eq,              // eq(x, y interface{}) bool
		"ne":       ne,              // ne(x, y interface{}) bool
		"env":      env,             // env() map[string]string
	}).Parse(string(bytesIn)))
	err = tmpl.Execute(os.Stdout, &context{})
	if err != nil {
		log.Fatal(err)
	}
}
