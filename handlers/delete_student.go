package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/ahab94/ApnaSchool"
	"github.com/ahab94/ApnaSchool/gen/restapi/operations"
)

// NewDeleteStudent function will delete the student
func NewDeleteStudent(rt *runtime.Runtime) operations.DeleteStudentHandler {
	return &deleteStudent{rt: rt}
}

type deleteStudent struct {
	rt *runtime.Runtime
}

// Handle the delete entry request
func (d *deleteStudent) Handle(params operations.DeleteStudentParams) middleware.Responder {
	if err := d.rt.Service().DeleteStudent(params.ID); err != nil {
		return operations.NewDeleteStudentNotFound()
	}
	return operations.NewDeleteStudentNoContent()
}
