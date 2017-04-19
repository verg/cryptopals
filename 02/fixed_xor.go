package cyrptopals

import (
	"encoding/hex"
	"errors"
)

func fixedXor(a, b string) (string, error) {
	aDecoded, _ := hex.DecodeString(a)
	bDecoded, _ := hex.DecodeString(b)
	if len(aDecoded) != len(bDecoded) {
		return "", errors.New("Hex strings must be equal length")
	}

	res := make([]byte, len(aDecoded))
	for i := range aDecoded {
		res[i] = aDecoded[i] ^ bDecoded[i]
	}
	return hex.EncodeToString(res), nil
}
