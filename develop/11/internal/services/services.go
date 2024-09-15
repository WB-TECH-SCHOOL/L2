package services

import (
	"http/internal/models/dto"
	"time"
)

type Events interface {
	CreateEvent(event dto.EventCreate) (int, error)
	GetEventsForDay(date time.Time) []dto.Event
	GetEventsForWeek(date time.Time) []dto.Event
	GetEventsForMonth(date time.Time) []dto.Event
	UpdateEvent(event dto.EventUpdate) error
	DeleteEvent(id int) error
}
