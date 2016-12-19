package escpos

const (
	FontA        byte = 0
	FontB        byte = 1
	FontC        byte = 2
	FontD        byte = 3
	FontE        byte = 4
	FontSpecialA byte = 97
	FontSpecialB byte = 98
)

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

func (e *Escpos) Font(n byte) error {
	return e.write([]byte{27, 77, n})
}

// [Name]	Select font for HRI characters
// [Format]
// 	ASCII	   	GS	  	f	  	n
// 	Hex		1D		66		n
// 	Decimal		29		102		n
// [Range]	n: different depending on the printers
// [Default]	n = 0
// [Description] 	Selects a font for the Human Readable Interpretation (HRI) characters when printing a barcode, using n as follows:
// 	n	Font of HRI characters
// 	0, 48	Font A
// 	1, 49	Font B
// 	2, 50	Font C
// 	3, 51	Font D
// 	4, 52	Font E
// 	97	Special font A
// 	98	Special font B

func (e *Escpos) BarcodeFont(n byte) error {
	return e.write([]byte{29, 102, n})
}
