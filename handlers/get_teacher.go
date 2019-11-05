package handlers

import (
	runtime "github.com/ahab94/ApnaSchool"
	"github.com/ahab94/ApnaSchool/gen/models"
	"github.com/ahab94/ApnaSchool/gen/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetTeacher handles a request for retrieving teacher
func NewGetTeacher(rt *runtime.Runtime) operations.GetTeacherHandler {
	return &getTeacher{rt: rt}
}

type getTeacher struct {
	rt *runtime.Runtime
}

// Handle the get teacher request
func (d *getTeacher) Handle(params operations.GetTeacherParams) middleware.Responder {
	teacher, err := d.rt.Service().RetrieveTeacher(params.ID)
	if err != nil {
		return operations.NewGetTeacherNotFound()
	}
	return operations.NewGetTeacherOK().WithPayload(&models.Teacher{
		Age:        teacher.Age,
		Department: teacher.Department,
		ID:         teacher.ID,
		Name:       teacher.Name,
		Salary:     int64(teacher.Salary),
		Password:   teacher.Password,
	})
}
