package escpos

// [Name]	Print barcode
// [Format]
// 	(A)	   	ASCII	   	GS	  	k	  	m	  	d1 ... dk	  	NUL
// 			Hex		1D		6B		m		d1 ... dk		NUL
// 			Decimal		29		107		m		d1 ... dk		NUL
// 	(B)	   	ASCII	   	GS	  	k	  	m	  	n	 	 	d1 ... dn
// 			Hex		1D		6B		m		n			d1 ... dn
// 			Decimal		29		107		m		n			d1 ... dn
// [Range]	m: different depending on the printers d, k of (A), and d, n of (B): different depending on the barcode format. Refer to the tables in the following [Description].
// [Default]	None
// [Description] 	Prints the barcode using the barcode system specified by m.
// 	<Function A>
// 	m	Barcode system				Barcode data ("SP" in the table indicates space.)
// 					Number of bytes		The range of k		Characters (ASCII)	Data ( d )
//
// 	0	UPC-A			Fixed			k = 11, 12		0 – 9			d = 48 – 57
//
// 	1	UPC-E			Fixed			k = 6 – 8, 11, 12	0 – 9			d = 48 – 57
// 														(However, d1 = 48 when k = 7, 8, 11, 12)
//
// 	2	JAN13 (EAN13)		Fixed			k = 12, 13		0 – 9			d = 48 – 57
//
// 	3	JAN8 (EAN8)		Fixed			k = 7, 8		0 – 9			d = 48 – 57
//
// 	4	CODE39			Can be changed		1 ≤ k			0 – 9, A – Z, SP, $,    d = 48 – 57, 65 – 90, 32, 36, 37,
// 											%, *, +, -, ., /     	42, 43, 45, 46, 47
//
// 	5	ITF			Can be changed		2 ≤ k			0 – 9			d = 48 – 57
// 		(Interleaved 2 of 5)				(even number)
//
// 	6	CODABAR (NW-7)		Can be changed		2 ≤ k			0 – 9, A – D, a – d,	d = 48 – 57, 65 – 68, 97 – 100,
// 											$, +, −, ., /, :	36, 43, 45, 46, 47, 58
// 														(However, d1 = 65 – 68, dk = 65 – 68, d1 = 97 – 100, dk = 97 – 100)
//
// k indicates the number of bytes of barcode data. k is an explanation parameter; therefore it does not need to be transmitted.
// d specifies the character code data of the barcode data to be printed.

const (
	BarcodeUPCA    byte = 0
	BarcodeUPCE    byte = 1
	BarcodeJAN13   byte = 2
	BarcodeJAN8    byte = 3
	BarcodeCODE39  byte = 4
	BarcodeITF     byte = 5
	BarcodeCODABAR byte = 6
)

func (e *Escpos) Barcode(m byte, data string) ([]byte, error) {
	// TODO: validate data according to the barcode system
	return append([]byte{29, 107, m}, data...), nil
}
