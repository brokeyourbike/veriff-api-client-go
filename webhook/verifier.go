package webhook

import (
	"context"
	"fmt"

	"github.com/brokeyourbike/veriff-api-client-go"
)

type Verifier interface {
	Verify(ctx context.Context, message, signature string) error
}

type verifier struct {
	secrets []string
}

func NewVerifier(secrets []string) *verifier {
	return &verifier{secrets: secrets}
}

func (v *verifier) Verify(ctx context.Context, message, signature string) error {
	for _, secret := range v.secrets {
		if veriff.SignPayload(secret, message) == signature {
			return nil
		}
	}

	return fmt.Errorf("signature missmatch want %s", signature)
}
