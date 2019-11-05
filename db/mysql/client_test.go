package mysql

import (
	"os"
	"reflect"
	"testing"

	"github.com/ahab94/ApnaSchool/db"
	"github.com/ahab94/ApnaSchool/models"
)

func Test_client_AddStudent(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "apna-school-mysql-db")
	os.Setenv("DB_USER", "root")

	type args struct {
		student *models.Student
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - add student in db",
			args:    args{student: &models.Student{Name: "Ahmed", Age: "12", Department: "Army", Grade: 430, Password: "mypass"}},
			wantErr: false,
		},
		{
			name:    "fail - add invalid student in db",
			args:    args{student: &models.Student{ID: "4", Name: "Talha", Age: "23", Department: "SE", Grade: 85, Password: "mypass"}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(db.Option{})
			_, err := c.AddStudent(tt.args.student)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_client_AddTeacher(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "apna-school-mysql-db")
	os.Setenv("DB_USER", "root")

	type args struct {
		teacher *models.Teacher
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - add teacher in db",
			args:    args{teacher: &models.Teacher{Name: "Ahmed", Age: "12", Department: "Army", Salary: 430, Password: "mypass"}},
			wantErr: false,
		},
		{
			name:    "fail - add invalid teacher in db",
			args:    args{teacher: &models.Teacher{ID: "2", Name: "Talha", Age: "23", Department: "SE", Salary: 85, Password: "mypass"}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(db.Option{})
			_, err := c.AddTeacher(tt.args.teacher)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTeacher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_client_DeleteStudent(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "apna-school-mysql-db")
	os.Setenv("DB_USER", "root")

	c, _ := NewClient(db.Option{})
	student := &models.Student{Name: "Qamar Zaman Bajwa", Age: "12", Department: "Army", Grade: 430, Password: "mypass"}
	_, _ = c.AddStudent(student)
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - delete student from db",
			args:    args{id: student.ID},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.DeleteStudent(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_DeleteTeacher(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "apna-school-mysql-db")
	os.Setenv("DB_USER", "root")

	c, _ := NewClient(db.Option{})
	teacher := &models.Teacher{Name: "Qamar Zaman Bajwa", Age: "12", Department: "Army", Salary: 430, Password: "mypass"}
	_, _ = c.AddTeacher(teacher)
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - delete student from db",
			args:    args{id: teacher.ID},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.DeleteTeacher(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTeacher() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_GetStudent(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "apna-school-mysql-db")
	os.Setenv("DB_USER", "root")

	c, _ := NewClient(db.Option{})
	student := &models.Student{Name: "Qamar Zaman Bajwa", Age: "12", Department: "Army", Grade: 430, Password: "mypass"}
	_, _ = c.AddStudent(student)
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Student
		wantErr bool
	}{
		{
			name:    "success - get student from db",
			args:    args{id: student.ID},
			want:    student,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetStudent(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStudent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_GetTeacher(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "apna-school-mysql-db")
	os.Setenv("DB_USER", "root")

	c, _ := NewClient(db.Option{})
	teacher := &models.Teacher{Name: "Qamar Zaman Bajwa", Age: "12", Department: "Army", Salary: 430, Password: "mypass"}
	_, _ = c.AddTeacher(teacher)
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Teacher
		wantErr bool
	}{
		{
			name:    "success - get teacher from db",
			args:    args{id: teacher.ID},
			want:    teacher,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetTeacher(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTeacher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTeacher() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_UpdateStudent(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "apna-school-mysql-db")
	os.Setenv("DB_USER", "root")

	c, _ := NewClient(db.Option{})
	type args struct {
		student *models.Student
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - update student in db",
			args:    args{student: &models.Student{Name: "Talha", Age: "22", Department: "Wanclouds", Grade: 30, Password: "mypass"}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.UpdateStudent(tt.args.student); (err != nil) != tt.wantErr {
				t.Errorf("UpdateStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_UpdateTeacher(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "apna-school-mysql-db")
	os.Setenv("DB_USER", "root")

	c, _ := NewClient(db.Option{})
	type args struct {
		teacher *models.Teacher
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - updata teacher in db",
			args:    args{teacher: &models.Teacher{Name: "Talha", Age: "22", Department: "Wanclouds", Salary: 30, Password: "mypass"}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.UpdateTeacher(tt.args.teacher); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTeacher() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
