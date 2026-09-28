package api

import (
	"net/http"

	"github.com/jackc/pgx/v5"
)

func Router(db *pgx.Conn) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/health", Health)
	mux.Handle("/api/sales", Sales(db))

	return mux
}
