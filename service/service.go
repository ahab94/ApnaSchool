package service

import runtime "github.com/ahab94/ApnaSchool"

// Service initiates database service at runtime
type Service struct {
	rt *runtime.Runtime
}

// NewService create a new service at runtime
func NewService() (*Service, error) {
	rt, err := runtime.NewRuntime()
	if err != nil {
		return nil, err
	}
	return &Service{rt}, nil
}
