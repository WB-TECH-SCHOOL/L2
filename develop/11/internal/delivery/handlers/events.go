package handlers

import (
	"encoding/json"
	"errors"
	"http/internal/delivery/responses"
	"http/internal/errs"
	"http/internal/models/dto"
	"http/internal/services"
	"net/http"
	"strconv"
	"time"
)

type EventHandler struct {
	service services.Events
}

func InitEventHandler(service services.Events) EventHandler {
	return EventHandler{service: service}
}

// CreateEvent
// @Summary Create Event
// @Tags event
// @Description Create a new event
// @Accept  json
// @Produce  json
// @Param input body dto.EventCreate true "Event data"
// @Success 200 {object} responses.SuccessResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 503 {object} responses.ErrorResponse
// @Router /create_event [post]
func (e EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event dto.EventCreate
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		responses.BadRequest(w, responses.ResponseBadBody)
		return
	}

	id, err := e.service.CreateEvent(event)
	if err != nil {
		responses.ServiceUnavailable(w, err.Error())
		return
	}

	responses.WriteSuccess(w, http.StatusOK, strconv.Itoa(id))
}

// GetEventsForDay
// @Summary Get Events for a Day
// @Tags event
// @Description Get events for a specific day
// @Produce  json
// @Param date query string true "Date in YYYY-MM-DD format"
// @Success 200 {object} responses.SuccessResponse{result=[]dto.Event}
// @Failure 400 {object} responses.ErrorResponse
// @Router /events_for_day [get]
func (e EventHandler) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		responses.BadRequest(w, responses.ResponseBadQuery)
		return
	}

	events := e.service.GetEventsForDay(date)
	responses.WriteSuccess(w, http.StatusOK, events)
}

// GetEventsForWeek
// @Summary Get Events for a Week
// @Tags event
// @Description Get events for a specific week
// @Produce  json
// @Param date query string true "Date in YYYY-MM-DD format"
// @Success 200 {object} responses.SuccessResponse{result=[]dto.Event}
// @Failure 400 {object} responses.ErrorResponse
// @Router /events_for_week [get]
func (e EventHandler) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		responses.BadRequest(w, responses.ResponseBadQuery)
		return
	}

	events := e.service.GetEventsForWeek(date)
	responses.WriteSuccess(w, http.StatusOK, events)
}

// GetEventsForMonth
// @Summary Get Events for a Month
// @Tags event
// @Description Get events for a specific month
// @Produce  json
// @Param date query string true "Date in YYYY-MM-DD format"
// @Success 200 {object} responses.SuccessResponse{result=[]dto.Event}
// @Failure 400 {object} responses.ErrorResponse
// @Router /events_for_month [get]
func (e EventHandler) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		responses.BadRequest(w, responses.ResponseBadQuery)
		return
	}

	events := e.service.GetEventsForMonth(date)
	responses.WriteSuccess(w, http.StatusOK, events)
}

// UpdateEvent
// @Summary Update Event
// @Tags event
// @Description Update an existing event
// @Accept  json
// @Produce  json
// @Param input body dto.EventUpdate true "Event data"
// @Success 200 {object} responses.SuccessResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 503 {object} responses.ErrorResponse
// @Router /update_event [put]
func (e EventHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var event dto.EventUpdate
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		responses.BadRequest(w, responses.ResponseBadBody)
		return
	}

	err = e.service.UpdateEvent(event)
	if err != nil {
		if errors.Is(err, errs.ErrorNotFound) {
			responses.NotFound(w)
			return
		}
		responses.ServiceUnavailable(w, err.Error())
		return
	}

	responses.WriteSuccess(w, http.StatusOK, responses.ResponseSuccess)
}

// DeleteEvent
// @Summary Delete Event
// @Tags event
// @Description Delete an event by ID
// @Produce  json
// @Param id query int true "Event ID"
// @Success 200 {object} responses.SuccessResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Router /delete_event [delete]
func (e EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		responses.BadRequest(w, responses.ResponseBadQuery)
		return
	}

	err = e.service.DeleteEvent(id)
	if err != nil {
		if errors.Is(err, errs.ErrorNotFound) {
			responses.NotFound(w)
			return
		}
		responses.ServiceUnavailable(w, err.Error())
		return
	}

	responses.WriteSuccess(w, http.StatusOK, responses.ResponseSuccess)
}
