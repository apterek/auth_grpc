package storage

import "errors"

var (
	ErrUserExists    = errors.New("user already exists")
	ErrUserNotFound1 = errors.New("user not found")
	ErrAppNotFound   = errors.New("app not found")
)
