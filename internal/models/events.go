package models

import (
	"database/sql"
	"time"
)

type Event struct {
	ID          int
	Title       string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
	Created     time.Time
}

type EventModel struct {
	DB *sql.DB
}

func (m *EventModel) Insert(title, description, location string, dateTime time.Time, userId int) (int, error) {
	stmt := `
	INSERT INTO events(title, description, location, dateTime, userId )
	VALUES(?, ?, ?, ?, ?)
`

	result, err := m.DB.Exec(stmt, title, description, location, dateTime, userId)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *EventModel) GetAll() ([]Event, error) {
	query := `SELECT id, title, description, location, dateTime, userId	FROM events`
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Title, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}
	return events, nil
}
