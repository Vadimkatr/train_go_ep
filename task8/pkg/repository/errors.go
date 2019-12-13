package repository

import (
	"errors"
)

var (
	ErrContactListIsEmpty    = errors.New("email or phone are not unique")
	ErrNotUniqueEmailOrPhone = errors.New("email or phone are not unique")
	ErrRecordNotFound        = errors.New("record not found")
)