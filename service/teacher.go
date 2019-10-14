package service

import (
	"github.com/ahab94/ApnaSchool/models"
)

// AddTeacher adds teacher into database
func (s *Service) AddTeacher(teacher *models.Teacher) (string, error) {
	id, err := s.rt.Db.AddTeacher(teacher)
	if err != nil {
		return "", err
	}
	return id, nil
}

// RetrieveTeacher gets teacher from database
func (s *Service) RetrieveTeacher(id string) (*models.Teacher, error) {
	stu, err := s.rt.Db.GetTeacher(id)
	if err != nil {
		return nil, err
	}
	return stu, nil
}

// DeleteTeacher deletes teacher from database
func (s *Service) DeleteTeacher(id string) error {
	err := s.rt.Db.DeleteTeacher(id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateTeacher updates teacher record in database
func (s *Service) UpdateTeacher(teacher *models.Teacher) error {
	err := s.rt.Db.UpdateTeacher(teacher)
	if err != nil {
		return err
	}
	return nil
}
