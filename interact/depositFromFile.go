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
		addressSub := string(addressRaw)[2:]
		addressSub = strings.Replace(addressSub, " ", "", -1)
		res := Deposit(addressSub)
		time.Sleep(1 * time.Second)
		_, _ = fmt.Fprintln(logsFile, res)
		//if err != nil {
		//	return
		//}
	}
}
