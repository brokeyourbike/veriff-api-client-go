package veriff

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func SignPayload(secret, payload string) string {
	digest := hmac.New(sha256.New, []byte(secret))
	digest.Write([]byte(payload))
	return fmt.Sprintf("%x", digest.Sum(nil))
}
