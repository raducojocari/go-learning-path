package controllers

import (
	"fmt"
	"net/http"
)

type healthController struct {
}

func newHealthController() *healthController {
	return &healthController{}
}

func (h healthController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ping")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pong"))
}
