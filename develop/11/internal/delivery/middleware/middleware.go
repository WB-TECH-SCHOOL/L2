package middleware

import (
	"bytes"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	lrw.body.Write(b)
	return lrw.ResponseWriter.Write(b)
}

func LogRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Чтение тела запроса
		var requestBody bytes.Buffer
		if r.Body != nil {
			bodyReader := io.TeeReader(r.Body, &requestBody)
			r.Body = io.NopCloser(bodyReader)
		}

		// Оборачиваем ResponseWriter
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK, body: &bytes.Buffer{}}

		// Логируем информацию о запросе
		start := time.Now()
		log.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Str("request_body", requestBody.String()).
			Msg("Request started")

		// Передаем управление следующему обработчику
		next.ServeHTTP(lrw, r)

		// Логируем информацию о ответе
		duration := time.Since(start).Milliseconds() // Преобразуем в миллисекунды
		responseBody := lrw.body.String()

		// Убираем лишние символы, такие как \n
		if len(responseBody) > 0 && responseBody[len(responseBody)-1] == '\n' {
			responseBody = responseBody[:len(responseBody)-1]
		}

		// Логируем информацию о запросе
		log.Info().
			Str("path", r.URL.Path).
			Int64("duration_ms", duration).
			Int("status", lrw.statusCode).
			Str("response_body", responseBody).
			Msg("Request completed")
	})
}
