package main

import (
	"fmt"

	"github.com/xenoryt/gork/shape"
	"github.com/xenoryt/gork/ui"
)

func main() {
	display, err := ui.Init(ui.TextDisplay)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer display.Close()

loop:
	for {
		display.Update(shape.NewRect(0, 0, 20, 20))
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
