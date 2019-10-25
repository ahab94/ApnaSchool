package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/ahab94/ApnaSchool"
	"github.com/ahab94/ApnaSchool/gen/restapi/operations"
)

// NewEditTeacher handles a request for editing a teacher
func NewEditTeacher(rt *runtime.Runtime) operations.EditTeacherHandler {
	return &editTeacher{rt: rt}
}

type editTeacher struct {
	rt *runtime.Runtime
}

// Handle the put teacher request
func (d *editTeacher) Handle(params operations.EditTeacherParams) middleware.Responder {
	teacher, err := d.rt.Service().RetrieveTeacher(params.ID)
	if err != nil {
		return operations.NewGetTeacherNotFound()
	}
	if err := d.rt.Service().UpdateTeacher(teacher); err != nil {
		return operations.NewEditTeacherInternalServerError()
	}

	return operations.NewEditTeacherOK()
}
