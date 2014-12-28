package TextDisplay

import (
	. "github.com/xenoryt/gork/errors"
	//. "../../errors"
)

type pane struct {
	buffer   [][]byte
	width    int
	height   int
	caretPos int
}

func newPane(buffer []byte, bwidth, x, y, w, h int) pane {
	panelBuffer := make([][]byte, h)
	for i := 0; i < h; i++ {
		start := (y+i)*bwidth + x
		end := (y+i)*bwidth + x + w
		panelBuffer[i] = buffer[start:end]
	}
	return pane{panelBuffer, w, h, 0}
}

func (p pane) inBounds(x, y int) bool {
	if x < 0 || x >= p.width {
		return false
	}
	if y < 0 || y >= p.height {
		return false
	}
	return true
}

func (p *pane) put(index int, char byte) {
	p.buffer[index/p.width][index%p.width] = char
}

//Prints text to the pane
func (p *pane) print(text []byte) error {
	//We want to start from the previous print's position
	bpos := p.caretPos

	//We want to save the new position of the caret after printing
	var i int
	defer func() {
		p.caretPos = bpos + i
	}()

	for i = range text {
		caret := i + bpos
		//Make sure we're not out of bounds
		if caret >= p.width*p.height {
			return GenericError("Error: couldn't print everything. Pane overflowed!")
		}

		//Add the character to the buffer
		switch text[i] {
		case '\n':
			bpos += p.width % i
		case '\t':
			bpos += 4 - (i % 4)
			//if (i+bpos) % p.width < 4 {
			//	//We went onto another line.
			//}
		default:
			p.put(caret, text[i])
		}
	}
	return nil
}

//Clears the pane of any text.
func (p *pane) clear() {
	//We shall start from the current caret position and keep
	//rolling back until we hit the start again.
	i := &p.caretPos
	for ; *i >= 0; *i-- {
		p.put(*i, ' ')
	}
}
