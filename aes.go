package passgen

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"strings"
)

type AES struct{}

func (a *AES) fillKey(key []byte) []byte {
	if len(key) == 32 {
		return key
	}
	if len(key) > 32 {
		return key[:32]
	}
	return append(key, []byte(strings.Repeat("0", 32-len(key)))...)
}

func (a *AES) Encrypt(plaintext []byte, key []byte) (ciphertext []byte, err error) {
	key = a.fillKey(key)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func (a *AES) Decrypt(ciphertext []byte, key []byte) (plaintext []byte, err error) {
	key = a.fillKey(key)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("malformed ciphertext")
	}

	return gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
}
