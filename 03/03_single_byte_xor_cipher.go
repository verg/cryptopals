package main

import (
	"cryptopals/single_byte_xor"
	"fmt"
)

func main() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	res := single_byte_xor.SingleByteXorCipher(input)
	fmt.Println(res)
}
