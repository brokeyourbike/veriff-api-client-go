package webhook_test

import (
	"context"
	"testing"

	"github.com/brokeyourbike/veriff-api-client-go/webhook"
	"github.com/stretchr/testify/assert"
)

func TestVerifier(t *testing.T) {
	message := `{"verification":{"callback":"https://veriff.com","person":{"firstName":"John","lastName":"Smith"},"document":{"type":"PASSPORT","country":"EE"},"vendorData":"unique id of the end-user","timestamp":"2016-05-19T08:30:25.597Z"}}`
	signature := "0dcab73ddd20062616d104231c7439657546a5c24e4691977da93bb854c31e25"

	v := webhook.NewVerifier([]string{"abc", "abcdef12-abcd-abcd-abcd-abcdef012345", "cde"})
	assert.NoError(t, v.Verify(context.Background(), []byte(message), signature))
}
