package main

import (
	"flag"
	"fortaFastRegister/interact"
)

func main() {
	var fileAddr string
	var configAddr string
	flag.StringVar(&fileAddr, "f", "", "批量部署的txt文件路径")
	flag.StringVar(&configAddr, "c", "", "配置文件路径")
	interact.ReadConfig(configAddr)
	if fileAddr != "" {
		interact.DepositFromFile(fileAddr)
	} else {
		interact.DepositFromFile("addr.txt")
	}
}
