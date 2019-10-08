package models

import (
	"github.com/fatih/structs"
)

type Student struct {
	ID         string `json:"id" structs:"id"  bson:"_id" structs:"id"`
	Name       string `json:"name" structs:"name"  bson:"name" structs:"name"`
	Age        string `json:"age" structs:"age" bson:"age" structs:"age"`
	Department string `json:"department" structs:"department" bson:"department" structs:"department"`
	Grade      int    `json:"grade" structs:"grade" bson:"grade" structs:"grade"`
}

func (s *Student) Map() map[string]interface{} {
	return structs.Map(s)
}

func (s *Student) Names() []string {
	return structs.Names(s)
}
