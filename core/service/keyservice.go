// Default implementation of key service.
//
// @author TSS

package service

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/binary"
	"io"

	"github.com/mashmb/1pass/core/domain"
	"github.com/mashmb/1pass/port/out"
)

type dfltKeyService struct {
	cryptoUtils out.CrytpoUtils
	profileRepo out.ProfileRepo
}

func NewDfltKeyService(cryptoUtils out.CrytpoUtils, profileRepo out.ProfileRepo) *dfltKeyService {
	return &dfltKeyService{
		cryptoUtils: cryptoUtils,
		profileRepo: profileRepo,
	}
}

func (s *dfltKeyService) CheckHmac(msg, key, desiredHmac []byte) error {
	hash := hmac.New(sha256.New, key)
	size, err := hash.Write(msg)

	if err != nil {
		return err
	}

	if size != len(msg) {
		return io.ErrShortWrite
	}

	computed := hash.Sum(nil)

	if !hmac.Equal(computed, desiredHmac) {
		return domain.ErrInvalidHmac
	}

	return nil
}

func (s *dfltKeyService) DecodeData(key, initVector, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, initVector)
	mode.CryptBlocks(data, data)

	return data, nil
}

func (s *dfltKeyService) DecodeKeys(key, derivedKey, derivedMac []byte) ([]byte, []byte, error) {
	base, err := s.DecodeOpdata(key, derivedKey, derivedMac)

	if err != nil {
		return nil, nil, err
	}

	hash := sha512.New()

	if _, err := hash.Write(base); err != nil {
		return nil, nil, err
	}

	keys := hash.Sum(nil)

	return keys[:32], keys[32:64], nil
}

func (s *dfltKeyService) DecodeOpdata(cipherText, key, macKey []byte) ([]byte, error) {
	data, mac := cipherText[:len(cipherText)-32], cipherText[len(cipherText)-32:]

	if err := s.CheckHmac(data, macKey, mac); err != nil {
		return nil, err
	}

	plain, err := s.DecodeData(key, data[16:32], data[32:])

	if err != nil {
		return nil, err
	}

	var plainSize int
	reader := bytes.NewReader(plain[8:16])

	if err := binary.Read(reader, binary.LittleEndian, &plainSize); err != nil {
		return nil, err
	}

	return plain[len(plain)-plainSize:], nil
}

func (s *dfltKeyService) DerivedKeys(password string) ([]byte, []byte, error) {
	iterations := s.profileRepo.GetIterations()
	salt, err := base64.StdEncoding.DecodeString(s.profileRepo.GetSalt())

	if err != nil {
		return nil, nil, err
	}

	keys := s.cryptoUtils.DeriveKey([]byte(password), []byte(salt), iterations, 64, sha512.New)

	return keys[:32], keys[:32], nil
}

func (s *dfltKeyService) MasterKeys(derivedKey, derivedMac []byte) ([]byte, []byte, error) {
	encoded, err := base64.StdEncoding.DecodeString(s.profileRepo.GetMasterKey())

	if err != nil {
		return nil, nil, err
	}

	return s.DecodeKeys(encoded, derivedKey, derivedMac)
}

func (s *dfltKeyService) OverviewKeys(derivedKey, derivedMac []byte) ([]byte, []byte, error) {
	encoded, err := base64.StdEncoding.DecodeString(s.profileRepo.GetOverviewKey())

	if err != nil {
		return nil, nil, err
	}

	return s.DecodeKeys(encoded, derivedKey, derivedMac)
}
