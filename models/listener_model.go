package models

import (
	"database/sql"
	"fmt"
)

type Listener struct {
	ID      int    `json:"id"`
	EventId int    `json:"event_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (listener *Listener) GetListener(db *sql.DB) error {
	query := fmt.Sprintf("SELECT id, event_id, name, address FROM listeners WHERE id=%d", listener.ID)

	return db.QueryRow(query).Scan(&listener.ID, &listener.EventId, &listener.Name, &listener.Address)
}

func (listener *Listener) DeleteListener(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM listeners WHERE id=%d", listener.ID)
	_, err := db.Exec(statement)

	return err
}

func (listener *Listener) CreateListener(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO listeners(event_id, name, address) VALUES('%d', '%s', '%s')", listener.EventId, listener.Name, listener.Address)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&listener.ID)

	if err != nil {
		return err
	}

	return nil
}

func (listener *Listener) UpdateListener(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE listeners SET name='%s', event_id=%d, address='%s' WHERE id=%d", listener.Name, listener.EventId, listener.Address, listener.ID)
	_, err := db.Exec(statement)

	return err
}
