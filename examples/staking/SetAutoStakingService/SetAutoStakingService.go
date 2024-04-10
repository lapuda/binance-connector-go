package main

import (
	"context"
	"fmt"

	binance_connector "github.com/lapuda/binance-connector-go"
)

func main() {
	SetAutoStaking()
}

func SetAutoStaking() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	setAutoStaking, err := client.NewSetAutoStakingService().Product("STAKING").PositionId("123").Renewable("true").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(setAutoStaking))
}
