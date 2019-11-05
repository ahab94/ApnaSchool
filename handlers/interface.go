package handlers

import (
	"github.com/go-openapi/loads"

	runtime "github.com/ahab94/ApnaSchool"
	"github.com/ahab94/ApnaSchool/gen/restapi/operations"
)

// Handler replaces swagger handler
type Handler *operations.ApnaSchoolAPI

// NewHandler overrides swagger api handlers
func NewHandler(rt *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewApnaSchoolAPI(spec)

	// teacher handlers
	handler.SignUpTeacherHandler = NewSignUpTeacher(rt)
	handler.AddTeacherHandler = NewAddTeacher(rt)
	handler.GetTeacherHandler = NewGetTeacher(rt)
	handler.EditTeacherHandler = NewEditTeacher(rt)
	handler.DeleteTeacherHandler = NewDeleteTeacher(rt)

	// student handlers
	handler.SignUpStudentHandler = NewSignUpStudent(rt)
	handler.AddStudentHandler = NewAddStudent(rt)
	handler.GetStudentHandler = NewGetStudent(rt)
	handler.EditStudentHandler = NewEditStudent(rt)
	handler.DeleteStudentHandler = NewDeleteStudent(rt)

	return handler
}
