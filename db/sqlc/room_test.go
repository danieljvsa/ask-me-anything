package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/danieljvsa/ask-me-anything/util"
	"github.com/stretchr/testify/require"
)

func createRandomRoom(t *testing.T) Room {
	argUser := CreateUserParams{
		Username: util.RandomNameString(), // randomly generated?
		Password: util.RandomNameString(),
		Email: util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), argUser)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	room, err := testQueries.CreateRoom(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, room)
	
	require.Equal(t, user.ID, room.UserID)

	require.NotZero(t, room.ID)
	require.NotZero(t, room.CreatedAt)
	require.NotZero(t, room.UpdatedAt)

	return room
}

func TestCreateRoom(t *testing.T) {
	createRandomRoom(t)
}

func TestGetRoom(t *testing.T) {
	room1 := createRandomRoom(t)
	room2, err := testQueries.GetRoom(context.Background(), room1.ID)
	
	require.NoError(t, err)
	require.NotEmpty(t, room1)

	require.Equal(t, room1.ID, room2.ID)
	require.Equal(t, room1.UserID, room2.UserID)
	require.Equal(t, room1.CreatedAt, room2.CreatedAt)
	require.Equal(t, room1.UpdatedAt, room2.UpdatedAt)

	require.WithinDuration(t, room1.CreatedAt.Time, room2.CreatedAt.Time, time.Second)
}

func TestUpdateRoomUser(t *testing.T) {
	room1 := createRandomRoom(t)
	user1 := createRandomUser(t)

	arg := UpdateRoomParams{
		ID: room1.ID,
		UserID: user1.ID,
	}

	room2, err := testQueries.UpdateRoom(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user1)

	require.Equal(t, room1.ID, room2.ID)
	require.Equal(t, arg.UserID, room2.UserID)
	require.Equal(t, room1.CreatedAt, room2.CreatedAt)
	require.Equal(t, room1.UpdatedAt, room2.UpdatedAt)

	require.WithinDuration(t, user1.CreatedAt.Time, room2.CreatedAt.Time, time.Second)
}

func TestDeleteRoom(t *testing.T) {
	room1 := createRandomRoom(t)

	err := testQueries.DeleteRoom(context.Background(), room1.ID)
	require.NoError(t, err)

	room2, err := testQueries.GetRoom(context.Background(), room1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, room2)
}

func TestListRooms(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomRoom(t)
	}

	arg := ListRoomsParams{
		Limit: 5,
		Offset: 10,
	}

	rooms, err := testQueries.ListRooms(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, rooms, 5)

	for _, room := range rooms {
		require.NotEmpty(t, room)
	}
}