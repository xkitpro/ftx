package ftx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func sign(secret []byte, payload string) string {
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(payload))

	return hex.EncodeToString(mac.Sum(nil))
}
