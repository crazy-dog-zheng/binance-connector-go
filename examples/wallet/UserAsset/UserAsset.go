package main

import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

func main() {
	UserAsset()
}

func UserAsset() {
	apiKey := ""
	secretKey := ""
	baseURL := "https://api4.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// UserAssetService - /sapi/v3/asset/getUserAsset
	//userAsset, err := client.NewUserAssetService().Asset("BTC").Do(context.Background())
	allCoinsInfo, err := client.NewGetAllCoinsInfoService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(allCoinsInfo))
}
