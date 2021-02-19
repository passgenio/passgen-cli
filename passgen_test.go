package passgen

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassGen(t *testing.T) {
	p := &PassGen{
		Master:    []byte("master"),
		Algorithm: &SHA512{},
	}
	unique := "unique"
	out1, err := p.GenFor(unique)
	assert.NoError(t, err)

	out2, err := p.GenFor(unique)

	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%x", out1), fmt.Sprintf("%x", out2))
}
