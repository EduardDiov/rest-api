package handlers

import (
	"fmt"
	"net/http"

	"diov.local/krasivo/components"
	"diov.local/krasivo/internal/database"
	"github.com/go-chi/chi"
)

func Currency(w http.ResponseWriter, r *http.Request) {
	symbol := chi.URLParam(r, "symbol")

	currencies, err := database.GetAllCurrencies()
	if err != nil {
		fmt.Println(err)
	}

	for _, currency := range currencies {
		if currency.Symbol == symbol {
			component := components.CurrencyDetail(currency)
			component.Render(r.Context(), w)
			return
		}
	}
	// btcPage, err := template.ParseFiles(util.GetExecutablePath() + "/public/crypto/btc.html")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = btcPage.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
