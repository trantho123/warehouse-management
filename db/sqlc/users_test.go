package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/trantho123/warehouse-management/utils"
)

func CreateUserForTest(t *testing.T) User {
	randPass := utils.RandPassword()
	hashPass, err := utils.HashPassword(randPass)
	require.NoError(t, err)

	creaUserTestParam := CreateUserParams{
		Username:  utils.RandString(10),
		Email:     utils.RandEmail(),
		Password:  hashPass,
		RoleID:    pgtype.Int4{Int32: 1, Valid: true},
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	user, err := TestStore.CreateUser(context.Background(), creaUserTestParam)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, creaUserTestParam.Username, user.Username)
	require.Equal(t, creaUserTestParam.Email, user.Email)
	require.Equal(t, creaUserTestParam.Password, user.Password)
	require.Equal(t, creaUserTestParam.RoleID, user.RoleID)
	require.NotZero(t, user.CreatedAt)
	return user
}
func TestCreateUser(t *testing.T) {
	CreateUserForTest(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateUserForTest(t)

	user2, err := TestStore.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.RoleID, user2.RoleID)
}

func TestGetUserByEmail(t *testing.T) {
	user1 := CreateUserForTest(t)

	user2, err := TestStore.GetUserByEmail(context.Background(), user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.RoleID, user2.RoleID)

}

func TestGetUserByUsername(t *testing.T) {
	user1 := CreateUserForTest(t)
	user2, err := TestStore.GetUserByUsername(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.RoleID, user2.RoleID)
}

func TestUpdateUserOnlyUserName(t *testing.T) {
	oldUser := CreateUserForTest(t)
	newUserName := utils.RandString(5)
	arg := UpdateUserParams{
		ID: oldUser.ID,
		Username: pgtype.Text{
			String: newUserName,
			Valid:  true,
		},
	}
	updatedUser, err := TestStore.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)

	require.Equal(t, oldUser.ID, updatedUser.ID)
	require.NotEqual(t, oldUser.Username, updatedUser.Username)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, oldUser.Password, updatedUser.Password)
	require.WithinDuration(t, time.Now(), updatedUser.UpdatedAt.Time, time.Second)

}
func TestUpdateUserOnlyEmail(t *testing.T) {
	oldUser := CreateUserForTest(t)
	newEmail := utils.RandEmail()
	arg := UpdateUserParams{
		ID: oldUser.ID,
		Email: pgtype.Text{
			String: newEmail,
			Valid:  true,
		},
	}
	updatedUser, err := TestStore.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)

	require.Equal(t, oldUser.ID, updatedUser.ID)
	require.NotEqual(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, oldUser.Username, updatedUser.Username)
	require.Equal(t, oldUser.Password, updatedUser.Password)
	require.WithinDuration(t, time.Now(), updatedUser.UpdatedAt.Time, time.Second)
}
func TestUpdateUserOnlyPassword(t *testing.T) {
	oldUser := CreateUserForTest(t)
	newPass := utils.RandPassword()
	arg := UpdateUserParams{
		ID: oldUser.ID,
		Password: pgtype.Text{
			String: newPass,
			Valid:  true,
		},
	}
	updatedUser, err := TestStore.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)

	require.Equal(t, oldUser.ID, updatedUser.ID)
	require.NotEqual(t, oldUser.Password, updatedUser.Password)
	require.Equal(t, oldUser.Username, updatedUser.Username)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.WithinDuration(t, time.Now(), updatedUser.UpdatedAt.Time, time.Second)
}

func TestUpdateUserAllFields(t *testing.T) {
	oldUser := CreateUserForTest(t)
	newUserName := utils.RandString(5)
	newPass := utils.RandPassword()
	newEmai := utils.RandEmail()
	arg := UpdateUserParams{
		ID: oldUser.ID,
		Username: pgtype.Text{
			String: newUserName,
			Valid:  true,
		},
		Email: pgtype.Text{
			String: newEmai,
			Valid:  true,
		},
		Password: pgtype.Text{
			String: newPass,
			Valid:  true,
		},
	}
	updatedUser, err := TestStore.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)

	require.Equal(t, oldUser.ID, updatedUser.ID)
	require.NotEqual(t, oldUser.Password, updatedUser.Password)
	require.NotEqual(t, oldUser.Username, updatedUser.Username)
	require.NotEqual(t, oldUser.Email, updatedUser.Email)
	require.WithinDuration(t, time.Now(), updatedUser.UpdatedAt.Time, time.Second)
}
