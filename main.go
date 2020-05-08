package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/komsec/konstant/checks"
)

func main() {
	c := flag.String("checks-dir", "/etc/konstant", "Directory that contain konstant checks yaml files")
	flag.Parse()
	res, ok, err := checks.RunAudit(*c)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if !ok {
		os.Exit(1)
	}
}
