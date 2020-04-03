package main

import (
	"fmt"
	"net/http"
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
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) { fmt.Println("hello world") })
	http.ListenAndServe(":8080", nil)
	return nil
}

type HelloHandler struct {
	HelloUsecase HelloUsecaseInterface
}

func (h HelloHandler) Handle(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, r.ContentLength)
	if _, err := r.Body.Read(b); err != nil {
		fmt.Fprintln(w, err)
		return
	}
	h.HelloUsecase.Do()
}

type HelloUsecaseInterface interface {
	Do()
}

type HelloUsecase struct{}

func (u *HelloUsecase) Do() {
	fmt.Println("hello")
}
