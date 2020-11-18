package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func ComputeHmacSha256(src, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(src))


	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
