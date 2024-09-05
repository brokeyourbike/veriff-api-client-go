package veriff_test

import (
	"testing"

	"github.com/brokeyourbike/veriff-api-client-go"
	"github.com/stretchr/testify/assert"
)

func TestSignPayload(t *testing.T) {
	signed1 := veriff.SignPayload("secret", "")
	signed2 := veriff.SignPayload("secret", string([]byte{}))

	assert.Equal(t, signed1, signed2)
}
