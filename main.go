package main

import (
	"fmt"
	"github.com/xenoryt/gork/ui"
)

func main() {
	if err := ui.Init(); err != nil {
		fmt.Println(err)
		return
	}
	defer ui.Close()
}
