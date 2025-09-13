package ver1

import (
	"encoding/json"
	"net/http"

	"github.com/amagkn/my-go-clean-architecture-template/pkg/logger"
)

func successResponse(w http.ResponseWriter, statusCode int, output any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if output == nil {
		return
	}

	err := writeJSON(w, output)
	if err != nil {
		logger.Error(err, "writeJSON")
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}
}

func errorResponse(w http.ResponseWriter, statusCode int, payload errorPayload) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := writeJSON(w, errorJSONBody{
		errorJSONPayload{
			Type:    payload.Type.Error(),
			Details: payload.Details,
		},
	})

	if err != nil {
		logger.Error(err, "writeJSON")
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}
}

type errorPayload struct {
	Type    error
	Details any
}

type errorJSONPayload struct {
	Type    string `json:"type"`
	Details any    `json:"details"`
}

type errorJSONBody struct {
	Error errorJSONPayload `json:"error"`
}

func writeJSON(w http.ResponseWriter, output any) error {
	b, err := json.Marshal(output)
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	if err != nil {
		return err
	}

	return nil
}
