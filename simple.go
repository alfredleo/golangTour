package main

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"fmt"
	mathrand "math/rand"
)

func main() {
	var b [8]byte
	_, err := cryptorand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	mathrand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
	for _, _ = range "ABCDEFGHIJKLMNOP" {
		//fmt.Println(i, c)
		fmt.Println("My favorite number is: ", mathrand.Intn(10))
	}
}
