package escpos

// [Name]	Set left margin
// [Format]
// 	ASCII	   	GS	  	L	  	nL	  	nH
// 	Hex		1D		4C		nL		nH
// 	Decimal		29		76		nL		nH
// [Range]	(nL + nH × 256) = 0 – 65535
// [Default]	(nL + nH × 256) = 0
// [Description] 	In Standard mode, sets the left margin to (nL + nH × 256) × (horizontal motion unit) from the left edge of the printable area.

func (e *Escpos) MarginLeft(n uint16) []byte {
	return []byte{29, 76, n & 0xff, n >> 8}
}