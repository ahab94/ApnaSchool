package db

import (
	"log"

	"github.com/ahab94/ApnaSchool/models"
)

type DataStore interface {
	AddStudent(student *models.Student) (string, error)
	GetStudent(id string) (*models.Student, error)
	DeleteStudent(id string) error
	UpdateStudent(student *models.Student) error
}

type Option struct {
	TestMode bool
}

type DataStoreFactory func(conf Option) (DataStore, error)

var datastoreFactories = make(map[string]DataStoreFactory)

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
