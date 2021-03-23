package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	req, err := http.NewRequest("GET", "https://exmaple.com/", nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	if err := retry(client, req); err != nil {
		log.Fatal(err)
	}
}

func retry(client *http.Client, req *http.Request) error {
	for i := 0; i < 3; i++ {
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		fmt.Println(resp)
		time.Sleep(1 * time.Second)
	}
	return nil
}
