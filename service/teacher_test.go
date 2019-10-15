package service

import (
	"reflect"
	"testing"

	runtime "github.com/ahab94/ApnaSchool"
	"github.com/ahab94/ApnaSchool/models"
)

func TestService_AddTeacher(t *testing.T) {
	type fields struct {
		rt *runtime.Runtime
	}
	type args struct {
		teacher *models.Teacher
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
			got, err := s.AddTeacher(tt.args.teacher)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTeacher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddTeacher() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_DeleteTeacher(t *testing.T) {
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
			if err := s.DeleteTeacher(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTeacher() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_RetrieveTeacher(t *testing.T) {
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
		want    *models.Teacher
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				rt: tt.fields.rt,
			}
			got, err := s.RetrieveTeacher(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("RetrieveTeacher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RetrieveTeacher() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_UpdateTeacher(t *testing.T) {
	type fields struct {
		rt *runtime.Runtime
	}
	type args struct {
		teacher *models.Teacher
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
			if err := s.UpdateTeacher(tt.args.teacher); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTeacher() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
