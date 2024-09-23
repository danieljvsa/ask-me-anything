package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/danieljvsa/ask-me-anything/util"
	"github.com/stretchr/testify/require"
)

func createRandomMessage(t *testing.T) Message {
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

	argMessage := CreateMessageParams{
		Message: sql.NullString{String: util.RandomMessageString(), Valid: true},
		UserID: user.ID,
		ParentID: sql.NullString{String: "0", Valid: true},
		RoomID: room.ID,
	}

	message, err := testQueries.CreateMessage(context.Background(), argMessage)
	require.NoError(t, err)
	require.NotEmpty(t, room)
	
	require.Equal(t, user.ID, message.UserID)
	require.Equal(t, room.ID, message.RoomID)
	require.Equal(t, argMessage.Message, message.Message)
	require.Equal(t, argMessage.ParentID, message.ParentID)

	require.NotZero(t, message.ID)
	require.NotZero(t, message.CreatedAt)
	require.NotZero(t, message.UpdatedAt)

	return message
}

func TestCreateMessage(t *testing.T) {
	createRandomMessage(t)
}

func TestGetMessage(t *testing.T) {
	message1 := createRandomMessage(t)
	message2, err := testQueries.GetMessage(context.Background(), message1.ID)
	
	require.NoError(t, err)
	require.NotEmpty(t, message1)

	require.Equal(t, message1.ID, message2.ID)
	require.Equal(t, message1.UserID, message2.UserID)
	require.Equal(t, message1.CreatedAt, message2.CreatedAt)
	require.Equal(t, message1.UpdatedAt, message2.UpdatedAt)

	require.WithinDuration(t, message1.CreatedAt.Time, message2.CreatedAt.Time, time.Second)
}

func TestUpdateMessageLikes(t *testing.T) {
	message1 := createRandomMessage(t)

	arg := UpdateLikesParams{
		ID: message1.ID,
		LikesCount: sql.NullInt64{Int64: util.RandomInt(1, 100), Valid: true},
	}

	message2, err := testQueries.UpdateLikes(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, message2)

	require.Equal(t, message1.ID, message2.ID)
	require.Equal(t, arg.LikesCount, message2.LikesCount)

	require.WithinDuration(t, message1.CreatedAt.Time, message2.CreatedAt.Time, time.Second)
}

func TestUpdateMessageAnswered(t *testing.T) {
	message1 := createRandomMessage(t)

	arg := UpdateAnsweredParams{
		ID: message1.ID,
		Answered: sql.NullBool{Bool: true, Valid: true},
	}

	message2, err := testQueries.UpdateAnswered(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, message2)

	require.Equal(t, message1.ID, message2.ID)
	require.Equal(t, arg.Answered, message2.Answered)

	require.WithinDuration(t, message1.CreatedAt.Time, message2.CreatedAt.Time, time.Second)
}

func TestUpdateMessageBody(t *testing.T) {
	message1 := createRandomMessage(t)

	arg := UpdateMessageParams{
		ID: message1.ID,
		Message: sql.NullString{String: util.RandomMessageString(), Valid: true},
	}

	message2, err := testQueries.UpdateMessage(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, message2)

	require.Equal(t, message1.ID, message2.ID)
	require.Equal(t, arg.Message, message2.Message)

	require.WithinDuration(t, message1.CreatedAt.Time, message2.CreatedAt.Time, time.Second)
}

func TestUpdateMessageParent(t *testing.T) {
	message1 := createRandomMessage(t)

	argUser := CreateUserParams{
		Username: util.RandomNameString(), // randomly generated?
		Password: util.RandomNameString(),
		Email: util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), argUser)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	arg := CreateMessageParams{
		UserID: user.ID,
		Message: sql.NullString{String: util.RandomMessageString(), Valid: true},
		ParentID: sql.NullString{String: "0", Valid: true},
		RoomID: message1.RoomID,
	}

	message2, err := testQueries.CreateMessage(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, message2)

	
	message1IDString := fmt.Sprintf("%d", message1.ID)
	argUpdate := UpdateParentParams{
		ID: message2.ID,
		ParentID: sql.NullString{String: message1IDString, Valid: true},
	}

	message3, err := testQueries.UpdateParent(context.Background(), argUpdate)
	require.NoError(t, err)
	require.NotEmpty(t, message3)

	require.Equal(t, message2.ID, message3.ID)
	require.Equal(t, argUpdate.ParentID, message3.ParentID)

	require.WithinDuration(t, message2.CreatedAt.Time, message3.CreatedAt.Time, time.Second)
}

func TestDeleteMessage(t *testing.T) {
	message1 := createRandomMessage(t)

	err := testQueries.DeleteMessage(context.Background(), message1.ID)
	require.NoError(t, err)

	message2, err := testQueries.GetMessage(context.Background(), message1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, message2)
}

func TestListMessages(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomMessage(t)
	}

	arg := ListMessagesParams{
		Limit: 5,
		Offset: 10,
	}

	messages, err := testQueries.ListMessages(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, messages, 5)

	for _, message := range messages {
		require.NotEmpty(t, message)
	}
}

