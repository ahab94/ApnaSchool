package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/ahab94/ApnaSchool"
	"github.com/ahab94/ApnaSchool/gen/restapi/operations"
	"github.com/ahab94/ApnaSchool/models"
)

// NewAddTeacher handles request for saving teacher
func NewAddTeacher(rt *runtime.Runtime) operations.AddTeacherHandler {
	return &addTeacher{rt: rt}
}

type addTeacher struct {
	rt *runtime.Runtime
}

// Handle the add teacher request
func (d *addTeacher) Handle(params operations.AddTeacherParams) middleware.Responder {
	id, err := d.rt.Service().AddTeacher(&models.Teacher{
		Name:       params.Teacher.Name,
		Age:        params.Teacher.Age,
		Department: params.Teacher.Department,
		Salary:     int(params.Teacher.Salary),
		Password:   params.Teacher.Password,
	})
	if err != nil {
		log().Errorf("failed to add teacher: %s", err)
		return operations.NewAddTeacherConflict()
	}

	params.Teacher.ID = id
	return operations.NewAddTeacherCreated().WithPayload(params.Teacher)
}
