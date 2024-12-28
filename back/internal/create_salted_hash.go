package internal

import (
	"github.com/alexedwards/argon2id"
)

func CreateSaltedHash(password string) (string, error) {
	params := &argon2id.Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	return argon2id.CreateHash(password, params)
}
