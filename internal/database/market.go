package database

import (
	"fmt"

	"diov.local/krasivo/internal/models"
)

func GetAllCurrencies() ([]models.Currency, error) {
	sql := `
        SELECT name, symbol
        FROM market
    `

	rows, err := pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("error querying tasks: %w", err)
	}
	defer rows.Close()

	var currencies []models.Currency
	for rows.Next() {
		var currency models.Currency
		err := rows.Scan(
			&currency.Name,
			&currency.Symbol,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning task row: %w", err)
		}
		currencies = append(currencies, currency)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating task rows: %w", err)
	}

	return currencies, nil
}
