package main

import (
	"flag"
	"fortaFastRegister/interact"
)

func main() {
	var fileAddr string
	var configAddr string
	var act string
	flag.StringVar(&fileAddr, "f", "", "批量部署的txt文件路径")
	flag.StringVar(&configAddr, "c", "", "配置文件路径")
	flag.StringVar(&act, "a", "", "需要执行的操作: \n 质押： -a stake \n 解除质押： -a unStake ")
	interact.ReadConfig(configAddr)
	if act == "stake" {
		if fileAddr != "" {
			interact.DepositFromFile(fileAddr)
		} else {
			interact.DepositFromFile("addr.txt")
		}
	}
	if act == "unStake" {
		if fileAddr != "" {
			interact.UnStakeFromFile(fileAddr)
		} else {
			interact.UnStakeFromFile("addr.txt")
		}
	}
}
