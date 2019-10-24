package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/ahab94/ApnaSchool"
	"github.com/ahab94/ApnaSchool/gen/restapi/operations"
)

// NewDeleteTeacher handles a request for deleting a teacher
func NewDeleteTeacher(rt *runtime.Runtime) operations.DeleteTeacherHandler {
	return &deleteTeacher{rt: rt}
}

type deleteTeacher struct {
	rt *runtime.Runtime
}

// Handle the delete teacher request
func (d *deleteTeacher) Handle(params operations.DeleteTeacherParams) middleware.Responder {
	if err := d.rt.Service().DeleteTeacher(params.ID); err != nil {
		return operations.NewDeleteTeacherNotFound()
	}
	return operations.NewDeleteTeacherNoContent()
}
