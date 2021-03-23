package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func main() {
	if err := a(); err != nil {
		if err, ok := err.(stackTracer); ok {
			for _, f := range err.StackTrace() {
				fmt.Printf("%+s:%d\n", f, f)
			}
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}

func a() error {
	return b()
}

func b() error {
	return c()
}

func c() error {
	return d()
}

func d() error {
	return errors.WithStack(fmt.Errorf("%s", "hoge"))
}
