package codec

import (
	"encoding/base32"
	"encoding/base64"
)

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64Decode(s string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func Base32Encode(s string) string {
	return base32.StdEncoding.EncodeToString([]byte(s))
}

func Base32Decode(s string) (string, error) {
	b, err := base32.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
