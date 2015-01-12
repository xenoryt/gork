package main

import (
	"fmt"
	"github.com/xenoryt/gork/rect"
	"github.com/xenoryt/gork/ui"
	"time"
)

//Create a ui/drawable.Drawable
type Player struct {
	x, y   int
	symbol byte
	dir    byte
}

func (player Player) Loc() (int, int) {
	return player.x, player.y
}
func (player Player) Symbol() byte {
	return player.symbol
}
func (player *Player) Move(maxwidth int) {
	if player.dir == 0 {
		player.dir = 1
	}
	player.x += int(player.dir)
	if player.x >= maxwidth {
		player.x = maxwidth - 1
		player.dir = -player.dir
	} else if player.x < 0 {
		player.x = 0
		player.dir = -player.dir
	}
}

func main() {
	display, err := ui.Init(false)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer display.Close()

	//display.Timeout(200)

	player := Player{
		y:      2,
		dir:    1,
		symbol: '@',
	}
	display.TrackDrawable(&player)
	inputChan := display.GetInputChan()
loop:
	for {
		w := display.Width()
		player.Move(w)
		//display.Error(string(byte(player.x) + '0'))
		display.SetView(rect.New(0, 0, 20, 20))
		switch ch := <-inputChan; ch {
		case 'q':
			break loop
		}
		time.Sleep(200)
	}

}
