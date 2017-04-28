package main

import "encoding/hex"

func RepeatingKeyXOR(expr string, keys []byte) string {
	res := []byte(expr)
	keyIndex := 0
	numKeys := len(keys)
	for i := 0; i < len(expr); i++ {
		res[i] = res[i] ^ keys[keyIndex]
		keyIndex = (keyIndex + 1) % numKeys
	}
	return hex.EncodeToString(res)
}
