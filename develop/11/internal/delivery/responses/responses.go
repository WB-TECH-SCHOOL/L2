package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	ResponseBadQuery = "Bad query provided"
	ResponseBadBody  = "Bad body provided"

	ResponseSuccess = "Success"
)

type SuccessResponse struct {
	Result interface{} `json:"result,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}

func writeResponse(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("error occurred when encoding response body: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func WriteSuccess(w http.ResponseWriter, status int, data interface{}) {
	respBody := SuccessResponse{Result: data}
	writeResponse(w, status, respBody)
}

func writeError(w http.ResponseWriter, status int, errMsg string) {
	respBody := ErrorResponse{Error: errMsg}
	writeResponse(w, status, respBody)
}

func BadRequest(w http.ResponseWriter, errMsg string) {
	writeError(w, http.StatusBadRequest, errMsg)
}

func ServiceUnavailable(w http.ResponseWriter, errMsg string) {
	writeError(w, http.StatusServiceUnavailable, errMsg)
}

func NotFound(w http.ResponseWriter) {
	writeError(w, http.StatusNotFound, "resource not found")
}
