package archive

import (
	"os"
)

type Reader interface {
	Next() (os.File, error)
	Read(buf []byte) (int, error)
}

func Register(factory func(name string) Reader, signature []byte) {

}
