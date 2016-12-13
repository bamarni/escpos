package escpos

// [Name]	Set barcode height
// [Format]
//	 ASCII	   	GS	  	h	  	n
//	 Hex		1D		68		n
//	 Decimal	29		104		n
// [Range]	n = 1 – 255
// [Default]	n: different depending on the printers
// [Description] 	Sets the height of a barcode to n dots.

func (e *Escpos) BarcodeHeight(n byte) []byte {
	return []byte{29, 104, n}
}
