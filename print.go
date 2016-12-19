package escpos

func (e *Escpos) Print(s string) error {
	b, err := e.enc.Bytes([]byte(s))
	if err != nil {
		return err
	}
	return e.write(b)
}

// [Name]	Select print mode(s)
// [Format]
// 	ASCII	   	ESC	  	!	  	n
// 	Hex		1B	 	21	 	n
// 	Decimal		27	 	33	 	n
// [Range]	n = 0 – 255
// [Default]	different depending on the printers
// [Description] 	Selects the character font and styles (emphasized, double-height, double-width, and underline) together as follows:
// 	n: Bit	Binary	Function		Hex	Decimal
// 	0	0	Selects Font 1		00	0
// 		1	Selects Font 2		01	1
// 	1	−	(Reserved)		−	−
// 	2	−	(Reserved)		−	−
// 	3	0	Emphasized mode: OFF	00	0
// 		1	Emphasized mode: ON	08	8
// 	4	0	Double-height mode: OFF	00	0
// 		1	Double-height mode: ON	10	16
// 	5	0	Double-width mode: OFF	00	0
// 		1	Double-width mode: ON	20	32
// 	6	−	(Reserved)		−	−
// 	7	0	Underline mode: OFF	00	0
// 		1	Underline mode: ON	80	128

const (
	PrintModeFontB        byte = 0x01
	PrintModeEmphasized   byte = 0x08
	PrintModeDoubleHeight byte = 0x10
	PrintModeDoubleWidth  byte = 0x20
	PrintModeUnderline    byte = 0x80
)

func (e *Escpos) PrintMode(n byte) error {
	return e.write([]byte{27, 33, n})
}
