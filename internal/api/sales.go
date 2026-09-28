package api

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type Sale struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Category    string  `json:"category"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	TotalAmount float64 `json:"total_amount"`
}

func Sales(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		rows, err := db.Query(r.Context(), `
			SELECT id, product_name, category, quantity, unit_price, total_amount
			FROM sales
			ORDER BY id
			LIMIT 2
		`)
		if err != nil {
			http.Error(w, "database error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var sales []Sale

		for rows.Next() {
			var sale Sale

			err := rows.Scan(
				&sale.ID,
				&sale.ProductName,
				&sale.Category,
				&sale.Quantity,
				&sale.UnitPrice,
				&sale.TotalAmount,
			)
			if err != nil {
				http.Error(w, "scan error", http.StatusInternalServerError)
				return
			}

			sales = append(sales, sale)
		}

		w.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		encoder.Encode(sales)
	}
}
