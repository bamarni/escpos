package escpos

// [Name]	Select character font
// [Format]
// 	ASCII	   	ESC	  	M	  	n
// 	Hex		1B		4D		n
// 	Decimal		27		77		n
// [Range]	different depending on the printers
// [Default]	different depending on the printers
// [Description] 	Selects a character font, using n as follows:
// 	n	Font
// 	0, 48	Font A
// 	1, 49	Font B
// 	2, 50	Font C
// 	3, 51	Font D
// 	4, 52	Font E
// 	97	Special font A
// 	98	Special font B

const (
	FontA        byte = 0
	FontB        byte = 1
	FontC        byte = 2
	FontD        byte = 3
	FontE        byte = 4
	FontSpecialA byte = 97
	FontSpecialB byte = 98
)

func (e *Escpos) Font(n byte) []byte {
	return []byte{27, 77, n}
}
