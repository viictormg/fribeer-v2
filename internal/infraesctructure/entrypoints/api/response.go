package api

import "time"

type Response struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Error     []string    `json:"error,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}
