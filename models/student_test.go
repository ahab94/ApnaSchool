package models

import (
	"reflect"
	"testing"
)

func TestStudent_Map(t *testing.T) {
	type fields struct {
		ID         string
		Name       string
		Age        string
		Department string
		Grade      int
		Password   string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: " success - student struct to map",
			fields: fields{
				ID:         "12345",
				Name:       "ahmed",
				Age:        "22",
				Department: "IT",
				Grade:      10,
				Password:   "mypass",
			},
			want: map[string]interface{}{
				"id":         "12345",
				"name":       "ahmed",
				"age":        "22",
				"department": "IT",
				"grade":      10,
				"password":   "mypass",
			},
		},
		{
			name: " success - student struct to map with fewer fields",
			fields: fields{
				Name:       "ahmed",
				Age:        "22",
				Department: "IT",
			},
			want: map[string]interface{}{
				"id":         "",
				"name":       "ahmed",
				"age":        "22",
				"department": "IT",
				"grade":      0,
				"password":   "mypass",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Student{
				ID:         tt.fields.ID,
				Name:       tt.fields.Name,
				Age:        tt.fields.Age,
				Department: tt.fields.Department,
				Grade:      tt.fields.Grade,
				Password:   tt.fields.Password,
			}
			if got := s.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudent_Names(t *testing.T) {
	type fields struct {
		ID         string
		Name       string
		Age        string
		Department string
		Grade      int
		Password   string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: " success - names of student struct fields",
			fields: fields{
				ID:         "12345",
				Name:       "ahmed",
				Age:        "22",
				Department: "IT",
				Grade:      10,
				Password:   "mypass",
			},
			want: []string{"id", "name", "age", "department", "grade", "password"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Student{
				ID:         tt.fields.ID,
				Name:       tt.fields.Name,
				Age:        tt.fields.Age,
				Department: tt.fields.Department,
				Grade:      tt.fields.Grade,
				Password:   tt.fields.Password,
			}
			if got := s.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
