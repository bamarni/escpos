package escpos

// [Name]	Set horizontal and vertical motion units
// [Format]
// 	ASCII	   	GS	  	P	  	x	  	y
// 	Hex		1D		50		x		y
// 	Decimal		29		80		x		y
// [Range]	x = 0 – 255
// 		y = 0 – 255
// [Default]	x, y: different depending on the printers
// [Description] 	Sets the horizontal and vertical motion units to approximately 25.4/x mm {1/x"} and approximately 25.4/y mm {1/y"}, respectively.
// 			When x = 0, the default value of the horizontal motion unit is used.
// 			When y = 0, the default value of the vertical motion unit is used.

func (e *Escpos) Units(x, y byte) error {
	return e.write([]byte{29, 80, x, y})
}
