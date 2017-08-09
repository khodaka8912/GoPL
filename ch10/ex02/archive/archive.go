package archive

import (
	"bytes"
	"errors"
	"os"
)

var unarchivers = []unarchiver{}

var maxSignature = 0

type Reader interface {
	Next() (os.File, error)
	Read(buf []byte) (int, error)
}

type unarchiver struct {
	newReader func(string) (Reader, error)
	signature []byte
}

func NewReader(name string) (Reader, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	sign, err := readSignature(file)
	if err != nil {
		return nil, err
	}
	for _, u := range unarchivers {
		if bytes.Contains(sign, u.signature) {
			return u.newReader(name)
		}
	}
	return nil, errors.New("unknown file format")
}

func readSignature(file *os.File) ([]byte, error) {
	defer file.Close()
	buf := make([]byte, maxSignature)
	for read := 0; read < maxSignature; {
		n, err := file.Read(buf[read:])
		if err != nil {
			return nil, err
		}
		read += n
	}
	return buf, nil
}

func Register(newReader func(name string) Reader, signature []byte) {

}
