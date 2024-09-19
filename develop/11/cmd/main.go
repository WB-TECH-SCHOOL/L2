package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"
	"http/internal/delivery/docs"
	"http/internal/delivery/routers"
	"net/http"
	"os"
)

func main() {
	logFile, err := os.OpenFile("requests.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("could not open log file")
	}
	defer logFile.Close()

	log.Logger = zerolog.New(logFile).With().Timestamp().Logger()

	docs.SwaggerInfo.BasePath = "/"
	http.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	routers.InitRouting()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(fmt.Sprintf("could not start server: %v\n", err))
	}
}
