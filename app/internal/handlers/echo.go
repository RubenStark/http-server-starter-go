package handlers

import (
	"fmt"
	"net/http"
)

type EchoHandler struct {
	Message string
}

// Constructor que recibe el mensaje como prop
func NewEchoHandler(message string) *EchoHandler {
	return &EchoHandler{Message: message}
}

func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, h.Message)
}
