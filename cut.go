package escpos

// [Name]	Select cut mode and cut paper
// [Format]
// 	<Function A>	ASCII	   	GS	  	V	  	m
// 			Hex		1D		56		m
// 			Decimal		29		86		m
// 	<Function B>	ASCII	   	GS	  	V	  	m	  	n
// 			Hex		1D		56		m		n
// 			Decimal		29		86		m		n
// 	<Function C>	ASCII	   	GS		V	  	m	  	n
// 			Hex		1D		56		m		n
// 			Decimal		29		86		m		n
// 	<Function D>	ASCII	   	GS	  	V	  	m	  	n
// 			Hex		1D		56		m		n
// 			Decimal		29		86		m		n
// [Range]	m, n: different depending on the printers
// [Default]	None
// [Description] 	Executes paper cutting specified by m, as follows:
// 	Function name	Function									m	Cutting shape
// 	<Function A>	Cuts the paper									0, 48	Full cut
// 													1, 49	Partial cut
// 	<Function B>	Feeds paper and cuts the paper							65	Full cut
// 													66	Partial cut
// 	<Function C>	Sets the paper cutting position							97	Full cut
// 													98	Partial cut
// 	<Function D>	Feeds paper and cuts the paper, and feeds paper to print starting position	103	Full cut
// 													104	Partial cut
// The printer operation of the each function is as follows:
// 	Function name	Function
// 	<Function A>	Executes paper cut
// 	<Function B>	Feeds paper to [cutting position + (n × vertical motion unit)] and executes paper cut
// 	<Function C>	Preset [cutting position + (n × vertical motion unit)] to the paper cutting position, and executes paper cut when it reaches the autocutter position after printing and feeding
// 	<Function D>	Feeds paper to [cutting position + (n × vertical motion unit)] and executes paper cut, then moves paper to the print start position by reverse feeding

func (e *Escpos) Cut(partial bool) error {
	if partial {
		return e.write([]byte{29, 86, 1})
	}
	return e.write([]byte{29, 86, 0})
}

func (e *Escpos) CutB(partial bool, n byte) error {
	if partial {
		return e.write([]byte{29, 86, 66, n})
	}
	return e.write([]byte{29, 86, 65, n})
}

func (e *Escpos) CutC(partial bool, n byte) error {
	if partial {
		return e.write([]byte{29, 86, 98, n})
	}
	return e.write([]byte{29, 86, 97, n})
}

func (e *Escpos) CutD(partial bool, n byte) error {
	if partial {
		return e.write([]byte{29, 86, 104, n})
	}
	return e.write([]byte{29, 86, 103, n})
}
