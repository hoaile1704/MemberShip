package models

// ResponseFormat represents the structure of the API response
type ResponseFormat struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
