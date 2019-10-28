package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/ahab94/ApnaSchool"
	domainErr "github.com/ahab94/ApnaSchool/errors"
	"github.com/ahab94/ApnaSchool/gen/models"
	"github.com/ahab94/ApnaSchool/gen/restapi/operations"
)

// NewGetStudent handles a request for retrieving student
func NewGetStudent(rt *runtime.Runtime) operations.GetStudentHandler {
	return &getStudent{rt: rt}
}

type getStudent struct {
	rt *runtime.Runtime
}

// Handle the get student request
func (d *getStudent) Handle(params operations.GetStudentParams) middleware.Responder {
	student, err := d.rt.Service().RetrieveStudent(params.ID)
	if err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewGetStudentNotFound()
		default:
			return operations.NewGetStudentInternalServerError()
		}
	}

	return operations.NewGetStudentOK().WithPayload(&models.Student{
		Age:        student.Age,
		Department: student.Department,
		ID:         student.ID,
		Name:       student.Name,
		Grade:      int64(student.Grade),
	})
}
