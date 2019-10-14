package service

import (
	"reflect"
	"testing"

	runtime "github.com/ahab94/ApnaSchool"
	"github.com/ahab94/ApnaSchool/models"
)

func TestService_AddStudent(t *testing.T) {
	type fields struct {
		rt *runtime.Runtime
	}
	type args struct {
		student *models.Student
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				rt: tt.fields.rt,
			}
			got, err := s.AddStudent(tt.args.student)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddStudent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_DeleteStudent(t *testing.T) {
	type fields struct {
		rt *runtime.Runtime
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				rt: tt.fields.rt,
			}
			if err := s.DeleteStudent(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_RetrieveStudent(t *testing.T) {
	type fields struct {
		rt *runtime.Runtime
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Student
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				rt: tt.fields.rt,
			}
			got, err := s.RetrieveStudent(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("RetrieveStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RetrieveStudent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_UpdateStudent(t *testing.T) {
	type fields struct {
		rt *runtime.Runtime
	}
	type args struct {
		student *models.Student
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				rt: tt.fields.rt,
			}
			if err := s.UpdateStudent(tt.args.student); (err != nil) != tt.wantErr {
				t.Errorf("UpdateStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
