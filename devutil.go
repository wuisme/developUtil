package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "1.0.0-snapshot"
var name = "devutil"
var (
	help       bool
	optainType string
	method     string
)

func init() {
	flag.BoolVar(&help, "h", false, "this help")
	flag.StringVar(&optainType, "t", "", opsUsage())
	flag.StringVar(&method, "m", "", "获取详细内容请")
	// 改变默认的 Usage
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, name+` version: `+version+`
Usage: devutil [-t 操作类型] [-m 操作类型下的方法]

Options:
`)
	flag.PrintDefaults()
}
func main() {
	flag.Parse()

	if help {
		flag.Usage()
	}
	switch optainType {
	case "ops":
		fmt.Println("获取到ops命令")
	default:
		flag.Usage()
	}
}
func opsUsage() string {
	returnStr := `操作类型，暂时支持为:
ops 运维类指令`
	return returnStr
}
