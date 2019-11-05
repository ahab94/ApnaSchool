package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/ahab94/ApnaSchool"
	"github.com/ahab94/ApnaSchool/gen/restapi/operations"
	"github.com/ahab94/ApnaSchool/models"
)

// NewSignUpTeacher handles request for saving teacher
func NewSignUpTeacher(rt *runtime.Runtime) operations.SignUpTeacherHandler {
	return &signUpTeacher{rt: rt}
}

type signUpTeacher struct {
	rt *runtime.Runtime
}

// Handle the sign up teacher request
func (d *signUpTeacher) Handle(params operations.SignUpTeacherParams) middleware.Responder {
	teacher := models.Teacher{
		Name:       params.Teacher.Name,
		Age:        params.Teacher.Age,
		Department: params.Teacher.Department,
		Salary:     int(params.Teacher.Salary),
		Password:   params.Teacher.Password,
	}

	id, err := d.rt.Service().SignUpTeacher(&teacher)
	if err != nil {
		log().Errorf("failed to sign up teacher: %s", err)
		return operations.NewSignUpTeacherConflict()
	}

	params.Teacher.ID = id
	return operations.NewSignUpTeacherCreated().WithPayload(params.Teacher)
}
