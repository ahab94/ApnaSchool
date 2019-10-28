// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/ahab94/ApnaSchool/gen/models"
)

// GetTeacherOKCode is the HTTP code returned for type GetTeacherOK
const GetTeacherOKCode int = 200

/*GetTeacherOK teacher response

swagger:response getTeacherOK
*/
type GetTeacherOK struct {

	/*
	  In: Body
	*/
	Payload *models.Teacher `json:"body,omitempty"`
}

// NewGetTeacherOK creates GetTeacherOK with default headers values
func NewGetTeacherOK() *GetTeacherOK {

	return &GetTeacherOK{}
}

// WithPayload adds the payload to the get teacher o k response
func (o *GetTeacherOK) WithPayload(payload *models.Teacher) *GetTeacherOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get teacher o k response
func (o *GetTeacherOK) SetPayload(payload *models.Teacher) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTeacherOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTeacherNotFoundCode is the HTTP code returned for type GetTeacherNotFound
const GetTeacherNotFoundCode int = 404

/*GetTeacherNotFound teacher not found

swagger:response getTeacherNotFound
*/
type GetTeacherNotFound struct {
}

// NewGetTeacherNotFound creates GetTeacherNotFound with default headers values
func NewGetTeacherNotFound() *GetTeacherNotFound {

	return &GetTeacherNotFound{}
}

// WriteResponse to the client
func (o *GetTeacherNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetTeacherInternalServerErrorCode is the HTTP code returned for type GetTeacherInternalServerError
const GetTeacherInternalServerErrorCode int = 500

/*GetTeacherInternalServerError internal server error

swagger:response getTeacherInternalServerError
*/
type GetTeacherInternalServerError struct {
}

// NewGetTeacherInternalServerError creates GetTeacherInternalServerError with default headers values
func NewGetTeacherInternalServerError() *GetTeacherInternalServerError {

	return &GetTeacherInternalServerError{}
}

// WriteResponse to the client
func (o *GetTeacherInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
