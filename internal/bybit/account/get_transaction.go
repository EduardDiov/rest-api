package account

import (
	"context"
	"fmt"

	bybit "github.com/bybit-exchange/bybit.go.api"
)

func GetTransaction(client *bybit.Client) {
	params := map[string]interface{}{"accountType": "UNIFIED", "category": "linear"}
	accountResult, err := client.NewUtaBybitServiceWithParams(params).GetTransactionLog(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(accountResult))
}
