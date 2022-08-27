package interact

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func DepositFromFile(addr string) {
	logsName := fmt.Sprint("logs", time.Now().UnixNano(), ".txt")
	logsFile, _ := os.OpenFile(logsName, os.O_RDWR|os.O_CREATE, os.ModeAppend)
	addrListFile, err := os.Open(addr)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer addrListFile.Close()

	addrList := bufio.NewReader(addrListFile)
	for {
		addressRaw, _, c := addrList.ReadLine()
		if c == io.EOF {
			break
		}
		//addressRaw := "0xe4bbaf48d0202dc4209b09c7b173cd7e9aee669c"
		addressSub := string(addressRaw)[2:]
		addressSub = strings.Replace(addressSub, " ", "", -1)
		//address = 0xB2174c3Cb47bC9E417908905B7F8D65d06f4140c
		res := Deposit(addressSub)
		time.Sleep(1 * time.Second)
		_, _ = fmt.Fprintln(logsFile, res)
		//if err != nil {
		//	return
		//}
	}
}
