package repositories

import (
	"DevDash/internal/models"
	"DevDash/pkg/utils"
	"context"
	"errors"
	"fmt"
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

func (r *UserRepositoryMock) GetByID(ctx context.Context, id int64) (*models.User, error) {
	for _, user := range r.DB.Users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("no user found")
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
	if user.UUID == "" {
		user.UUID = fmt.Sprintf("user-%d", len(r.DB.Users)+1)
	}
	if user.ID == 0 {
		user.ID = int64(len(r.DB.Users) + 1)
	}
	now := utils.NowUTC()
	user.CreatedAt = now
	user.UpdatedAt = now
	r.DB.Users[user.UUID] = *user
	return nil
}

func (r *UserRepositoryMock) Update(ctx context.Context, user *models.User) error {
	user.UpdatedAt = utils.NowUTC()
	user.NormalizeTimestamps()
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

func (r *ProjectRepositoryMock) GetByID(ctx context.Context, id int64) (*models.Project, error) {
	for _, project := range r.DB.Projects {
		if project.ID == id {
			return &project, nil
		}
	}
	return nil, errors.New("no project found")
}

func (r *ProjectRepositoryMock) GetByUUID(ctx context.Context, id string) (*models.Project, error) {
	project, ok := r.DB.Projects[id]
	if !ok {
		return nil, errors.New("no project found")
	}
	return &project, nil
}

func (r *ProjectRepositoryMock) GetAllByUserID(ctx context.Context, userID int64) ([]models.Project, error) {
	var projects []models.Project
	for _, p := range r.DB.Projects {
		if p.UserID == userID {
			projects = append(projects, p)
		}
	}
	return projects, nil
}

func (r *ProjectRepositoryMock) Create(ctx context.Context, project *models.Project) error {
	if project.UUID == "" {
		project.UUID = fmt.Sprintf("project-%d", len(r.DB.Projects)+1)
	}
	if project.ID == 0 {
		project.ID = int64(len(r.DB.Projects) + 1)
	}
	now := utils.NowUTC()
	project.CreatedAt = now
	project.UpdatedAt = now
	r.DB.Projects[project.UUID] = *project
	return nil
}

func (r *ProjectRepositoryMock) Update(ctx context.Context, project *models.Project) error {
	project.UpdatedAt = utils.NowUTC()
	project.NormalizeTimestamps()
	r.DB.Projects[project.UUID] = *project
	return nil
}

func (r *ProjectRepositoryMock) Delete(ctx context.Context, id string) error {
	delete(r.DB.Projects, id)
	return nil
}
