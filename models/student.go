package models

import (
	"github.com/fatih/structs"
)

// Teacher represents the domain model for a teacher
type Student struct {
	Name       string `json:"name" structs:"name"`
	Age        string `json:"age" structs:"age"`
	Department string `json:"department" structs:"department"`
	Grade     int     `json:"grade" structs:"grade"`
}

// Map converts struct into map
func (s *Student) Map() map[string]interface{} {
	return structs.Map(s)
}

// Names returns the list of fields
func (s *Student) Names() []string {
	return structs.Names(s)
}