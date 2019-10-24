// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AddStudentCreatedCode is the HTTP code returned for type AddStudentCreated
const AddStudentCreatedCode int = 201

/*AddStudentCreated student added

swagger:response addStudentCreated
*/
type AddStudentCreated struct {
}

// NewAddStudentCreated creates AddStudentCreated with default headers values
func NewAddStudentCreated() *AddStudentCreated {

	return &AddStudentCreated{}
}

// WriteResponse to the client
func (o *AddStudentCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}