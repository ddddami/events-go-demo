package models

import "time"

type Event struct {
	ID          int
	Title       string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
	Created     time.Time
}

var events = []Event{}

func (e Event) Save() {
	event := e
	events = append(events, event)
}

func GetAllEvents() []Event {
	return events
}
