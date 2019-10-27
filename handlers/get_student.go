package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/ahab94/ApnaSchool"
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
	_, err := d.rt.Service().RetrieveStudent(params.ID)
	if err != nil {
		return operations.NewGetStudentNotFound()
	}
	return operations.NewGetStudentOK()
}
