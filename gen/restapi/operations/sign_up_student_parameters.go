// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "github.com/ahab94/ApnaSchool/gen/models"
)

// NewSignUpStudentParams creates a new SignUpStudentParams object
// no default values defined in spec.
func NewSignUpStudentParams() SignUpStudentParams {

	return SignUpStudentParams{}
}

// SignUpStudentParams contains all the bound params for the sign up student operation
// typically these are obtained from a http.Request
//
// swagger:parameters SignUpStudent
type SignUpStudentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*student's details
	  In: body
	*/
	Student *models.Student
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSignUpStudentParams() beforehand.
func (o *SignUpStudentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Student
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("student", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Student = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}