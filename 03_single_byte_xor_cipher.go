package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	res := SingleByteXorCipher(input)
	fmt.Println(res)
}

func SingleByteXorCipher(hexString string) string {
	var maxString []byte
	maxScore := 0

	h, _ := hex.DecodeString(hexString)
	for i := 0; i < 256; i++ {
		decoded := xorWithByte(h, byte(i))
		score := score(decoded)
		if score > maxScore {
			maxString = decoded
			maxScore = score
		}
	}
	return string(maxString)
}

func xorWithByte(encoded []byte, b byte) (res []byte) {
	res = make([]byte, len(encoded))
	for i, a := range encoded {
		res[i] = a ^ b
	}
	return res
}

func score(s []byte) (score int) {
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
