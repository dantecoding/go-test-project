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

func (event *Event) GetEvent(db *sql.DB) error {
	query := fmt.Sprintf("SELECT * FROM events WHERE id=%d", event.ID)

	return db.QueryRow(query).Scan(&event)
}

func (event *Event) DeleteEvent(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM events WHERE id=%d", event.ID)
	_, err := db.Exec(statement)

	return err
}

func (event *Event) CreateEvent(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO events(name, published) VALUES('%s', '%d')", event.Name, event.isPublished())
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&event.ID)

	if err != nil {
		return err
	}

	return nil
}

func (event *Event) UpdateEvent(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE events SET name='%s', published=%d WHERE id=%d", event.Name, event.isPublished(), event.ID)
	_, err := db.Exec(statement)

	return err
}

func (event *Event) SetPublished(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE events SET published=%d WHERE id=%d", event.isPublished(), event.ID)
	_, err := db.Exec(statement)

	return err
}

func (event *Event) isPublished() int {
	published := 0
	if event.Published {
		published = 1
	}

	return published
}
