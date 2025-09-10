package response

import (
	"encoding/json"
	"net/http"

	"github.com/amagkn/my-go-clean-architecture-template/pkg/logger"
)

func Success(w http.ResponseWriter, statusCode int, output any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if output != nil {
		writeJSON(w, output)
	}
}

func Error(w http.ResponseWriter, statusCode int, payload ErrorPayload) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	writeJSON(w, errorJSONBody{
		errorJSONPayload{
			Type:    payload.Type.Error(),
			Details: payload.Details,
		},
	})
}

func writeJSON(w http.ResponseWriter, output any) {
	b, err := json.Marshal(output)
	if err != nil {
		logger.Error(err, "response.writeJSON json.Marshal")
		http.Error(w, "marshal json error", http.StatusInternalServerError)

		return
	}

	_, err = w.Write(b)
	if err != nil {
		logger.Error(err, "response.writeJSON w.Write")
		http.Error(w, "response write error", http.StatusInternalServerError)
	}
}
