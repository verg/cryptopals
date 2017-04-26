package main

import (
	"bufio"
	sbx "cryptopals/single_byte_xor"
	"fmt"
	"io"
	"os"
)

func main() {
	input, _ := os.Open("04.txt")
	res := findBestSingleByteXorCandidate(input)
	fmt.Printf("res = %+v\n", res)
}

func findBestSingleByteXorCandidate(r io.Reader) string {
	sc := bufio.NewScanner(r)

	bestScore := 0
	var best string
	for sc.Scan() {
		line := sc.Bytes()
		candidateString := sbx.SingleByteXorCipher(string(line))
		score := sbx.EnglishScore([]byte(candidateString))
		if score > bestScore {
			bestScore = score
			best = candidateString
		}
	}
	return best
}
