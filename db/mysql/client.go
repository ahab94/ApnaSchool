package mysql

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"

	"github.com/ahab94/ApnaSchool/config"
	"github.com/ahab94/ApnaSchool/db"
	"github.com/ahab94/ApnaSchool/models"
)

const (
	studentTableName = "student"
	teacherTableName = "teacher"
)

func init() {
	db.Register("mysql", NewClient)
}

//The first implementation.
type client struct {
	db *sqlx.DB
}

func formatDSN() string {
	cfg := mysql.NewConfig()
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	cfg.DBName = viper.GetString(config.DbName)
	cfg.ParseTime = true
	cfg.User = viper.GetString(config.DbUser)
	cfg.Passwd = viper.GetString(config.DbPass)
	return cfg.FormatDSN()
}

// NewClient initializes a mysql database connection
func NewClient(conf db.Option) (db.DataStore, error) {
	log().Info("initializing mysql connection: " + formatDSN())
	cli, err := sqlx.Connect("mysql", formatDSN())
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}
	return &client{db: cli}, nil
}

func (c *client) SignUpStudent(student *models.Student) (string, error) {
	if student.ID != "" {
		return "", errors.New("id is not empty")
	}
	student.ID = uuid.NewV4().String()

	names := student.Names()
	if _, err := c.db.NamedExec(fmt.Sprintf(`INSERT INTO %s (%s) VALUES(%s)`, studentTableName, strings.Join(names, ","), strings.Join(mkPlaceHolder(names, ":", func(name, prefix string) string {
		return prefix + name
	}), ",")), student); err != nil {
		return "", errors.Wrap(err, "failed to add student")
	}

	return "", nil
}

func (c *client) AddStudent(student *models.Student) (string, error) {
	if student.ID != "" {
		return "", errors.New("id is not empty")
	}
	student.ID = uuid.NewV4().String()

	names := student.Names()
	if _, err := c.db.NamedExec(fmt.Sprintf(`INSERT INTO %s (%s) VALUES(%s)`, studentTableName, strings.Join(names, ","), strings.Join(mkPlaceHolder(names, ":", func(name, prefix string) string {
		return prefix + name
	}), ",")), student); err != nil {
		return "", errors.Wrap(err, "failed to add student")
	}

	return "", nil
}

func (c *client) GetStudent(id string) (*models.Student, error) {
	var stu models.Student
	if err := c.db.Get(&stu, fmt.Sprintf(`SELECT * FROM %s WHERE id = '%s'`, studentTableName, id)); err != nil {
		return nil, errors.Wrap(err, "failed to get student")
	}

	return &stu, nil
}

func (c *client) UpdateStudent(student *models.Student) error {
	names := student.Names()
	if _, err := c.db.NamedExec(fmt.Sprintf(`UPDATE %s SET %s`, studentTableName, strings.Join(mkPlaceHolder(names, "=:", func(name, prefix string) string {
		return name + prefix + name
	}), ",")), student); err != nil {
		return errors.Wrap(err, "failed to update student")
	}

	return nil
}

func (c *client) DeleteStudent(id string) error {
	if _, err := c.db.Query(fmt.Sprintf(`DELETE FROM %s WHERE id= '%s'`, studentTableName, id)); err != nil {
		return errors.Wrap(err, "failed to delete student")
	}

	return nil
}

func (c *client) SignUpTeacher(teacher *models.Teacher) (string, error) {
	if teacher.ID != "" {
		return "", errors.New("id is not empty")
	}
	teacher.ID = uuid.NewV4().String()

	names := teacher.Names()
	if _, err := c.db.NamedExec(fmt.Sprintf(`INSERT INTO %s (%s) VALUES(%s)`, teacherTableName, strings.Join(names, ","), strings.Join(mkPlaceHolder(names, ":", func(name, prefix string) string {
		return prefix + name
	}), ",")), teacher); err != nil {
		return "", errors.Wrap(err, "failed to add teacher")
	}

	return "", nil
}

func (c *client) AddTeacher(teacher *models.Teacher) (string, error) {
	if teacher.ID != "" {
		return "", errors.New("id is not empty")
	}
	teacher.ID = uuid.NewV4().String()

	names := teacher.Names()
	if _, err := c.db.NamedExec(fmt.Sprintf(`INSERT INTO %s (%s) VALUES(%s)`, teacherTableName, strings.Join(names, ","), strings.Join(mkPlaceHolder(names, ":", func(name, prefix string) string {
		return prefix + name
	}), ",")), teacher); err != nil {
		return "", errors.Wrap(err, "failed to add teacher")
	}

	return "", nil
}

func (c *client) GetTeacher(id string) (*models.Teacher, error) {
	var tch models.Teacher
	if err := c.db.Get(&tch, fmt.Sprintf(`SELECT * FROM %s WHERE id = '%s'`, teacherTableName, id)); err != nil {
		return nil, errors.Wrap(err, "failed to get teacher")
	}

	return &tch, nil
}

func (c *client) UpdateTeacher(teacher *models.Teacher) error {
	names := teacher.Names()
	if _, err := c.db.NamedExec(fmt.Sprintf(`UPDATE %s SET %s`, teacherTableName, strings.Join(mkPlaceHolder(names, "=:", func(name, prefix string) string {
		return name + prefix + name
	}), ",")), teacher); err != nil {
		return errors.Wrap(err, "failed to update teacher")
	}

	return nil
}

func (c *client) DeleteTeacher(id string) error {
	if _, err := c.db.Query(fmt.Sprintf(`DELETE FROM %s WHERE id= '%s'`, teacherTableName, id)); err != nil {
		return errors.Wrap(err, "failed to delete teacher")
	}

	return nil
}

func mkPlaceHolder(names []string, prefix string, formatName func(name, prefix string) string) []string {
	ph := make([]string, len(names))
	for i, name := range names {
		ph[i] = formatName(name, prefix)
	}

	return ph
}
