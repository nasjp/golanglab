package main

import (
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

func run() error {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8, 9, 10}
	s3 := []int{11, 12, 13, 14, 15}

	s1 = append(append(s1, s2...), s3...)
	fmt.Println(s1)
	return nil
}
