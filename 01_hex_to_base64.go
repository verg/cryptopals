package main

import (
	"encoding/base64"
	"encoding/hex"
)

func hexToBase64(hexString string) string {
	decoded, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(decoded)
}
