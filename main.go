package main

import (
	"fmt"
	"github.com/xenoryt/gork/rect"
	"github.com/xenoryt/gork/ui"
)

func main() {
	display, err := ui.Init(false)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer display.Close()

	display.Timeout(200)

loop:
	for {
		display.Update(rect.New(0, 0, 20, 20))
		switch ch := display.GetInput(); ch {
		case 'q':
			break loop
		case 0:
			continue
		default:
			display.Print("char: " + string(ch))
		}
		//ui.Update
	}

}
