package service

import (
	"Reto1/internal/model"
	"Reto1/internal/repository"
)

type EmpleadoService struct {
	repo *repository.EmpleadoRepository
}

func NewEmpleadoService(repo *repository.EmpleadoRepository) *EmpleadoService {
	return &EmpleadoService{
		repo: repo,
	}
}

func (s *EmpleadoService) RegistrarEmpleado(empleado model.Empleado) model.Empleado {
	return s.repo.Guardar(empleado)
}

func (s *EmpleadoService) ObtenerEmpleado(id string) (model.Empleado, error) {
	return s.repo.ObtenerPorID(id)
}
