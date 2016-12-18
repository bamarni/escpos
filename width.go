package escpos

// [Name]	Set print area width
// [Format]
// 	ASCII	   	GS	  	W	  	nL	  	nH
// 	Hex		1D		57		nL		nH
// 	Decimal		29		87		nL		nH
// [Range]	(nL + nH × 256) = 0 – 65535
// [Default]	Entire printable area
// [Description] 	In Standard mode, sets the print area width to (nL + nH × 256) × (horizontal motion unit).

func (e *Escpos) Width(n uint16) []byte {
	return []byte{29, 87, byte(n & 0xff), byte(n >> 8)}
}
