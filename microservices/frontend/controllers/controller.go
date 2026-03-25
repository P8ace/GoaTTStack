package controllers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	responses "github.com/P8ace/GoaTTStack/microservices/common/types/Responses"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (self *Controller) HealthController(w http.ResponseWriter, r *http.Request) {
	response := responses.GenericResponse{
		Status: "Success",
		Body:   "API is healthy",
	}

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&response)
	if err != nil {
		slog.Error("Error while encoding to json", "Error:", err.Error())
		return
	}
}
