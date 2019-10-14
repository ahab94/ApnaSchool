package runtime

import (
	"github.com/ahab94/ApnaSchool/db"
	"github.com/ahab94/ApnaSchool/db/mongo"
)

// Runtime initializes values for entry point to our application
type Runtime struct {
	Db db.DataStore
}

// NewRuntime creates a connection to our database
func NewRuntime() (*Runtime, error) {
	store, err := mongo.NewClient(db.Option{})
	if err != nil {
		return nil, err
	}
	return &Runtime{store}, err
}
