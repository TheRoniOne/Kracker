package factories

import (
	"context"

	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/brianvoe/gofakeit/v7"
)

type UserFactory struct {
	Queries *sqlc.Queries
}

func (f *UserFactory) CreateOne() sqlc.User {
	User, err := f.Queries.CreateUser(
		context.Background(),
		sqlc.CreateUserParams{
			Username:   gofakeit.Username(),
			Email:      gofakeit.Email(),
			SaltedHash: gofakeit.UUID(),
			Firstname:  gofakeit.FirstName(),
			Lastname:   gofakeit.LastName(),
		})
	if err != nil {
		panic(err)
	}

	return User
}

func (f *UserFactory) CreateMany(count int) ([]sqlc.User, error) {
	var c []sqlc.User
	for i := 0; i < count; i++ {
		User := f.CreateOne()
		c = append(c, User)
	}
	return c, nil
}
