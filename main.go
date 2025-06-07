package main

import (
	"fmt"

	enc "github.com/jguillaumes/go_encoding/encodings"
)

func main() {
	encoding := enc.NewEncoding()
	_, err := encoding.GetEncodingMapFor("IBM-1047")
	if err != nil {
		panic(err)
	}
	_, err = encoding.GetEncodingMapFor("IBM-1145")
	if err != nil {
		panic(err)
	}
	fmt.Println(encoding)
	fmt.Println(encoding.ListEncodings())

}
