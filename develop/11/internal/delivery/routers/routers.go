package routers

import (
	"http/internal/delivery/handlers"
	"http/internal/delivery/middleware"
	"http/internal/services"
	"net/http"
)

func InitRouting() {
	eventService := services.InitEventService()
	eventHandler := handlers.InitEventHandler(eventService)

	http.Handle("/create_event", middleware.LogRequests(http.HandlerFunc(eventHandler.CreateEvent)))
	http.Handle("/events_for_day", middleware.LogRequests(http.HandlerFunc(eventHandler.GetEventsForDay)))
	http.Handle("/events_for_week", middleware.LogRequests(http.HandlerFunc(eventHandler.GetEventsForWeek)))
	http.Handle("/events_for_month", middleware.LogRequests(http.HandlerFunc(eventHandler.GetEventsForMonth)))
	http.Handle("/update_event", middleware.LogRequests(http.HandlerFunc(eventHandler.UpdateEvent)))
	http.Handle("/delete_event", middleware.LogRequests(http.HandlerFunc(eventHandler.DeleteEvent)))
}
