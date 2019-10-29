package service

import (
	"github.com/ahab94/ApnaSchool/models"
)

// AddStudent adds student into database
func (s *Service) AddStudent(student *models.Student) (string, error) {
	return s.db.AddStudent(student)
}

// RetrieveStudent gets student from database
func (s *Service) RetrieveStudent(id string) (*models.Student, error) {
	student, err := s.db.GetStudent(id)
	if err != nil {
		return nil, err
	}

	return student, nil
}

// DeleteStudent deletes student from database
func (s *Service) DeleteStudent(id string) error {
	_, err := s.db.GetStudent(id)
	if err != nil {
		return err
	}

	return s.db.DeleteStudent(id)
}

// UpdateStudent updates student record in database
func (s *Service) UpdateStudent(student *models.Student) error {
	student, err := s.db.GetStudent(student.ID)
	if err != nil {
		return err
	}

	return s.db.UpdateStudent(student)
}
