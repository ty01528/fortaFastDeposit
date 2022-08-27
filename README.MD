# forta合约批量Deposit

----

## 使用说明

- 请自行构建代码，```Releases```下可执行文件<font color="#ff0000">仅提供了</font>根据文件地址进行质押的操作


在 ```config.conf``` 下输入填入需要签名的私钥 ( 执行合约操作的私钥 ) 与 Polygon的RPC地址

在 ```addr.txt``` 中输入要质押的地址，每行一个

使用 ```go build main.go```编译成可执行程序 ```main.go```

确保以上三个文件在同一目录下，双击```main.go```执行