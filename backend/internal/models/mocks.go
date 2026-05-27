package models

import "time"

type MockDB struct {
	Users    map[string]User
	Projects map[string]Project
}

func NewMockDB() *MockDB {
	db := &MockDB{
		Users:    make(map[string]User),
		Projects: make(map[string]Project),
	}
	db.Users["01"] = User{
		ID: 1, UUID: "01", Name: "User 1",
		Email: "user1@example.com", PasswordHash: "123123",
		CreatedAt: time.Now(), UpdatedAt: time.Now(),
	}
	db.Users["02"] = User{
		ID: 2, UUID: "02", Name: "User 2",
		Email: "user2@example.com", PasswordHash: "123123",
		CreatedAt: time.Now(), UpdatedAt: time.Now(),
	}

	db.Projects["01"] = Project{
		ID: 1, UUID: "01", Name: "Project 1",
		Description: "this is a description for project 1",
		Status:      "in progress", Stack: "golang, react",
		RepositoryURL: "example.com", DeploymentURL: "example.com",
		UserID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(),
	}
	return db
}
