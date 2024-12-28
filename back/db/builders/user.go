package builders

import (
	"context"

	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/TheRoniOne/Kracker/internal"
	"github.com/brianvoe/gofakeit/v7"
)

type UserBuilder struct {
	Queries *sqlc.Queries
	Params  sqlc.CreateUserParams
}

func NewUserBuilder(queries *sqlc.Queries) *UserBuilder {
	return &UserBuilder{Queries: queries, Params: sqlc.CreateUserParams{
		Username:   gofakeit.Username(),
		Email:      gofakeit.Email(),
		SaltedHash: gofakeit.UUID(),
		Firstname:  gofakeit.FirstName(),
		Lastname:   gofakeit.LastName(),
	}}
}

func (b *UserBuilder) Username(username string) *UserBuilder {
	b.Params.Username = username

	return b
}

func (b *UserBuilder) Password(password string) *UserBuilder {
	saltedHash, err := internal.CreateSaltedHash(password)
	if err != nil {
		panic(err)
	}
	b.Params.SaltedHash = saltedHash

	return b
}

func (b *UserBuilder) CreateOne() sqlc.User {
	User, err := b.Queries.CreateUser(
		context.Background(),
		b.Params)
	if err != nil {
		panic(err)
	}

	return User
}

func (b *UserBuilder) CreateMany(count int) ([]sqlc.User, error) {
	var c []sqlc.User
	for i := 0; i < count; i++ {
		User := b.CreateOne()
		c = append(c, User)
	}
	return c, nil
}
