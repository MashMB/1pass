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
	"encoding/binary"
	"io"
	"log"
)

type dfltKeyService struct {
}

func NewDfltKeyService() *dfltKeyService {
	return &dfltKeyService{}
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

func (s *dfltKeyService) DecodeOpdata(cipherText, key, macKey []byte) []byte {
	data, mac := cipherText[:len(cipherText)-32], cipherText[len(cipherText)-32:]
	s.CheckHmac(data, macKey, mac)
	plain := s.DecodeData(key, data[16:32], data[32:])
	var plainSize int
	reader := bytes.NewReader(plain[8:16])
	binary.Read(reader, binary.LittleEndian, &plainSize)

	return plain[len(plain)-plainSize:]
}
