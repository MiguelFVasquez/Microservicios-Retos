// Package repository provides data access layer for employees.
package repository

import (
	"Reto1/internal/model"
	"errors"
	"sync"
)

// EmpleadoRepository manages employee storage in memory.
type EmpleadoRepository struct {
	empleados map[string]model.Empleado
	mutex     sync.RWMutex
}

// NewEmpleadoRepository creates a new instance of EmpleadoRepository.
func NewEmpleadoRepository() *EmpleadoRepository {
	return &EmpleadoRepository{
		empleados: make(map[string]model.Empleado),
	}
}

// Guardar stores an employee in the repository.
func (r *EmpleadoRepository) Guardar(empleado model.Empleado) model.Empleado {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.empleados[empleado.ID] = empleado
	return empleado
}

// ObtenerPorID retrieves an employee by ID from the repository.
func (r *EmpleadoRepository) ObtenerPorID(id string) (model.Empleado, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	empleado, existe := r.empleados[id]

	if !existe {
		return model.Empleado{}, errors.New("empleado no encontrado")
	}

	return empleado, nil
}
