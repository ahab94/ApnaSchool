package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/ahab94/ApnaSchool"
	"github.com/ahab94/ApnaSchool/gen/restapi/operations"
)

// NewEditStudent function is for edit student
func NewEditStudent(rt *runtime.Runtime) operations.EditStudentHandler {
	return &editStudent{rt: rt}
}

type editStudent struct {
	rt *runtime.Runtime
}

// Handle the put student request
func (d *editStudent) Handle(params operations.EditStudentParams) middleware.Responder {
	student, err := d.rt.Service().RetrieveStudent(params.ID)
	if err != nil {
		return operations.NewGetStudentNotFound()
	}
	if err := d.rt.Service().UpdateStudent(student); err != nil {
		return operations.NewEditStudentInternalServerError()
	}

	return operations.NewEditStudentOK()
}
