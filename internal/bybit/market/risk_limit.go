package market

import (
	"context"
	"fmt"

	bybit "github.com/bybit-exchange/bybit.go.api"
)

func GetRiskLimit(client *bybit.Client) {
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetMarketRiskLimits(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(serverResult))
}
