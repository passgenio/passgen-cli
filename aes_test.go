package passgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAES(t *testing.T) {
	a := &AES{}
	data := "Amirreza"
	key := []byte("secret_key_should_be_32_bytes_of")
	encrypted, err := a.Encrypt([]byte(data), key)
	assert.NoError(t, err)

	decrypted, err := a.Decrypt(encrypted, key)
	assert.NoError(t, err)
	assert.Equal(t, data, string(decrypted))
}
