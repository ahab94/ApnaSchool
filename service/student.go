package service

import (
	"github.com/ahab94/ApnaSchool/models"
)

// AddStudent adds student into database
func (s *Service) AddStudent(student *models.Student) (string, error) {
	id, err := s.rt.Db.AddStudent(student)
	if err != nil {
		return "", err
	}
	return id, nil
}

// RetrieveStudent gets student from database
func (s *Service) RetrieveStudent(id string) (*models.Student, error) {
	stu, err := s.rt.Db.GetStudent(id)
	if err != nil {
		return nil, err
	}
	return stu, nil
}

// DeleteStudent deletes student from database
func (s *Service) DeleteStudent(id string) error {
	err := s.rt.Db.DeleteStudent(id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateStudent updates student record in database
func (s *Service) UpdateStudent(student *models.Student) error {
	err := s.rt.Db.UpdateStudent(student)
	if err != nil {
		return err
	}
	return nil
}
