package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jguillaumes/go-encoding/encodings"
)

var ebcdic_digits = []byte{0xF0, 0xF1, 0xF2, 0xF3, 0xF4, 0xF5, 0xF6, 0xF7, 0xF8, 0xF9}
var ascii_digits = []byte{0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39}
var string_digits = "0123456789"

var string_lowercase = "abcdefghijklmnopqrstuvwxyz"
var ebcdic_lowercase = []byte{0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88, 0x89, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97, 0x98, 0x99, 0xa2, 0xa3, 0xa4, 0xa5, 0xa6, 0xa7, 0xa8, 0xa9}
var ascii_lowercase = []byte{0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7a}

var string_uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var ebcdic_uppercase = []byte{0xc1, 0xc2, 0xc3, 0xc4, 0xc5, 0xc6, 0xc7, 0xc8, 0xc9, 0xd1, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xd8, 0xd9, 0xe2, 0xe3, 0xe4, 0xe5, 0xe6, 0xe7, 0xe8, 0xe9}
var ascii_uppercase = []byte{0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5a}

var string_special = "ñÑçÇ[]{}@#|áéíóúàèòùÁÉÍÓÚÀÈÌÒÙ"
var ibm037_special = []byte{0x49, 0x69, 0x48, 0x68, 0xba, 0xbb, 0xc0, 0xd0, 0x7c, 0x7b, 0x4f, 0x45, 0x51, 0x55, 0xce, 0xde, 0x44, 0x54, 0xcd, 0xdd, 0x65, 0x71, 0x75, 0xee, 0xfe, 0x64, 0x74, 0x78, 0xed, 0xfd}
var ibm284_special = []byte{0x6a, 0x7b, 0x48, 0x68, 0x4a, 0x5a, 0xc0, 0xd0, 0x7c, 0x69, 0x4f, 0x45, 0x51, 0x55, 0xce, 0xde, 0x44, 0x54, 0xcd, 0xdd, 0x65, 0x71, 0x75, 0xee, 0xfe, 0x64, 0x74, 0x78, 0xed, 0xfd}
var ibm1047_special = []byte{0x49, 0x69, 0x48, 0x68, 0xad, 0xbd, 0xc0, 0xd0, 0x7c, 0x7b, 0x4f, 0x45, 0x51, 0x55, 0xce, 0xde, 0x44, 0x54, 0xcd, 0xdd, 0x65, 0x71, 0x75, 0xee, 0xfe, 0x64, 0x74, 0x78, 0xed, 0xfd}
var ibm1145_special = []byte{0x6a, 0x7b, 0x48, 0x68, 0x4a, 0x5a, 0xc0, 0xd0, 0x7c, 0x69, 0x4f, 0x45, 0x51, 0x55, 0xce, 0xde, 0x44, 0x54, 0xcd, 0xdd, 0x65, 0x71, 0x75, 0xee, 0xfe, 0x64, 0x74, 0x78, 0xed, 0xfd}
var iso8859_special = []byte{0xf1, 0xd1, 0xe7, 0xc7, 0x5b, 0x5d, 0x7b, 0x7d, 0x40, 0x23, 0x7c, 0xe1, 0xe9, 0xed, 0xf3, 0xfa, 0xe0, 0xe8, 0xf2, 0xf9, 0xc1, 0xc9, 0xcd, 0xd3, 0xda, 0xc0, 0xc8, 0xcc, 0xd2, 0xd9}
var euro_rune = '€'
var ibm1145_euro byte = 0x9f

