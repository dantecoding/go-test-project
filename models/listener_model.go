package models

import (
	"database/sql"
	"fmt"
)

type Listener struct {
	ID      int    `json:"id"`
	Event   string `json:"event"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (list *Listener) GetListener(db *sql.DB) error {
	query := fmt.Sprintf("SELECT id, event, name, address FROM listeners WHERE id=%d", list.ID)

	return db.QueryRow(query).Scan(&list.ID, &list.Event, &list.Name, &list.Address)
}

func (list *Listener) DeleteListener(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM listeners WHERE id=%d", list.ID)
	_, err := db.Exec(statement)

	return err
}

func (list *Listener) CreateListener(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO listeners(event, name, address) VALUES('%s', '%s', '%s')", list.Event, list.Name, list.Address)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&list.ID)

	if err != nil {
		return err
	}

	return nil
}

func (list *Listener) UpdateListener(db *sql.DB) {

}
