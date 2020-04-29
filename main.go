package main

import (
	"fmt"
	"os"

	"github.com/komsec/konstant/checks"
)

func main() {
	res, err := checks.RunAudit()
	fmt.Println(res)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
