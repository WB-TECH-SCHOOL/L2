package main

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"http/internal/delivery/docs"
	"http/internal/delivery/routers"
	"log"
	"net/http"
)

func main() {
	docs.SwaggerInfo.BasePath = "/"
	http.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	routers.InitRouting()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
