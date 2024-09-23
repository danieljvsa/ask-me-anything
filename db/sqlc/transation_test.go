package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/danieljvsa/ask-me-anything/util"
	"github.com/stretchr/testify/require"
)

func TestAnswerTx(t *testing.T) {
	store := NewStore(testDB)

	message1 := createRandomMessage(t)
	
	argUser := CreateUserParams{
		Username: util.RandomNameString(), // randomly generated?
		Password: util.RandomNameString(),
		Email: util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), argUser)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	n := 5 

	errs := make(chan error)
	results := make(chan AnswerMessageResult)

	for i := 0; i < n; i++ {
		
		go func() {
			arg := AnswerMessageParams{
				UserID: user.ID,
				Body: util.RandomMessageString(),
				ParentMessage: message1.ID,
				RoomID: message1.RoomID,
			}

			result, err := store.AnswerTx(context.Background(), arg)
			require.NoError(t, err)
			require.NotEmpty(t, result)

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		parentMessage := result.ParentMessage
		message := result.Message

		require.NotEmpty(t, parentMessage)
		require.NotEmpty(t, message)
		
		parentId := fmt.Sprintf("%d", parentMessage.ID)
		if message.ParentID.Valid {
			require.Equal(t, parentId, message.ParentID.String)
		} else {
			return
		}

		require.Equal(t, user.ID, message.UserID)
		require.Equal(t, parentMessage.RoomID, message.RoomID)

		_, err = store.GetMessage(context.Background(), message.ID)
		require.NoError(t, err)

	}
}