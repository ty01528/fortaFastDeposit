package main

import (
	"flag"
	"fortaFastRegister/interact"
)

func main() {
	var fileAddr string
	flag.StringVar(&fileAddr, "f", "", "批量部署的文件地址")
	interact.ReadConfig()
	if fileAddr != "" {
		interact.DepositFromFile(fileAddr)
	} else {
		interact.DepositFromFile("addr.txt")
	}
}
