package main

import "uscan/es"

//func main() {
//
//	// 创建 RPC 配置
//	rpcConfig := &rpcclient.ConnConfig{
//		Host:         "127.0.0.1:8332",
//		User:         "username",
//		Pass:         "passowrd",
//		HTTPPostMode: true,
//		DisableTLS:   true,
//	}
//
//	// 连接到节点
//	client, err := rpcclient.New(rpcConfig, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 调用 RPC 方法
//	// ...
//	// 获取比特币地址的交易记录
//	addressTxs, err := client.ListTransactionsCount(address, 1000)
//	if err != nil {
//		log.Fatal(err)
//	}
//	// 获取比特币地址的余额
//	addressBalance, err := client.GetBalance(address)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func main() {
	//chain.GenData()
	es.Init()
}
