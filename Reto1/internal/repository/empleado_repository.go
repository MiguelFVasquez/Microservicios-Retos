package repository

import (
	"Reto1/internal/model"
	"errors"
	"sync"
)

type EmpleadoRepository struct {
	empleados map[string]model.Empleado
	mutex     sync.RWMutex
}

func NewEmpleadoRepository() *EmpleadoRepository {
	return &EmpleadoRepository{
		empleados: make(map[string]model.Empleado),
	}
}

func (r *EmpleadoRepository) Guardar(empleado model.Empleado) model.Empleado {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.empleados[empleado.ID] = empleado
	return empleado
}

func (r *EmpleadoRepository) ObtenerPorID(id string) (model.Empleado, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	empleado, existe := r.empleados[id]

	if !existe {
		return model.Empleado{}, errors.New("empleado no encontrado")
	}

	return empleado, nil
}
