package models

import (
	"github.com/fatih/structs"
)

// Teacher holds information for a teacher
type Teacher struct {
	ID         string `json:"id" structs:"id" bson:"_id" db:"id"`
	Name       string `json:"name" structs:"name" bson:"name" db:"name"`
	Age        string `json:"age" structs:"age" bson:"age" db:"age"`
	Department string `json:"department" structs:"department" bson:"department" db:"department"`
	Salary     int    `json:"salary" structs:"salary" bson:"salary" db:"salary"`
	Password   string `json:"password" structs:"password" bson:"password" db:"password"`
}

// Map converts structs to a map representation
func (t *Teacher) Map() map[string]interface{} {
	return structs.Map(t)
}

// Names returns the field names of Teacher model
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
