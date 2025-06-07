# Codepage handling module for golang

This module allows to encode golang strings into single-byte codepages (like the different variants of EBCDIC) and to decode byte arrays to strings.

It's inspired in the `ebcdic` python module, and it uses the same mechanism to generate translation tables, as it is explained below.

## Building and using this module

To use it, just import into your golang module:

github.com/jguillaumes/go_encoding

Out of the box, it supports the IBM-037, IBM-284, IBM-1047 and IBM-1145 EBCDIC encodings. It's very easy to add new encodings, not limited to EBCDIC codepages. The only restriction is they must encode a character in one byte.

### General usage

The module offers a single public interface, Encoding, with the following methods:

- `EncodeString(gostring string, codepage string) ([]byte, error)`: Encodes the `gostring`string into a byte slice using the `codepage` translation table.
- `DecodeBytes(buffer []bytes, codepage string) (string, error)`: Decodes the contents of `buffer`into a golang string using the `codepage` translation table.
- `EncodeRune(r rune, codepage string) (byte, error)`: Encodes a single rune into a byte using the `codepage`translation table.
- `ListEncodings() []string`: Provides a list of the supported encodings.
- `GetEncodingMapFor(codepage string) (*EncodingMap, error)`: Returns a pointer to an EncodingMap object described below.
- `GetDecodingTableFor(codepage string) (*DecodingTable, error)`: Returns a pointer to a DecodingTable object described below.

### Encoding maps and decoding tables

`EncodingMap`is an alias for `map[rune]byte`. It can be used to perform your own encoding.

`DecodingTable` is an alias for []rune. It is a 256 element slice with the rune corresponding to each possible byte value, and can be used to perform your own decoding.

## Extending and customizing this module.

This version requires the codeset tables to be present under the path `./encoding/tables` at build time. I have plans to allow to add external tables, but at this time if you want to support additional encodings you have to add the tables and rebuild the module. 

To generate additional tables, you can use the `CodecMapper` utility you can find here:

 https://github.com/roskakori/CodecMapper

 Please note this is a Java utility, and you will need to have a JDK and the `ant` utility installed in your workstation to build it. Once you have built it, its use is pretty simple. Lets say we want to generate the table for IBM-500:

 ```
 $ java -jar CodecMapper.jar IBM-500
write "IBM-500.txt"
 ```

That's it! The utility generates the IBM-500.txt file. You just have to copy or move it into `./encodings/tables` and rebuild the module.
Kudos and thanks to Thomas Aglassinger for providing the CodecMapper tool. Any mistake in the use of his program in the context of this module is my fault, not his.



