package escpos

// [Name]	Print and line feed
// [Format]
// 	ASCII	   	LF
// 	Hex		0A
// 	Decimal		10
// [Description] 	Prints the data in the print buffer and feeds one line, based on the current line spacing.

// [Name]		Print and feed n lines
// [Format]
// 	ASCII	   	ESC	  	d	  	n
// 	Hex		1B		64		n
// 	Decimal		27		100		n
// [Range]		n = 0 – 255
// [Default]		None
// [Description] 	Prints the data in the print buffer and feeds n lines.

func (e *Escpos) Lf(n byte) []byte {
	if n > 1 {
		return []byte{27, 100, n}
	}
	return []byte{10}
}
