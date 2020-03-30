package main

import (
	"bytes"
	"encoding/json"
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
	var i interface{}
	i = &User{Name: "hoge"}
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(i); err != nil {
		return err
	}
	fmt.Println(buf.String())
	return nil
}

type User struct {
	Name string `json:"name"`
}
