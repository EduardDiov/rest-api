package handlers

import (
	"fmt"
	"net/http"

	"diov.local/krasivo/components"
	"diov.local/krasivo/internal/database"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	currencies, err := database.GetAllCurrencies()
	if err != nil {
		fmt.Println(err)
	}

	component := components.Currencies(currencies)
	component.Render(r.Context(), w)
	// indexPage, err := template.ParseFiles(
	// 	util.GetExecutablePath()+"/public/layout.html",
	// 	util.GetExecutablePath()+"/public/index.html",
	// )

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = indexPage.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
