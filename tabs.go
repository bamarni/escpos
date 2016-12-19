package escpos

import (
	"errors"
)

// [Name]	Horizontal tab
// [Format]
// 	ASCII	   	HT
// 	Hex		09
// 	Decimal		9
// [Description] 	Moves the print position to the next horizontal tab position.

func (e *Escpos) Tab() error {
	_, err := e.w.Write([]byte{9})
	return err
}

// [Name]	Set horizontal tab positions
// [Format]
// 	ASCII	   	ESC	  	D	  	n1	  	...	  	nk	  	NUL
// 	Hex		1B		44		n1		...		nk		00
// 	Decimal		27		68		n1		...		nk		0
// [Range]	n = 1 – 255
// 		k = 0 – 32
// [Default]	n = 8, 16, 24, 32, ... (Every eight characters for the default font set)
// [Description] 	Sets horizontal tab positions.

func (e *Escpos) TabPositions(positions ...byte) error {
	if len(positions) > 32 {
		return errors.New("tab positions cannot exceed 32")
	}
	return e.write(append(append([]byte{27, 68}, positions...), 0))
}
