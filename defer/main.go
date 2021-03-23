package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}

var ErrFail = errors.New("fail")

func run() error {
	checker, ok := getChecker()

	defer checker()

	if isOK() {
		ok()

		return nil
	}

	return ErrFail
}

func isOK() bool {
	chTrue := make(chan interface{})
	close(chTrue)

	chFalse := make(chan interface{})
	close(chFalse)

	select {
	case <-chTrue:
		return true
	case <-chFalse:
		return false
	}
}

func getChecker() (cherker func(), ok func()) {
	var success bool

	cherker = func() {
		if success {
			fmt.Printf("success: %t", success)
		}
	}

	ok = func() {
		success = true
	}

	return
}
