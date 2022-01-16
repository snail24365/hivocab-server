package db

import (
	"context"
	"database/sql"
	"fmt"
)
type Store interface {

}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return SQLStore{
		db: db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
}

// TransferTxResult is the result of the transfer transaction
type Exercise struct {
	Word 			Word 				`json:"word"`
	Usecase 	Usecase 		`json:"usecase"`
	Examples  []Example   `json:"examples"`
}


// TransferTx performs a money transfer from one account to the other.
// It creates the transfer, add account entries, and update accounts' balance within a database transaction
func (store *SQLStore) GetExercise(ctx context.Context, userId int64) (Exercise, error) {
	var exercise Exercise
	
	//exercise.Word = 


	return exercise, nil
}