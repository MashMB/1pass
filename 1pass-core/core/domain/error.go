// Definition of application errors.
//
// @author TSS

package domain

import (
	"errors"
)

var (
	ErrInvalidMd5      = errors.New("invalid MD5 checksum")
	ErrInvalidHmac     = errors.New("invalid HMAC")
	ErrInvalidPass     = errors.New("invalid password")
	ErrInvalidVault    = errors.New("invalid vault (%v)")
	ErrNoUpdate        = errors.New("up to date")
	ErrUnknownDataType = errors.New("unknown data type")
	ErrUnknownItemCat  = errors.New("unknown item category")
	ErrUnknownLogicVal = errors.New("unknown logic value")
)
