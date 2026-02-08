// Package model contains the data structures used in the application.
package model

// Empleado represents an employee with ID, name, and position.
type Empleado struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
	Cargo  string `json:"cargo"`
}
