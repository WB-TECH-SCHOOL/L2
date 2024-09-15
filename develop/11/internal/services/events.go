package services

import (
	"fmt"
	"http/internal/converters"
	"http/internal/models/domain"
	"http/internal/models/dto"
	"sync"
	"time"
)

type eventService struct {
	events map[int]domain.Event
	nextID int
	mu     sync.RWMutex

	converter converters.EventConverter
}

func InitEventService() Events {
	return &eventService{
		events: make(map[int]domain.Event),
		nextID: 1,
		mu:     sync.RWMutex{},

		converter: converters.NewEventConverter(),
	}
}

func (e *eventService) CreateEvent(event dto.EventCreate) (int, error) {
	domainEvent := domain.Event{
		ID:     e.nextID,
		UserID: event.UserID,
		Date:   event.Date,
		Title:  event.Title,
	}

	e.mu.Lock()
	e.events[domainEvent.ID] = domainEvent
	e.nextID++
	e.mu.Unlock()

	return domainEvent.ID, nil
}

func (e *eventService) GetEventsForDay(date time.Time) []dto.Event {
	e.mu.RLock()
	defer e.mu.RUnlock()
	var result []domain.Event
	for _, event := range e.events {
		if event.Date.Year() == date.Year() && event.Date.YearDay() == date.YearDay() {
			result = append(result, event)
		}
	}
	return e.converter.EventsDomainToDTO(result)
}

func (e *eventService) GetEventsForWeek(date time.Time) []dto.Event {
	e.mu.RLock()
	defer e.mu.RUnlock()
	var result []domain.Event
	year, week := date.ISOWeek()
	for _, event := range e.events {
		eYear, eWeek := event.Date.ISOWeek()
		if eYear == year && eWeek == week {
			result = append(result, event)
		}
	}
	return e.converter.EventsDomainToDTO(result)
}

func (e *eventService) GetEventsForMonth(date time.Time) []dto.Event {
	e.mu.RLock()
	defer e.mu.RUnlock()
	var result []domain.Event
	for _, event := range e.events {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() {
			result = append(result, event)
		}
	}
	return e.converter.EventsDomainToDTO(result)
}

func (e *eventService) UpdateEvent(event dto.EventUpdate) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, exists := e.events[event.ID]; !exists {
		return fmt.Errorf("event not found")
	}
	e.events[event.ID] = e.converter.EventUpdateDTOToDomain(event)
	return nil
}

func (e *eventService) DeleteEvent(id int) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, exists := e.events[id]; !exists {
		return fmt.Errorf("event not found")
	}
	delete(e.events, id)
	return nil
}
