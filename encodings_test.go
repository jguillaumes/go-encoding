package main

import (
	"fmt"
	"testing"

	"github.com/jguillaumes/go_encoding/encodings"
)

var ebcdic_digits = []byte{0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9}
var string_digits = "0123456789"

var string_lowercase = "abcdefghijklmnopqrstuvwxyz"
var ebcdic_lowercase = []byte{0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0xa2, 0xa3, 0xa4, 0xa5, 0xa6, 0xa7, 0xa8, 0xa9}

var string_uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var ebcdic_uppercase = []byte{0xc1, 0xc2, 0xc3, 0xc4, 0xc5, 0xc6, 0xc7, 0xc8, 0xc9, 0xd1, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xd8, 0xd9, 0xe2, 0xe3, 0xe4, 0xe5, 0xe6, 0xe7, 0xe8, 0xe9}

func TestEncoding(t *testing.T) {

	enc := encodings.NewEncoding()

	codes := enc.ListEncodings()

	for _, c := range codes {
		d, err := enc.DecodeBytes(ebcdic_digits, c)
		if err != nil {
			t.Error("Failed to decode digits array")
		} else if d != string_digits {
			t.Errorf("Wrong decoding of digits for %s. Expected %s, got %s", c, string_digits, d)
		} else {
			fmt.Printf("Digits decoding test for %s OK\n", c)
		}
	}

	for _, c := range codes {
		d, err := enc.DecodeBytes(ebcdic_lowercase, c)
		if err != nil {
			t.Error("Failed to decode lowercase array")
		} else if d != string_lowercase {
			t.Errorf("Wrong decoding of lowercase letters for %s. Expected %s, got %s", c, string_lowercase, d)
		} else {
			fmt.Printf("Lowercase decoding test for %s OK\n", c)
		}
	}

	for _, c := range codes {
		d, err := enc.DecodeBytes(ebcdic_uppercase, c)
		if err != nil {
			t.Error("Failed to uppercase array")
		} else if d != string_uppercase {
			t.Errorf("Wrong decoding of lowercase letters for %s. Expected %s, got %s", c, string_uppercase, d)
		} else {
			fmt.Printf("Uppercase decoding test for %s OK\n", c)
		}
	}

	e, _ := enc.EncodeString(string_uppercase, "IBM-037")
	for _, b := range e {
		fmt.Printf("0x%02x,", b)
	}
}
