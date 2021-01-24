package services

import (
	"github.com/galileoluna/bookstore_users-api/utils/errors"

	"github.com/galileoluna/bookstore_users-api/domain/users"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
