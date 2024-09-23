package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Transaction struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Transaction {
	return &Transaction{
		db: db, 
		Queries: New(db),
	}
}

func (transaction *Transaction) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := transaction.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb error: %v", err, rbErr)
		}

		return err
	}

	return tx.Commit()
}

type AnswerMessageParams struct {
	ParentMessage int64 "json:parent_message_id"
	Body string "json:message_body"
	UserID int64 "json:user_id"
	RoomID int64 "json:room_id" 
}

type AnswerMessageResult struct {
	ParentMessage Message "json:parent_message"
	Message Message "json:message"
}


func (transaction *Transaction) AnswerTx( ctx context.Context, arg AnswerMessageParams) (AnswerMessageResult, error) {
	var result AnswerMessageResult

	err := transaction.execTx(ctx, func(q *Queries) error {
		var err error

		ParentID := fmt.Sprintf("%d", arg.ParentMessage)
		result.Message, err = q.CreateMessage(ctx, CreateMessageParams{
			Message: sql.NullString{String: arg.Body, Valid: true},
			UserID: arg.UserID,
			RoomID: arg.RoomID,
			ParentID: sql.NullString{String: ParentID, Valid: true},
		})
		if err != nil {
			return nil
		}

		result.ParentMessage, err = q.UpdateAnswered(ctx, UpdateAnsweredParams{
			ID: arg.ParentMessage,
			Answered: sql.NullBool{Bool: true, Valid: true},
		})
		if err != nil {
			return nil
		}

		return nil
	})

	return result, err
}