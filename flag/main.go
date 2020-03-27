package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

func main() {
	pFlag()
}

/*
go run flag/main.go migrate up -all
だとできない

go run flag/main.go -all migrate up
じゃないとだめ(えー)
*/
func standardFlag() {
	all := flag.Bool("all", false, "all flag")

	fmt.Println(os.Args[1])
	fmt.Println(*all)

	flag.Parse()
	fmt.Println(*all)
	fmt.Println(flag.Args())
}

/*
pflag だとできる!
go run flag/main.go migrate up -all
*/
func pFlag() {
	all := pflag.BoolP("all", "a", false, "all flag")

	fmt.Println(os.Args[1])
	fmt.Println(*all)

	pflag.Parse()
	fmt.Println(*all)
	fmt.Println(flag.Args())
}
