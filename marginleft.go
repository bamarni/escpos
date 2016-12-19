package escpos

// [Name]	Set left margin
// [Format]
// 	ASCII	   	GS	  	L	  	nL	  	nH
// 	Hex		1D		4C		nL		nH
// 	Decimal		29		76		nL		nH
// [Range]	(nL + nH × 256) = 0 – 65535
// [Default]	(nL + nH × 256) = 0
// [Description] 	In Standard mode, sets the left margin to (nL + nH × 256) × (horizontal motion unit) from the left edge of the printable area.

func (e *Escpos) MarginLeft(n uint16) error {
	return e.write([]byte{29, 76, byte(n & 0xff), byte(n >> 8)})
}
