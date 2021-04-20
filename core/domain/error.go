// Definition of application errors.
//
// @author TSS

package domain

import (
	"errors"
)

var (
	ErrInvalidHmac    = errors.New("invalid HMAC")
	ErrUnknownItemCat = errors.New("unknown item category")
	ErrInvalidPass    = errors.New("invalid password")
)
