# forta合约批量交互

----

## 使用说明

- <font color="#ff0000">请自行检阅与构建代码</font>

- 目前提供 质押 与 解除质押 的操作

在 ```config.conf``` 下输入填入需要签名的私钥 ( 执行合约操作的私钥 ) 与 Polygon的RPC地址

在 ```addr.txt``` 中输入要质押的地址，每行一个

使用 ```go build main.go```编译成可执行程序 ```main.go```

使用  ```main.exe  -a 需要执行的操作 -f 地址文件路径 -c 配置文件路径``` 执行

- ```-a``` 操作类型，```stake``` 为质押，```unStake``` 为解除质押
- ```-f``` 地址文件路径
- ```-c``` 配置文件路径

配置文件示例：

```
contractOwnerPriKey = "=="
rpcClient           = "https://polygon-mainnet.infura.io/v3/=="
gasLimit            = 450000
```
第一行为私钥

第二行为RPC地址

第三行为gasLimit

## 为防止网络拥堵，默认1s提交一个！！！
### 运行结果示例
![image](https://user-images.githubusercontent.com/95566315/187038613-acf89e68-e974-4b06-84aa-3bb4ab4494fb.png)
