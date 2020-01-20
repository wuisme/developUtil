package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "1.0.0-snapshot"
var name = "devutil"
var (
	help bool
)

func init() {
	flag.BoolVar(&help, "h", false, "this help")
	// 改变默认的 Usage
	flag.Usage = usage
}
func usage() {
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0
Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]

Options:
`)
	flag.PrintDefaults()
}
func main() {
	flag.Parse()

	if help {
		flag.Usage()
	}
}
