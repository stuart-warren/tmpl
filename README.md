tmpl
====

``tmpl`` is a basic cli tool to take a golang template file and output the processed result to stdout.

Some simple utility functions are provided
* env - access to current environment variables
* eq - use with if to compare values for equality
* ne - opposite of eq
* parseURL - take a string URL and return a [url.URL](https://golang.org/pkg/net/url/#URL) type
* split - take a string and split into a slice
* replace - take a string and replace a substring with another a number of times (-1 for no limit)

Usage
-----

```
$ cat test.template | tmpl > myconfig.cfg
$ # or
$ tmpl test.template > myconfig.cfg
```

See ``test.template`` for examples

Why?
----

This was mostly created to replace ``erb`` in my workflow and remove the need to install ``ruby`` just to insert environment variables into config files (and easier that using ``sed``)

How
---

Use [gox](https://github.com/mitchellh/gox) to cross compile for multiple platforms

```
gox -osarch="linux/amd64 windows/amd64 darwin/amd64"
```
