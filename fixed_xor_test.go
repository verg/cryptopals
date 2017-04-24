package main

import (
	"fmt"
	"testing"
)

func TestFixedXor(t *testing.T) {
	a := "1c0111001f010100061a024b53535009181c"
	b := "686974207468652062756c6c277320657965"
	expected := "746865206b696420646f6e277420706c6179"

	if got, err := fixedXor(a, b); err != nil || got != expected {
		if err != nil {
			t.Errorf(fmt.Sprint(err))
		} else {
			t.Errorf("Got: %s, Expected: %s", got, expected)
		}
	}
}
