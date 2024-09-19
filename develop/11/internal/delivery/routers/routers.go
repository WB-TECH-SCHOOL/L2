package routers

import (
	"http/internal/delivery/handlers"
	"http/internal/services"
	"net/http"
)

func InitRouting() {
	eventService := services.InitEventService()
	eventHandler := handlers.InitEventHandler(eventService)

	http.HandleFunc("/create_event", eventHandler.CreateEvent)
	http.HandleFunc("/events_for_day", eventHandler.GetEventsForDay)
	http.HandleFunc("/events_for_week", eventHandler.GetEventsForWeek)
	http.HandleFunc("/events_for_month", eventHandler.GetEventsForMonth)
	http.HandleFunc("/update_event", eventHandler.UpdateEvent)
	http.HandleFunc("/delete_event", eventHandler.DeleteEvent)
}