func TestEncoding(t *testing.T) {

	enc := encodings.NewEncoding()

	codes := enc.ListEncodings()

	for _, c := range codes {
		var d string
		var err error
		if strings.HasPrefix(c, "IBM-") {
			d, err = enc.DecodeBytes(ebcdic_digits, c)
		} else {
			d, err = enc.DecodeBytes(ascii_digits, c)
		}
		if err != nil {
			t.Error("Failed to decode digits array")
		} else if d != string_digits {
			t.Errorf("Wrong decoding of digits for %s. Expected %s, got %s", c, string_digits, d)
		} else {
			fmt.Printf("Digits decoding test for %s OK\n", c)
		}
	}

	for _, c := range codes {
		var d string
		var err error
		if strings.HasPrefix(c, "IBM-") {
			d, err = enc.DecodeBytes(ebcdic_lowercase, c)
		} else {
			d, err = enc.DecodeBytes(ascii_lowercase, c)
		}
		if err != nil {
			t.Error("Failed to decode lowercase array")
		} else if d != string_lowercase {
			t.Errorf("Wrong decoding of lowercase letters for %s. Expected %s, got %s", c, string_lowercase, d)
		} else {
			fmt.Printf("Lowercase decoding test for %s OK\n", c)
		}
	}

	for _, c := range codes {
		var d string
		var err error
		if strings.HasPrefix(c, "IBM-") {
			d, err = enc.DecodeBytes(ebcdic_uppercase, c)
		} else {
			d, err = enc.DecodeBytes(ascii_uppercase, c)
		}
		if err != nil {
			t.Error("Failed decode to uppercase array")
		} else if d != string_uppercase {
			t.Errorf("Wrong decoding of uppercase letters for %s. Expected %s, got %s", c, string_uppercase, d)
		} else {
			fmt.Printf("Uppercase decoding test for %s OK\n", c)
		}
	}

	d, err := enc.DecodeBytes(ibm037_special, "IBM-037")
	if err != nil {
		t.Error("Failed decode of special array")
	} else if d != string_special {
		t.Errorf("Wrong decoding of special characters letters for IBM-037: Expected %s, got %s", string_special, d)
	} else {
		fmt.Printf("Special chars decoding test for IBM-037 OK\n")
	}

	d, err = enc.DecodeBytes(ibm284_special, "IBM-284")
	if err != nil {
		t.Error("Failed decode of special array")
	} else if d != string_special {
		t.Errorf("Wrong decoding of special characters letters for IBM-284: Expected %s, got %s", string_special, d)
	} else {
		fmt.Printf("Special chars decoding test for IBM-284 OK\n")
	}

	d, err = enc.DecodeBytes(ibm1047_special, "IBM-1047")
	if err != nil {
		t.Error("Failed decode of special array")
	} else if d != string_special {
		t.Errorf("Wrong decoding of special characters letters for IBM-1047: Expected %s, got %s", string_special, d)
	} else {
		fmt.Printf("Special chars decoding test for IBM-1047 OK\n")
	}

	d, err = enc.DecodeBytes(ibm1145_special, "IBM-1145")
	if err != nil {
		t.Error("Failed decode of special array")
	} else if d != string_special {
		t.Errorf("Wrong decoding of special characters letters for IBM-1145: Expected %s, got %s", string_special, d)
	} else {
		fmt.Printf("Special chars decoding test for IBM-1145 OK\n")
	}

	d, err = enc.DecodeBytes(iso8859_special, "ISO8859-1")
	if err != nil {
		t.Error("Failed decode of special array")
	} else if d != string_special {
		difs, err := compareTwoStrings(d, string_special)
		if err != nil {
			t.Error(err)
		} else {
			if len(difs) != 0 {
				for _, d := range difs {
					fmt.Printf("Difference at position %d: expected %c, found %c\n", d.P, d.E, d.F)
				}
				t.Errorf("Wrong decoding of special characters letters for ISO8859-1: Expected \n[%s], got \n[%s]", string_special, d)
			}
		}
	} else {
		fmt.Printf("Special chars decoding test for ISO8859-1 OK\n")
	}

	b, err := enc.EncodeRune(euro_rune, "IBM-1145")
	if err != nil {
		t.Error("Failed decode of euro rune")
	} else if b != ibm1145_euro {
		t.Errorf("Wrong decoding of euro rune for IBM-1145: Expected 0x%02x, got 0x%02x", ibm1145_euro, b)
	} else {
		fmt.Printf("Euro char decoding test for IBM-1145 OK\n")
	}

}

func compareTwoStrings(s1 string, s2 string) ([]struct {
	P int
	E rune
	F rune
}, error) {
	if len(s1) != len(s2) {
		// return nil, fmt.Errorf("the two strings do not have the same length (%d vs %d) ", len(s1), len(s2))
	}

	diffs := make([]struct {
		P int
		E rune
		F rune
	}, 0)
	r1 := []rune(s1)
	r2 := []rune(s2)
	for i := range r2 {
		c1 := r1[i]
		c2 := r2[i]
		if c1 != c2 {
			diffs = append(diffs, struct {
				P int
				E rune
				F rune
			}{
				P: i,
				E: c1,
				F: c2,
			})
		}
	}
	return diffs, nil

}
