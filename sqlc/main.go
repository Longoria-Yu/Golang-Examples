package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	db "github.com/Longoria-Yu/Golang-Examples/sqlc/db/sqlc"
)

// For connecting to the postgres database, we need to start the postgres server first.
// To start the postgres server in a container, run the following command:
// 1. docker run --name postgres -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=example -d -p 5432:5432 postgres:14-alpine
// 2. psql -h localhost -p 5432 -U user -d example

const dbSource = "postgresql://user:password@localhost:5432/example"

func main() {
	store, err := connectDB()
	if err != nil {
		log.Fatal("cannot connect to the example db", err)
	}

	users, getUsersErr := store.GetUsers(context.Background())
	if getUsersErr != nil {
		log.Fatal("cannot get the users from the database", getUsersErr)
	}
	log.Println("The current users in the database are:", len(users))

	// Use the store to interact with the database
	user, createUserErr := store.CreateUser(context.Background(), db.CreateUserParams{
		Username: "user2",
		Password: "password2",
		Email:    "user2@gmail.com",
		Mood:     db.NullEnumMood{EnumMood: db.EnumMoodHappy, Valid: true},
	})
	if createUserErr != nil {
		log.Fatal("cannot create the user in the database", createUserErr)
	}
	log.Println("The user created in the database is:", user)
}

func connectDB() (*db.Queries, error) {
	ctx := context.Background()
	connPool, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal("failed to cannot connect to the example db", err)
	}

	store := db.New(connPool)
	return store, err
}
