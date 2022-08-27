package interact

import (
	"github.com/astaxie/beego/config"
	"log"
	"strconv"
)

var (
	contractAddress     = "0xd2863157539b1D11F39ce23fC4834B62082F6874"
	contractOwnerPriKey = "-"
	rpcClient           = "https://polygon-mainnet.infura.io/v3/-"
	chainId             = 137
	gasLimit            uint64
)

func ReadConfig(filename string) {
	conf, err := config.NewConfig("ini", "config.conf")
	if filename != "" {
		conf, err = config.NewConfig("ini", filename)
	}
	if err != nil {
		log.Print("config read error!")
		log.Println(err)
	}
	contractOwnerPriKey = conf.String("contractOwnerPriKey")
	rpcClient = conf.String("rpcClient")
	gasLimit, _ = strconv.ParseUint(conf.String("gasLimit"), 10, 64)
}
