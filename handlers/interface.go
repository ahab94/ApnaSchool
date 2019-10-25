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
	handler.AddTeacherHandler = NewAddTeacher(rt)
	handler.GetTeacherHandler = NewGetTeacher(rt)
	handler.EditTeacherHandler = NewEditTeacher(rt)
	handler.DeleteTeacherHandler = NewDeleteTeacher(rt)
	return handler
}
