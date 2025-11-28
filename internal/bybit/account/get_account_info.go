package account

import (
	"context"
	"fmt"

	bybit "github.com/bybit-exchange/bybit.go.api"
)

func GetAccountInfo(client *bybit.Client) {
	accountResult, err := client.NewUtaBybitServiceNoParams().GetAccountInfo(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(accountResult))
}
