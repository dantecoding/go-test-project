package models

import (
	"database/sql"
	"fmt"
)

type Event struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Published bool   `json:"published"`
}

func (list *Event) GetListener(db *sql.DB) error {
	query := fmt.Sprintf("SELECT * FROM events WHERE id=%d", list.ID)

	return db.QueryRow(query).Scan(&list)
}
