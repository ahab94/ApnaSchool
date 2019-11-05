package service

import (
	"github.com/ahab94/ApnaSchool/models"
)

// SignUpTeacher signs up the teacher and make its account
func (s *Service) SignUpTeacher(teacher *models.Teacher) (string, error) {
	return s.db.AddTeacher(teacher)
}

// AddTeacher adds teacher into database
func (s *Service) AddTeacher(teacher *models.Teacher) (string, error) {
	return s.db.AddTeacher(teacher)
}

// RetrieveTeacher gets teacher from database
func (s *Service) RetrieveTeacher(id string) (*models.Teacher, error) {
	return s.db.GetTeacher(id)
}

// DeleteTeacher deletes teacher from database
func (s *Service) DeleteTeacher(id string) error {
	return s.db.DeleteTeacher(id)
}

// UpdateTeacher updates teacher record in database
func (s *Service) UpdateTeacher(teacher *models.Teacher) error {
	return s.db.UpdateTeacher(teacher)
}
