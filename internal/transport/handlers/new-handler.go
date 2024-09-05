package handlers

import pgx "github.com/jackc/pgx/v5"

type Handler struct {
	DBconn *pgx.Conn
}

// Create new handler and previous reports info from file it needed
func NewHandler() Handler {
	var h Handler

	return h
}
