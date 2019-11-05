package mongo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/ahab94/ApnaSchool/config"
	"github.com/ahab94/ApnaSchool/db"
	"github.com/ahab94/ApnaSchool/models"
)

const (
	stuCollection = "Student"
	tchCollection = "Teacher"
)

func init() {
	db.Register("mongo", NewClient)
}

type client struct {
	conn *mongo.Client
}

// NewClient initializes a mysql database connection
func NewClient(conf db.Option) (db.DataStore, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	log().Infof("initializing mongodb: %s", uri)
	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}

	return &client{conn: cli}, nil
}

func (c *client) SignUpStudent(student *models.Student) (string, error) {
	if student.ID != "" {
		return "", errors.New("id is not empty")
	}

	student.ID = uuid.NewV4().String()
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(stuCollection)
	if _, err := collection.InsertOne(context.TODO(), student); err != nil {
		return "", errors.Wrap(err, "failed to add student")
	}

	return student.ID, nil
}

func (c *client) AddStudent(student *models.Student) (string, error) {
	if student.ID != "" {
		return "", errors.New("id is not empty")
	}

	student.ID = uuid.NewV4().String()
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(stuCollection)
	if _, err := collection.InsertOne(context.TODO(), student); err != nil {
		return "", errors.Wrap(err, "failed to add student")
	}

	return student.ID, nil
}

func (c *client) GetStudent(id string) (*models.Student, error) {
	var stu *models.Student
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(stuCollection)
	if err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&stu); err != nil {
		return nil, errors.Wrap(err, "failed to get student")
	}

	return stu, nil
}

func (c *client) DeleteStudent(id string) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(stuCollection)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {
		return errors.Wrap(err, "failed to delete student")
	}

	return nil
}

func (c *client) UpdateStudent(student *models.Student) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(stuCollection)
	if _, err := collection.UpdateOne(context.TODO(), bson.M{"_id": student.ID}, bson.M{"$set": student}); err != nil {
		return errors.Wrap(err, "failed to update student")
	}

	return nil
}

func (c *client) SignUpTeacher(teacher *models.Teacher) (string, error) {
	if teacher.ID != "" {
		return "", errors.New("id is not empty")
	}

	teacher.ID = uuid.NewV4().String()
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(tchCollection)
	if _, err := collection.InsertOne(context.TODO(), teacher); err != nil {
		return "", errors.Wrap(err, "failed to add teacher")
	}

	return teacher.ID, nil
}

func (c *client) AddTeacher(teacher *models.Teacher) (string, error) {
	if teacher.ID != "" {
		return "", errors.New("id is not empty")
	}

	teacher.ID = uuid.NewV4().String()
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(tchCollection)
	if _, err := collection.InsertOne(context.TODO(), teacher); err != nil {
		return "", errors.Wrap(err, "failed to add teacher")
	}

	return teacher.ID, nil
}

func (c *client) GetTeacher(id string) (*models.Teacher, error) {
	var tch *models.Teacher
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(tchCollection)
	if err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&tch); err != nil {
		return nil, errors.Wrap(err, "failed to get teacher")
	}

	return tch, nil
}

func (c *client) DeleteTeacher(id string) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(tchCollection)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {
		return errors.Wrap(err, "failed to delete teacher")
	}

	return nil
}

func (c *client) UpdateTeacher(teacher *models.Teacher) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(tchCollection)
	if _, err := collection.UpdateOne(context.TODO(), bson.M{"_id": teacher.ID}, bson.M{"$set": teacher}); err != nil {
		return errors.Wrap(err, "failed to update teacher")
	}

	return nil
}
