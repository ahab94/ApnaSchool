package models

import (
	"github.com/fatih/structs"
)

type Teacher struct {
	Name       string `json:"name" structs:"name"`
	Age        string `json:"age" structs:"age"`
	Department string `json:"department" structs:"department"`
	Salary     int    `json:"salary" structs:"salary"`
}

func (t *Teacher) Map() map[string]interface{} {
	return structs.Map(t)
}

func (t *Teacher) Names() []string {
	return structs.Names(t)
}
