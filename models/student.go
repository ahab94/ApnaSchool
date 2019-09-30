package models

import (
	"github.com/fatih/structs"
)

// Teacher represents the domain model for a teacher
type Teacher struct {
	Name       string `json:"name" structs:"name"`
	Age        string `json:"age" structs:"age"`
	Department string `json:"department" structs:"department"`
	Salary     int    `json:"salary" structs:"salary"`
}

// Map converts struct into map
func (t *Teacher) Map() map[string]interface{} {
	return structs.Map(t)
}

// Names returns the list of fields
func (t *Teacher) Names() []string {
	return structs.Names(t)
}
