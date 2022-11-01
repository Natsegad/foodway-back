package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidatePassword(t *testing.T) {
	pass := "1"

	assert.NoError(t, ValidatePasswordOzzo(pass))
}
