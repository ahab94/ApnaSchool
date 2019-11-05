package db

import (
	"log"

	"github.com/ahab94/ApnaSchool/models"
)

// DataStore is an interface for query ops
type DataStore interface {
	SignUpStudent(student *models.Student) (string, error)
	AddStudent(student *models.Student) (string, error)
	GetStudent(id string) (*models.Student, error)
	DeleteStudent(id string) error
	UpdateStudent(student *models.Student) error
	SignUpTeacher(teacher *models.Teacher) (string, error)
	AddTeacher(teacher *models.Teacher) (string, error)
	GetTeacher(id string) (*models.Teacher, error)
	DeleteTeacher(id string) error
	UpdateTeacher(teacher *models.Teacher) error
}

// Option holds configuration for data store clients
type Option struct {
	TestMode bool
}

// DataStoreFactory holds configuration for data store
type DataStoreFactory func(conf Option) (DataStore, error)

var datastoreFactories = make(map[string]DataStoreFactory)

// Register saves data store into a data store factory
func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Fatalf("Datastore factory %s does not exist.", name)
		return
	}
	_, ok := datastoreFactories[name]
	if ok {
		log.Fatalf("Datastore factory %s already registered. Ignoring.", name)
		return
	}
	datastoreFactories[name] = factory
}
