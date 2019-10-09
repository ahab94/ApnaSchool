package models

import (
	"github.com/fatih/structs"
)

type Student struct {
	ID         string `json:"id" structs:"id"  bson:"_id" db:"id"`
	Name       string `json:"name" structs:"name"  bson:"name" db:"name"`
	Age        string `json:"age" structs:"age" bson:"age" db:"age"`
	Department string `json:"department" structs:"department" bson:"department" db:"department"`
	Grade      int    `json:"grade" structs:"grade" bson:"grade" db:"grade"`
}

func (s *Student) Map() map[string]interface{} {
	return structs.Map(s)
}

func (s *Student) Names() []string {
	fields := structs.Fields(s)
	names := make([]string, len(fields))

	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}

	return names
}
