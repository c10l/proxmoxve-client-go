package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalJSON(t *testing.T) {
	var b PVEBool
	err := b.UnmarshalJSON([]byte("1"))
	assert.NoError(t, err)
	assert.Equal(t, true, bool(b))
	err = b.UnmarshalJSON([]byte("0"))
	assert.NoError(t, err)
	assert.Equal(t, false, bool(b))
	err = b.UnmarshalJSON([]byte("2"))
	assert.Error(t, err)
}
