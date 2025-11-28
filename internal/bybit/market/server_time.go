package market

import (
	"context"
	"fmt"

	bybit "github.com/bybit-exchange/bybit.go.api"
)

func GetServerTime(client *bybit.Client) {
	serverResult, err := client.NewUtaBybitServiceNoParams().GetServerTime(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bybit.PrettyPrint(serverResult))
}
