// Definition of application errors.
//
// @author TSS

package domain

import (
	"errors"
)

var (
	ErrInvalidHmac = errors.New("invalid HMAC")
	ErrInvalidPass = errors.New("invalid password")
)
