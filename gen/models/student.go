// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Student student
// swagger:model Student
type Student struct {

	// age
	Age string `json:"Age,omitempty"`

	// department
	Department string `json:"Department,omitempty"`

	// grade
	Grade int64 `json:"Grade,omitempty"`

	// ID
	ID string `json:"ID,omitempty"`

	// name
	Name string `json:"Name,omitempty"`
}

// Validate validates this student
func (m *Student) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Student) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Student) UnmarshalBinary(b []byte) error {
	var res Student
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
