package webhook

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
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
	var pass bool
	for _, secret := range v.secrets {
		digest := hmac.New(sha256.New, []byte(secret))
		digest.Write([]byte(message))
		computed := fmt.Sprintf("%x", digest.Sum(nil))

		if computed == signature {
			pass = true
			break
		}
	}

	if pass {
		return nil
	}

	return fmt.Errorf("signature missmatch want %s", signature)
}
