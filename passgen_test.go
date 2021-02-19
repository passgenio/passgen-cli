package passgen

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassGen(t *testing.T) {
	p := &PassGen{
		Master:    []byte("master"),
		Algorithm: &AES{},
	}
	unique := "unique"
	out1, err := p.GenFor(unique)
	assert.NoError(t, err)

	out2, err := p.Algorithm.Decrypt(out1, p.Master)

	assert.NoError(t, err)
	assert.Equal(t, unique, fmt.Sprintf("%x", out2))
}
