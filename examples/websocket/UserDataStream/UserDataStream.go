package main

import (
	"context"
	"fmt"

	binance_connector "github.com/lapuda/binance-connector-go"
)

func main() {
	WsUserData()
}

func WsUserData() {
	apiKey := "PaCw4ghVt10x2XVKiFNLOcLsMfyxqa1yBu9zNyZgSrpsrE8FpQAD68bKZNmymVuJ"
	secretKey := "hEkHWCKWzsrq1VZZFWiJHfBEPTp4c5PpEZSBEeYKLWuZMF4z2wzLNDLG9X5DHtP4"
	baseURL := "https://testnet.binance.vision"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)
	client.Debug = true
	listenKey, err := client.NewCreateListenKeyService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	websocketStreamClient := binance_connector.NewWebsocketStreamClient(false, "wss://testnet.binance.vision")

	wsUserDataHandler := func(event *binance_connector.WsUserDataEvent) {
		fmt.Println(binance_connector.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, _, err := websocketStreamClient.WsUserDataServe(listenKey, wsUserDataHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneCh
}
