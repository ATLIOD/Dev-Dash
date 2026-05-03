package repositories

import (
	"DevDash/internal/models"
	"context"
	"errors"
)

func NewMockRepo(db *models.MockDB) *Repository {
	return &Repository{
		User:    &UserRepositoryMock{DB: db},
		Project: &ProjectRepositoryMock{DB: db},
	}
}

type UserRepositoryMock struct {
	DB *models.MockDB
}

func (r *UserRepositoryMock) GetByUUID(ctx context.Context, id string) (*models.User, error) {
	user := r.DB.Users[id]
	if user.UUID == "" {
		return &user, errors.New("no user found")
	}
	return &user, nil
}

func (r *UserRepositoryMock) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	for _, val := range r.DB.Users {
		if val.Email == email {
			user = val
			break
		}
	}
	return &user, nil
}

func (r *UserRepositoryMock) Create(ctx context.Context, user *models.User) error {
	r.DB.Users[user.UUID] = *user
	return nil
}

func (r *UserRepositoryMock) Update(ctx context.Context, user *models.User) error {
	r.DB.Users[user.UUID] = *user
	return nil
}

func (r *UserRepositoryMock) Delete(ctx context.Context, id string) error {
	delete(r.DB.Users, id)
	return nil
}

type ProjectRepositoryMock struct {
	DB *models.MockDB
}

func (r *ProjectRepositoryMock) GetByUUID(ctx context.Context, id string) (*models.Project, error) {
	project, ok := r.DB.Projects[id]
	if !ok {
		return nil, errors.New("no project found")
	}
	return &project, nil
}

func (r *ProjectRepositoryMock) Create(ctx context.Context, project *models.Project) error {
	r.DB.Projects[project.UUID] = *project
	return nil
}

func (r *ProjectRepositoryMock) Update(ctx context.Context, project *models.Project) error {
	r.DB.Projects[project.UUID] = *project
	return nil
}

func (r *ProjectRepositoryMock) Delete(ctx context.Context, id string) error {
	delete(r.DB.Projects, id)
	return nil
}
