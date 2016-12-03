package escpos

import (
	"fmt"
)

// [Name]	    Select justification
// [Format]
//      ASCII	   	ESC	  	a	  	n
//      Hex	    	1B		61		n
//      Decimal		27		97		n
// [Range]	    n = 0 – 2, 48 – 50
// [Default]	n = 0
// [Description] 	In Standard mode, aligns all the data in one line to the selected layout, using n as follows:
//      n   	Justification
//      0, 48	Left justification
//      1, 49	Centered
//      2, 50	Right justification

type align byte

const (
	AlignLeft   align = 0
	AlignCenter align = 1
	AlignRight  align = 2
)

func (e *escpos) Align(a align) string {
	return fmt.Sprintf("\x1Ba%c", a)
}
