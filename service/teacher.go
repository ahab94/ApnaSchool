package service

import (
	"github.com/ahab94/ApnaSchool/models"
)

// AddTeacher adds teacher into database
func (s *Service) AddTeacher(teacher *models.Teacher) (string, error) {
	return s.db.AddTeacher(teacher)
}

// RetrieveTeacher gets teacher from database
func (s *Service) RetrieveTeacher(id string) (*models.Teacher, error) {
	teacher, err := s.db.GetTeacher(id)
	if err != nil {
		return nil, err
	}

	return teacher, nil
}

// DeleteTeacher deletes teacher from database
func (s *Service) DeleteTeacher(id string) error {
	_, err := s.db.GetTeacher(id)
	if err != nil {
		return err
	}

	return s.db.DeleteTeacher(id)
}

// UpdateTeacher updates teacher record in database
func (s *Service) UpdateTeacher(teacher *models.Teacher) error {
	teacher, err := s.db.GetTeacher(teacher.ID)
	if err != nil {
		return err
	}

	return s.db.UpdateTeacher(teacher)
}
