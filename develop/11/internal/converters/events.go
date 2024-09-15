package converters

import (
	"http/internal/models/domain"
	"http/internal/models/dto"
)

type EventConverter interface {
	EventDomainToDTO(event domain.Event) dto.Event
	EventsDomainToDTO(events []domain.Event) []dto.Event

	EventUpdateDTOToDomain(event dto.EventUpdate) domain.Event
}

type eventConverter struct{}

func NewEventConverter() EventConverter {
	return eventConverter{}
}

// DOMAIN -> DTO

func (e eventConverter) EventDomainToDTO(event domain.Event) dto.Event {
	return dto.Event{
		EventCreate: dto.EventCreate{
			UserID: event.UserID,
			Date:   event.Date,
			Title:  event.Title,
		},
		ID: event.ID,
	}
}

func (e eventConverter) EventsDomainToDTO(events []domain.Event) []dto.Event {
	result := make([]dto.Event, len(events))

	for i, event := range events {
		result[i] = e.EventDomainToDTO(event)
	}

	return result
}

// DTO -> DOMAIN

func (e eventConverter) EventUpdateDTOToDomain(event dto.EventUpdate) domain.Event {
	return domain.Event{
		UserID: event.UserID,
		Date:   event.Date,
		Title:  event.Title,
		ID:     event.ID,
	}
}
