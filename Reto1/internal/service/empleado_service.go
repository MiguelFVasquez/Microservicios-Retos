// Package service provides business logic for employee operations.
package service

import (
	"Reto1/internal/model"
	"Reto1/internal/repository"
)

// EmpleadoService handles business logic for employees.
type EmpleadoService struct {
	repo *repository.EmpleadoRepository
}

// NewEmpleadoService creates a new instance of EmpleadoService.
func NewEmpleadoService(repo *repository.EmpleadoRepository) *EmpleadoService {
	return &EmpleadoService{
		repo: repo,
	}
}

// RegistrarEmpleado registers a new employee.
func (s *EmpleadoService) RegistrarEmpleado(empleado model.Empleado) model.Empleado {
	return s.repo.Guardar(empleado)
}

// ObtenerEmpleado retrieves an employee by ID.
func (s *EmpleadoService) ObtenerEmpleado(id string) (model.Empleado, error) {
	return s.repo.ObtenerPorID(id)
}
