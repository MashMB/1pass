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
	"log"

	"github.com/mashmb/1pass/port/out"
)

type dfltKeyService struct {
	profileRepo out.ProfileRepo
}

func NewDfltKeyService(profileRepo out.ProfileRepo) *dfltKeyService {
	return &dfltKeyService{
		profileRepo: profileRepo,
	}
}

func (s *dfltKeyService) CheckHmac(msg, key, desiredHmac []byte) {
	hash := hmac.New(sha256.New, key)
	size, err := hash.Write(msg)

	if err != nil {
		log.Fatalln(err)
	}

	if size != len(msg) {
		log.Fatalln(io.ErrShortWrite)
	}

	computed := hash.Sum(nil)

	if !hmac.Equal(computed, desiredHmac) {
		log.Fatalln("invalid HMAC")
	}
}

func (s *dfltKeyService) DecodeData(key, initVector, data []byte) []byte {
	block, err := aes.NewCipher(key)

	if err != nil {
		log.Fatalln(err)
	}

	mode := cipher.NewCBCDecrypter(block, initVector)
	mode.CryptBlocks(data, data)

	return data
}

func (s *dfltKeyService) DecodeKeys(key, derivedKey, derivedMac []byte) ([]byte, []byte) {
	base := s.DecodeOpdata(key, derivedKey, derivedMac)
	hash := sha512.New()

	if _, err := hash.Write(base); err != nil {
		log.Fatalln(err)
	}

	keys := hash.Sum(nil)

	return keys[:32], keys[32:64]
}

func (s *dfltKeyService) DecodeOpdata(cipherText, key, macKey []byte) []byte {
	data, mac := cipherText[:len(cipherText)-32], cipherText[len(cipherText)-32:]
	s.CheckHmac(data, macKey, mac)
	plain := s.DecodeData(key, data[16:32], data[32:])
	var plainSize int
	reader := bytes.NewReader(plain[8:16])
	binary.Read(reader, binary.LittleEndian, &plainSize)

	return plain[len(plain)-plainSize:]
}

func (s *dfltKeyService) MasterKeys(derivedKey, derivedMac []byte) ([]byte, []byte) {
	encoded, err := base64.StdEncoding.DecodeString(s.profileRepo.GetMasterKey())

	if err != nil {
		log.Fatalln(err)
	}

	return s.DecodeKeys(encoded, derivedKey, derivedMac)
}

func (s *dfltKeyService) OverviewKeys(derivedKey, derivedMac []byte) ([]byte, []byte) {
	encoded, err := base64.StdEncoding.DecodeString(s.profileRepo.GetOverviewKey())

	if err != nil {
		log.Fatalln(err)
	}

	return s.DecodeKeys(encoded, derivedKey, derivedMac)
}
