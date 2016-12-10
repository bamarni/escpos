package escpos

// [Name]	    Select justification
// [Format]
//      ASCII	   	ESC	  	a	  	n
//      Hex	    	1B		61		n
//      Decimal		27		97		n
// [Range]		n = 0 – 2, 48 – 50
// [Default]		n = 0
// [Description] 	In Standard mode, aligns all the data in one line to the selected layout, using n as follows:
//      n   	Justification
//      0, 48	Left justification
//      1, 49	Centered
//      2, 50	Right justification

const (
	AlignLeft   byte = 0
	AlignCenter byte = 1
	AlignRight  byte = 2
)

func (e *Escpos) Align(a byte) []byte {
	return []byte{27, 97, a}
}
