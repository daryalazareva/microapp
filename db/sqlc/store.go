package db

import (
	"context"
	"database/sql"
	"fmt"
)

//Store provides all functions to execute db queries and transactions
//Embedding queries inside store. All individual query functions are provided by queries
// will be available to store
type Store struct {
	*Queries
	db *sql.DB
}

//Returns a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

//execTx executes a function within a database transaction
func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//get back new queries object
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

//It searches for the user in the database and if it finds one with given email
//the func updates its password
func (s *Store) ChangePasswordTx(ctx context.Context, email string, newpassword string) error {
	err := s.execTx(ctx, func(q *Queries) error {

		_, err := q.GetUserForUpdate(ctx, email)
		if err != nil {
			return err
		}

		_, err = q.UpdateUser(ctx, UpdateUserParams{
			Email:             email,
			EncryptedPassword: newpassword,
		})
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

func (s *Store) CreateRecordUser(ctx context.Context, arg CreateUserParams) (User, error) {

	u, err := s.Queries.CreateUser(ctx, arg)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (s *Store) GetRecordUser(ctx context.Context, email string) (User, error) {

	u, err := s.Queries.GetUser(ctx, email)
	if err != nil {
		return User{}, err
	}

	return u, nil
}
