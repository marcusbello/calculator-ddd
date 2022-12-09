package utils

// Response example
type Response struct {
	Data  string `json:"data"`
	Error error  `json:"error,omitempty"`
} //@name Response
