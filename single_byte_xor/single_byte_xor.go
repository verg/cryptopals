package single_byte_xor

import "encoding/hex"

func SingleByteXorCipher(hexString string) string {
	var maxString []byte
	maxScore := 0

	h, _ := hex.DecodeString(hexString)
	for i := 0; i < 256; i++ {
		decoded := XorWithByte(h, byte(i))
		score := EnglishScore(decoded)
		if score > maxScore {
			maxString = decoded
			maxScore = score
		}
	}
	return string(maxString)
}

func XorWithByte(encoded []byte, b byte) (res []byte) {
	res = make([]byte, len(encoded))
	for i, a := range encoded {
		res[i] = a ^ b
	}
	return res
}

func EnglishScore(s []byte) (score int) {
	for _, b := range s {
		if freqMap[b] {
			score += 1
		}
	}
	return
}

var freqMap = map[byte]bool{
	' ': true,
	'e': true,
	'E': true,
	't': true,
	'T': true,
	'a': true,
	'A': true,
	'o': true,
	'O': true,
	'i': true,
	'I': true,
	'n': true,
	'N': true,
	's': true,
	'S': true,
	'h': true,
	'H': true,
	'r': true,
	'R': true,
	'd': true,
	'D': true,
	'l': true,
	'L': true,
	'u': true,
	'U': true,
}
