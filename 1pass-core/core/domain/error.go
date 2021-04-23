// Definition of application errors.
//
// @author TSS

package domain

import (
	"errors"
)

var (
	ErrInvalidHmac    = errors.New("invalid HMAC")
	ErrInvalidPass    = errors.New("invalid password")
	ErrInvalidVault   = errors.New("invalid vault (%v)")
	ErrUnknownItemCat = errors.New("unknown item category")
)
