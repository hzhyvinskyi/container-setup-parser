package main

import (
	"flag"
	"fmt"
)

func main() {
	setup := flag.String("setup", "[p(*)({[{}]})]()", "Container setup string")
	flag.Parse()

	fmt.Println(validate(*setup, len(*setup)))
}
