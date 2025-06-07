package encodings

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"path"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

//go:embed tables
var encoding_tables embed.FS

const table_regex = "0x([0-9a-f]{2})\\s+0x([0-9a-f]{4})\\s+.*"

type EncodingMap map[rune]byte
type DecodingTable []rune

type EncodingData struct {
	name          string
	encodingMap   *EncodingMap
	decodingTable *DecodingTable
}

type Encoding interface {
	GetEncodingMapFor(string) (*EncodingMap, error)
	GetDecodingTableFor(string) (*DecodingTable, error)
	EncodeString(string, string) ([]byte, error)
	DecodeBytes([]byte, string) (string, error)
	EncodeRune(rune, string) (byte, error)
	ListEncodings() []string

	generateEncoding(string) (*EncodingData, error)
}

type encodingImpl struct {
	Encoding
	encodings map[string]EncodingData
}

func (e *encodingImpl) GetDecodingTableFor(name string) (*DecodingTable, error) {
	var err error = nil
	ed, ok := e.encodings[name]
	if !ok {
		var edPtr *EncodingData
		edPtr, err = e.generateEncoding(name)
		if err == nil && edPtr != nil {
			ed = *edPtr
		}

	}
	if err != nil {
		return nil, err
	}
	return ed.decodingTable, nil
}

func (e *encodingImpl) GetEncodingMapFor(name string) (*EncodingMap, error) {
	var err error = nil
	ed, ok := e.encodings[name]
	if !ok {
		var edPtr *EncodingData
		edPtr, err = e.generateEncoding(name)
		if err == nil && edPtr != nil {
			ed = *edPtr
		}

	}
	if err != nil {
		return nil, err
	}
	return ed.encodingMap, nil
}

func (e *encodingImpl) DecodeBytes(bs []byte, code string) (string, error) {
	decoder, err := e.GetDecodingTableFor(code)
	if err != nil {
		return "", err
	}
	var builder strings.Builder
	for _, b := range bs {
		r := (*decoder)[b]
		builder.WriteRune(r)
	}
	return builder.String(), nil
}

func (e *encodingImpl) EncodeString(s string, code string) ([]byte, error) {
	tmpBuf := bytes.NewBuffer(make([]byte, 0, len(s)*2))
	encoder, err := e.GetEncodingMapFor(code)
	if err != nil {
		return nil, err
	}
	// Loop over the runes in s
	var r rune
	const blank rune = ' '
	for _, r = range s {
		b, ok := (*encoder)[r]
		if !ok {
			log.Warnf("No encoding for the rune %v in %s, set to an encoded space\n", r, code)

			b = (*encoder)[blank]
		}
		tmpBuf.WriteByte(b)
	}
	return tmpBuf.Bytes()[0:tmpBuf.Len()], nil
}

func (e *encodingImpl) EncodeRune(r rune, code string) (byte, error) {
	encoder, err := e.GetEncodingMapFor(code)
	if err != nil {
		return 0, err
	}
	b, ok := (*encoder)[r]
	if ok {
		return b, nil
	} else {
		return 0, fmt.Errorf("there is no encoding for the rune %v in the %s codemap", r, code)
	}
}

func (e *encodingImpl) generateEncoding(name string) (*EncodingData, error) {
	re, _ := regexp.Compile(table_regex)
	fileName := path.Join("tables", name+".txt")
	log.Debugf("Building the encoding table %s from %s\n", name, fileName)
	encodingMap := make(EncodingMap)
	decodingTable := make(DecodingTable, 256)
	tableData, err := encoding_tables.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(tableData))
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] != '#' {
			parts := re.FindStringSubmatch(line)
			if len(parts) != 3 {
				log.Printf("Malformed/unparsable line: [%s]", line)
			} else {
				i, _ := strconv.ParseUint(parts[1], 16, 8)
				v, _ := strconv.ParseUint(parts[2], 16, 16)
				decodingTable[i] = rune(v)
				encodingMap[rune(v)] = byte(i)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	encodingData := &EncodingData{
		name:          name,
		encodingMap:   &encodingMap,
		decodingTable: &decodingTable,
	}
	e.encodings[name] = *encodingData
	return encodingData, nil
}

func (e *encodingImpl) ListEncodings() []string {
	files, _ := encoding_tables.ReadDir("tables")
	list := make([]string, 0, len(files))
	for _, f := range files {
		n := f.Name()
		if strings.HasSuffix(n, ".txt") {
			list = append(list, strings.TrimSuffix(n, ".txt"))
		}
	}
	return list
}

func NewEncoding() Encoding {
	impl := &encodingImpl{
		encodings: make(map[string]EncodingData),
	}
	return impl
}
