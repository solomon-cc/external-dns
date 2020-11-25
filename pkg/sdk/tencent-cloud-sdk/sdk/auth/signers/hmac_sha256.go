package signers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func ComputeHmacSha256(source, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(source))


	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
