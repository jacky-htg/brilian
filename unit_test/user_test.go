package unit_test

import (
	"errors"
	"testing"
)

type MockUserRepository struct {
	Users map[int]*User
}

func (m *MockUserRepository) GetByID(id int) (*User, error) {
	if user, ok := m.Users[id]; ok {
		return user, nil
	}
	return nil, errors.New("user not found")
}

func TestGetUser(t *testing.T) {
	mockRepo := &MockUserRepository{
		Users: map[int]*User{
			1: {ID: 1, Name: "John Doe"},
		},
	}

	user, err := GetUser(mockRepo, 1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.Name != "John Doe" {
		t.Errorf("expected 'John Doe', got %v", user.Name)
	}

	_, err = GetUser(mockRepo, 2)
	if err == nil {
		t.Fatalf("expected error, got none")
	}
}
