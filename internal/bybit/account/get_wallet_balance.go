package account

import (
	"context"
	"fmt"

	bybit "github.com/bybit-exchange/bybit.go.api"
)

func GetWalletBalance(client *bybit.Client) {
	params := map[string]interface{}{"accountType": "UNIFIED"}
	accountResult, err := client.NewUtaBybitServiceWithParams(params).GetFeeRates(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(accountResult))
}
