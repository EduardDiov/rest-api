package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"diov.local/krasivo/components"
	"diov.local/krasivo/internal/database"
	"diov.local/krasivo/internal/models"
)

func Search(w http.ResponseWriter, r *http.Request) {
	currencies, err := database.GetAllCurrencies()
	if err != nil {
		fmt.Println(err)
	}

	query := strings.ToLower(r.URL.Query().Get("search"))
	var results []models.Currency

	for _, currency := range currencies {
		if strings.Contains(strings.ToLower(currency.Name), query) {
			results = append(results, currency)
		}
	}

	components.CurrenciesList(results).Render(r.Context(), w)
}
