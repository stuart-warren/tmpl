package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

var (
	versionFlag = flag.Bool("v", false, "prints current version")
	bytesIn     []byte
	err         error
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [filename|-]:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if *versionFlag {
		fmt.Printf("tmpl %s\n", version)
		os.Exit(0)
	}
	var name string
	if len(os.Args) > 1 {
		filename := os.Args[1]
		bytesIn, err = ioutil.ReadFile(filename)
		name = filename
	} else {
		bytesIn, err = ioutil.ReadAll(os.Stdin)
		name = "stdin"
	}
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf(string(bytesIn))
	tmpl, err := template.New(name).Funcs(template.FuncMap{
		"contains":     contains,        // contains(item map[string]string, key string) bool
		"split":        strings.Split,   // split(s, sep string) []string
		"join":         strings.Join,    // join(a []string, sep string) string
		"replace":      strings.Replace, // replace(s, old, new string, n int) string
		"parseURL":     parseURL,        // parseURL(rawurl string) *url.URL
		"default":      defaultStr,      // defaultStr(item string, defaultValue string) string
		"eq":           eq,              // eq(x, y interface{}) bool
		"ne":           ne,              // ne(x, y interface{}) bool
		"env":          env,             // env(key string) string
		"envAll":       envAll,          // envAll() map[string]string
		"filterPrefix": filterPrefix,    // filterPrefix(items map[string]string, prefix string) map[string]string
		"filterSuffix": filterSuffix,    // filterSuffix(items map[string]string, suffix string) map[string]string
	}).Parse(string(bytesIn))
	if err != nil {
		log.Fatal("Issue parsing template:", err)
	}
	err = tmpl.Execute(os.Stdout, &context{})
	if err != nil {
		log.Fatal("Issue executing template:", err)
	}
}
