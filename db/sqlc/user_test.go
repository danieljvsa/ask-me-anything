package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/danieljvsa/ask-me-anything/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: util.RandomNameString(), // randomly generated?
		Password: util.RandomNameString(),
		Email: util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)

	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}

func TestUpdateUserName(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUsernameParams{
		ID: user1.ID,
		Username: util.RandomNameString(),
	}

	user2, err := testQueries.UpdateUsername(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user1)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, arg.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)

	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}
func TestUpdateEmail(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateEmailParams{
		ID: user1.ID,
		Email: util.RandomEmail(),
	}

	user2, err := testQueries.UpdateEmail(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user1)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, arg.Email, user2.Email)

	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}
func TestUpdatePassword(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdatePasswordParams{
		ID: user1.ID,
		Password: util.RandomNameString(),
	}

	user2, err := testQueries.UpdatePassword(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user1)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, arg.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)

	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit: 5,
		Offset: 10,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}