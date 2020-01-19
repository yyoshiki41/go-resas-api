package resas

import (
	"fmt"
	"net/http"
)

type Body struct {
	StatusCode  string `json:"statusCode"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func (b *Body) GetStatusCode() int {
	switch b.StatusCode {
	case "":
		return http.StatusOK
	case "400":
		return http.StatusBadRequest
	case "403":
		return http.StatusForbidden
	case "404":
		return http.StatusNotFound
	default:
	}
	return http.StatusInternalServerError
}

func (b *Body) Err() error {
	if b.GetStatusCode() != http.StatusOK {
		return fmt.Errorf("[StatusCode %s] message: %s %s",
			b.StatusCode, b.Message, b.Description)
	}
	return nil
}
