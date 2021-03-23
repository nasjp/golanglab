package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"time"
	"unicode/utf8"
)

const keyLength = 32

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	ap := fmt.Sprintf("%s%v", "aaaa", time.Now())
	bp := fmt.Sprintf("%s%v", "bbbb", time.Now())
	cp := fmt.Sprintf("%s%v", "cccc", time.Now())

	for _, p := range []string{ap, bp, cp} {
		hash := getBinaryBySHA256(p)
		hashStr := fmt.Sprintf("%x", hash)
		fmt.Printf("pass  : %s\n", ap)
		fmt.Printf("hash  : %x\n", hashStr)
		fmt.Printf("cur   : %x\n", hashStr)
		fmt.Printf("len   : %x\n", utf8.RuneCountInString(hashStr))
	}

	return nil
}

func getBinaryBySHA256(s string) []byte {
	r := sha256.Sum256([]byte(s))
	return r[:]
}
