package escpos

// [Name]	Set barcode width
// [Format]
// 	ASCII	   	GS	  	w	  	n
// 	Hex		1D		77		n
// 	Decimal		29		119		n
// [Range]	n: different depending on the printers
// [Default]	n: different depending on the printers
// [Description]	Sets the horizontal size of a barcode.
// 			n specifies the barcode module width.

func (e *Escpos) BarcodeWidth(n byte) []byte {
	return []byte{29, 119, n}
}
