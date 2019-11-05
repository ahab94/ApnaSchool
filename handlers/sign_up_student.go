package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/ahab94/ApnaSchool"
	"github.com/ahab94/ApnaSchool/gen/restapi/operations"
	"github.com/ahab94/ApnaSchool/models"
)

// NewSignUpStudent handles request for saving student
func NewSignUpStudent(rt *runtime.Runtime) operations.SignUpStudentHandler {
	return &signUpStudent{rt: rt}
}

type signUpStudent struct {
	rt *runtime.Runtime
}

// Handle the sign up student request
func (d *signUpStudent) Handle(params operations.SignUpStudentParams) middleware.Responder {
	student := models.Student{
		Name:       params.Student.Name,
		Age:        params.Student.Age,
		Department: params.Student.Department,
		Grade:      int(params.Student.Grade),
		Password:   params.Student.Password,
	}

	id, err := d.rt.Service().SignUpStudent(&student)
	if err != nil {
		log().Errorf("failed to sign up student: %s", err)
		return operations.NewSignUpStudentConflict()
	}

	params.Student.ID = id
	return operations.NewSignUpStudentCreated().WithPayload(params.Student)
}
