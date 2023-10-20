package chain

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Block struct {
	Index     int    `json:"index"`
	Timestamp string `json:"timestamp"`
	Data      string `json:"data"`
	Hash      string `json:"hash"`
	PrevHash  string `json:"prevHash"`
}

var blockchain []Block

func GenData() {
	// 初始化区块链数据（模拟数据）
	blockchain = append(blockchain, Block{Index: 0, Timestamp: "2023-01-01", Data: "Genesis Block", Hash: "0", PrevHash: ""})
	blockchain = append(blockchain, Block{Index: 1, Timestamp: "2023-01-02", Data: "Transaction 1", Hash: "hash1", PrevHash: "0"})
	blockchain = append(blockchain, Block{Index: 2, Timestamp: "2023-01-03", Data: "Transaction 2", Hash: "hash2", PrevHash: "hash1"})

	// 创建HTTP路由
	http.HandleFunc("/blocks", GetBlocks)

	// 启动HTTP服务器
	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetBlocks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(blockchain)
	if err != nil {
		return
	}
}
