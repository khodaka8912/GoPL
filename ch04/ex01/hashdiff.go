package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	s1, s2 := "x", "X"
	args := os.Args[1:]
	if len(args) == 2 {
		s1, s2 = args[0], args[1]
	}
	c1 := sha256.Sum256([]byte(s1))
	c2 := sha256.Sum256([]byte(s2))
	diff := diff(&c1, &c2)
	fmt.Printf("SHA256 diff bits between %q and %q : %d", s1, s2, diff)
}

func diff(hash1, hash2 *[32]byte) int {
	diff := 0
	for i := 0; i < 32; i++ {
		diff += popCount(hash1[i] ^ hash2[i])
	}
	return diff
}

func popCount(bits byte) int {
	bits = (bits & 0x55) + (bits >> 1 & 0x55)
	bits = (bits & 0x33) + (bits >> 2 & 0x33)
	return int((bits & 0x0f) + (bits >> 4 & 0x0f))
}
