package main

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

func TestSha256(t *testing.T) {
	hash := sha256.New()
	in := bytes.NewReader([]byte("x"))
	if sum := sum(in, hash); len(sum) != 64 {
		t.Errorf("sum(%T, %T) wants hex string of length 64 but length is %d", in, hash, len(sum))
	}
}

func TestSha384(t *testing.T) {
	hash := sha512.New384()
	in := bytes.NewReader([]byte("x"))
	if sum := sum(in, hash); len(sum) != 96 {
		t.Errorf("sum(%T, %T) wants hex string of length 96 but length is %d", in, hash, len(sum))
	}
}

func TestSha512(t *testing.T) {
	hash := sha512.New()
	in := bytes.NewReader([]byte("x"))
	if sum := sum(in, hash); len(sum) != 128 {
		t.Errorf("sum(%T, %T) wants hex string of length 128 but length is %d", in, hash, len(sum))
	}
}
