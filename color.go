package escpos

// [Name]	Select print color
// [Format]
// 	ASCII	   	ESC	  	r	  	n
// 	Hex		1B		72		n
// 	Decimal		27		114		n
// [Range]	n = 0, 1, 48, 49
// [Default]	n = 0
// [Description] 	Selects a print color, using n as follows:
// 	n	Print color
// 	0, 48	Black
// 	1, 49	Red

const (
	ColorBlack byte = 0
	ColorRed   byte = 1
)

func (e *Escpos) Color(n byte) error {
	return e.write([]byte{27, 114, n})
}
