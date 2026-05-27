//go:build integration

package Unit

import (
	"DevDash/Test"
	"DevDash/Test/Unit/utils"
	"DevDash/internal/models"
	"DevDash/internal/services"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {
	repo := utils.Setup()
	svc := services.New(repo)

	t.Run("Create User", func(t *testing.T) {
		c := Test.NewChecker(t)
		req := models.CreateUserRequest{
			Name:     "New User",
			Email:    "new@example.com",
			Password: "password123",
		}

		resp, err := svc.User.Create(context.Background(), req)
		c.Check(assert.NoError(t, err))
		c.Check(assert.NotNil(t, resp))
		c.Check(assert.Equal(t, req.Name, resp.Name))
		c.Check(assert.Equal(t, req.Email, resp.Email))
	})

	t.Run("Get User", func(t *testing.T) {
		c := Test.NewChecker(t)
		resp, err := svc.User.GetByUUID(context.Background(), "01")
		c.Check(assert.NoError(t, err))
		c.Check(assert.NotNil(t, resp))
		c.Check(assert.Equal(t, "User 1", resp.Name))
	})

	t.Run("Update User", func(t *testing.T) {
		c := Test.NewChecker(t)
		req := models.UpdateUserRequest{
			Name:  "Updated User",
			Email: "updated@example.com",
		}
		resp, err := svc.User.Update(context.Background(), "02", req)
		c.Check(assert.NoError(t, err))
		c.Check(assert.NotNil(t, resp))
		c.Check(assert.Equal(t, req.Name, resp.Name))
	})

	t.Run("Delete User", func(t *testing.T) {
		c := Test.NewChecker(t)
		err := svc.User.Delete(context.Background(), "01")
		c.Check(assert.NoError(t, err))

		_, err = svc.User.GetByUUID(context.Background(), "01")
		c.Check(assert.Error(t, err))
		c.Check(assert.Contains(t, err.Error(), "no user found"))
	})
}
