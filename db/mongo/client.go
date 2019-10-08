package mongo

import (
	"context"
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ahab94/ApnaSchool/db"
	"github.com/ahab94/ApnaSchool/models"
)

const (
	dbName     = "ApnaSchool"
	collection = "Student"
)

func init() {
	db.Register("mongo", NewClient)
}

type client struct {
	conn *mongo.Client
}

func NewClient(conf db.Option) (db.DataStore, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		return nil, err
	}

	return &client{conn: cli}, nil
}

func (c *client) AddStudent(student *models.Student) (string, error) {
	if student.ID != "" {
		return "", errors.New("id is not empty")
	}

	student.ID = uuid.NewV4().String()
	collection := c.conn.Database(dbName).Collection(collection)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := collection.InsertOne(ctx, student)
	if err != nil {
		return "", err
	}

	return student.ID, nil
}

func (c *client) GetStudent(id string) (*models.Student, error) {
	var stu *models.Student
	collection := c.conn.Database(dbName).Collection(collection)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&stu)
	if err != nil {
		return nil, err
	}

	return stu, nil
}

func (c *client) DeleteStudent(id string) error {
	collection := c.conn.Database(dbName).Collection(collection)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}

func (c *client) UpdateStudent(student *models.Student) error {
	collection := c.conn.Database(dbName).Collection(collection)
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": student.ID}, bson.M{"$set": student})
	if err != nil {
		return err
	}

	return nil
}
