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
	a := map[string]interface{}{
		"user_id":   1,
		"user_name": "hoge",
	}

	b := map[string]interface{}{
		"shop_id":   2,
		"shop_name": "huga",
	}

	c := make([]map[string]interface{}, 0)

	c = append(c, a, b)

	d := []map[string]interface{}{
		map[string]interface{}{
			"dog_id":   3,
			"dog_name": "piyo",
		},
	}

	fmt.Println(append(d, c...))
	return nil
}
