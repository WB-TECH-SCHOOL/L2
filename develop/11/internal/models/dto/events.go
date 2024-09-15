package dto

import "time"

type EventCreate struct {
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
	Title  string    `json:"title"`
}

type Event struct {
	EventCreate
	ID int `json:"id"`
}

type EventUpdate struct {
	Event
}
