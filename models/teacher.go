package models

import (
	"github.com/fatih/structs"
)

type Teacher struct {
	ID         string `json:"id" structs:"id" bson:"_id"`
	Name       string `json:"name" structs:"name" bson:"name"`
	Age        string `json:"age" structs:"age" bson:"age"`
	Department string `json:"department" structs:"department" bson:"department"`
	Salary     int    `json:"salary" structs:"salary" bson:"salary"`
}

func (t *Teacher) Map() map[string]interface{} {
	return structs.Map(t)
}

func (t *Teacher) Names() []string {
	fields := structs.Fields(t)
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
