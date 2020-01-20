package main

import (
	"developUtil/ops"
	"flag"
	"fmt"
	"os"
)

var version = "1.0.0-snapshot"
var name = "devutil"
var (
	help       bool
	optainType string
)

func init() {
	flag.BoolVar(&help, "h", false, "this help")
	flag.StringVar(&optainType, "o", "", ops.OpsUsage())
	// 改变默认的 Usage
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, name+` version: `+version+`
Usage: devutil [-o 操作类型] [-s signal] [-c filename] [-p prefix] [-g directives]

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
