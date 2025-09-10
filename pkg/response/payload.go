package response

type ErrorPayload struct {
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
