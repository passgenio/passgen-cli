package passgen

import (
	"crypto/hmac"
	"crypto/sha512"
)

type SHA512 struct{}

func (s *SHA512) Hash(data []byte, key []byte) ([]byte, error) {
	h := hmac.New(sha512.New512_256, key)
	h.Write(data)
	return h.Sum(nil), nil
}
