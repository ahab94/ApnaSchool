package models

import (
	"github.com/fatih/structs"
)

type Student struct {
	Name       string `json:"name" structs:"name"`
	Age        string `json:"age" structs:"age"`
	Department string `json:"department" structs:"department"`
	Grade      int    `json:"grade" structs:"grade"`
}

func (s *Student) Map() map[string]interface{} {
	return structs.Map(s)
}

func (s *Student) Names() []string {
	return structs.Names(s)
}
