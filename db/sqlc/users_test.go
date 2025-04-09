package database

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/trantho123/warehouse-management/utils"
)

func TestCreateUser(t *testing.T) {
	hashPass, err := utils.HashPassword("password123")
	require.NoError(t, err)

	creaUserTestParam := CreateUserParams{
		Username:  "testuser",
		Email:     "test@example.com",
		Password:  hashPass,
		RoleID:    sql.NullInt32{Int32: 1, Valid: true},
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	user, err := TestQuerier.CreateUser(context.Background(), creaUserTestParam)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, creaUserTestParam.Username, user.Username)
	require.Equal(t, creaUserTestParam.Email, user.Email)
	require.Equal(t, creaUserTestParam.Password, user.Password)
	require.Equal(t, creaUserTestParam.RoleID, user.RoleID)
	require.NotZero(t, user.CreatedAt)
}
