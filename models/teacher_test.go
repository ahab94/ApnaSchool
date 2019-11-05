package models

import (
	"reflect"
	"testing"
)

func TestTeacher_Map(t1 *testing.T) {
	type fields struct {
		ID         string
		Name       string
		Age        string
		Department string
		Salary     int
		Password   string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: " success - teacher struct to map",
			fields: fields{
				ID:         "231",
				Name:       "Talha",
				Age:        "23",
				Department: "IT",
				Salary:     20000,
				Password:   "mypass",
			},
			want: map[string]interface{}{
				"id":         "231",
				"name":       "Talha",
				"age":        "23",
				"department": "IT",
				"salary":     20000,
				"password":   "mypass",
			},
		},
		{
			name: " success - teacher struct to map with fewer fields",
			fields: fields{
				Name:       "Talha",
				Age:        "23",
				Department: "IT",
			},
			want: map[string]interface{}{
				"id":         "",
				"name":       "Talha",
				"age":        "23",
				"department": "IT",
				"salary":     0,
				"password":   "mypass",
			},
		},
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Teacher{
				ID:         tt.fields.ID,
				Name:       tt.fields.Name,
				Age:        tt.fields.Age,
				Department: tt.fields.Department,
				Salary:     tt.fields.Salary,
				Password:   tt.fields.Password,
			}
			if got := t.Map(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTeacher_Names(t1 *testing.T) {
	type fields struct {
		ID         string
		Name       string
		Age        string
		Department string
		Salary     int
		Password   string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: " success - return list of teacher fields",
			fields: fields{
				ID:         "231",
				Name:       "Talha",
				Age:        "23",
				Department: "CS",
				Salary:     20000,
				Password:   "mypass",
			},
			want: []string{"id", "name", "age", "department", "salary", "password"},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Teacher{
				ID:         tt.fields.ID,
				Name:       tt.fields.Name,
				Age:        tt.fields.Age,
				Department: tt.fields.Department,
				Salary:     tt.fields.Salary,
				Password:   tt.fields.Password,
			}
			if got := t.Names(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
