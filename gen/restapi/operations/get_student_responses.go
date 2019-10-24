// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetStudentOKCode is the HTTP code returned for type GetStudentOK
const GetStudentOKCode int = 200

/*GetStudentOK student response

swagger:response getStudentOK
*/
type GetStudentOK struct {
}

// NewGetStudentOK creates GetStudentOK with default headers values
func NewGetStudentOK() *GetStudentOK {

	return &GetStudentOK{}
}

// WriteResponse to the client
func (o *GetStudentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetStudentNotFoundCode is the HTTP code returned for type GetStudentNotFound
const GetStudentNotFoundCode int = 404

/*GetStudentNotFound student not found

swagger:response getStudentNotFound
*/
type GetStudentNotFound struct {
}

// NewGetStudentNotFound creates GetStudentNotFound with default headers values
func NewGetStudentNotFound() *GetStudentNotFound {

	return &GetStudentNotFound{}
}

// WriteResponse to the client
func (o *GetStudentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}