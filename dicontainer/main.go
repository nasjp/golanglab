package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nasjp/golanglab/dicontainer/services"
	"go.uber.org/dig"
)

/*
$ curl localhost:8080/upper --header 'Content-Type: application/json' -i --data-raw '{"name": "hoge"}'
*/

func main() {
	if err := diContainer(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func normal() error {
	p := services.NewUpperPresenter()
	u := services.NewUpperUsecase(p)
	h := services.NewUpperHandler(u)
	http.HandleFunc("/upper", h.Handle)
	http.ListenAndServe(":8080", nil)
	return nil
}

func diContainer() error {
	c := dig.New()
	if err := c.Provide(services.NewUpperPresenter); err != nil {
		return err
	}
	if err := c.Provide(services.NewUpperUsecase); err != nil {
		return err
	}
	if err := c.Provide(services.NewUpperHandler); err != nil {
		return err
	}
	err := c.Invoke(func(h *services.UpperHandler) {
		http.HandleFunc("/upper", h.Handle)
		http.ListenAndServe(":8080", nil)
	})
	if err != nil {
		return err
	}
	return nil
}
