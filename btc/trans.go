package btc

import (
	"encoding/json"
	"github.com/btcsuite/btcd/rpcclient"
)

type btcTransaction struct {
	Txid     string `json:"txid"`
	Version  int    `json:"version"`
	LockTime int    `json:"locktime"`
	Size     int    `json:"size"`
	Vin      []struct {
		Txid      string `json:"txid"`
		Vout      int    `json:"vout"`
		ScriptSig struct {
			Asm string `json:"asm"`
			Hex string `json:"hex"`
		} `json:"scriptSig"`
		Sequence int `json:"sequence"`
	} `json:"vin"`
	Vout []struct {
		Value        float64 `json:"value"`
		N            int     `json:"n"`
		ScriptPubKey struct {
			Asm       string   `json:"asm"`
			Hex       string   `json:"hex"`
			ReqSigs   int      `json:"reqSigs"`
			Type      string   `json:"type"`
			Addresses []string `json:"addresses"`
		} `json:"scriptPubKey"`
	} `json:"vout"`
}

func getTransaction(client *rpcclient.Client, txid string) (*btcTransaction, error) {
	transactionJSON, err := client.GetRawTransactionVerbose(txid)
	if err != nil {
		return nil, err
	}
	var transaction btcTransaction
	err = json.Unmarshal([]byte(transactionJSON), &transaction)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}
