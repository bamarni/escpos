package escpos

// [Name]	Select print position of HRI characters
// [Format]
// 	ASCII	   	GS	  	H	  	n
// 	Hex		1D		48		n
// 	Decimal		29		72		n
// [Range]	n = 0 – 3
// 		n = 48 – 51
// [Default]	n = 0
// [Description] 	Selects the print position of Human Readable Interpretation (HRI) characters when printing a barcode, using n as follows:
// 	n	Print position
// 	0, 48	Not printed
// 	1, 49	Above the barcode
// 	2, 50	Below the barcode
// 	3, 51	Both above and below the barcode

const (
	BarcodeHRINone   byte = 0
	BarcodeHRITop    byte = 1
	BarcodeHRIBottom byte = 2
	BarcodeHRIBoth   byte = 3
)

func (e *Escpos) BarcodeHri(n byte) []byte {
	return []byte{29, 72, n}
}
