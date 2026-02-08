// Package handlers provides HTTP handlers for employee operations.
package handlers

import (
	"Reto1/internal/model"
	"Reto1/internal/service"
	"encoding/json"
	"net/http"
	"strings"
)

// EmpleadoHandler handles HTTP requests for employees.
type EmpleadoHandler struct {
	service *service.EmpleadoService
}

// NewEmpleadoHandler creates a new instance of EmpleadoHandler.
func NewEmpleadoHandler(service *service.EmpleadoService) *EmpleadoHandler {
	return &EmpleadoHandler{
		service: service,
	}
}

// RegistrarEmpleado handles POST requests to register a new employee.
func (h *EmpleadoHandler) RegistrarEmpleado(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Recurso no encontrado", http.StatusNotFound)
		return
	}

	var empleado model.Empleado

	err := json.NewDecoder(r.Body).Decode(&empleado)

	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	resultado := h.service.RegistrarEmpleado(empleado)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultado)
}

// ObtenerEmpleado handles GET requests to retrieve an employee by ID.
func (h *EmpleadoHandler) ObtenerEmpleado(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Recurso no encontrado", http.StatusNotFound)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/empleados/")

	empleado, err := h.service.ObtenerEmpleado(id)

	if err != nil {
		http.Error(w, "El empleado con id "+id+" no existe", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(empleado)
}
