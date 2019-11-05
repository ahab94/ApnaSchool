package service

import (
	"github.com/ahab94/ApnaSchool/models"
)

// SignUpStudent signs up the student and make its account
func (s *Service) SignUpStudent(student *models.Student) (string, error) {
	return s.db.AddStudent(student)
}

// AddStudent adds student into database
func (s *Service) AddStudent(student *models.Student) (string, error) {
	return s.db.AddStudent(student)
}

// RetrieveStudent gets student from database
func (s *Service) RetrieveStudent(id string) (*models.Student, error) {
	return s.db.GetStudent(id)
}

// DeleteStudent deletes student from database
func (s *Service) DeleteStudent(id string) error {
	return s.db.DeleteStudent(id)
}

// UpdateStudent updates student record in database
func (s *Service) UpdateStudent(student *models.Student) error {
	return s.db.UpdateStudent(student)
}
