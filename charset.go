package escpos

import (
	"fmt"

	"golang.org/x/text/encoding/charmap"
)

// [Name]	Select character code table
// [Format]
// 		ASCII	   	ESC	  	t	  	n
// 		Hex			1B		74		n
// 		Decimal		27		116		n
// [Range]	different depending on the printers
// [Default]
// 		n = 20	   	[Thai models]
// 		n = 0	   	[Other models]
// [Description] 	Selects a page n from the character code table as follows:
// 		n	Character code table
// 		0	Page 0 [PC437: USA, Standard Europe]
// 		1	Page 1 [Katakana]
// 		2	Page 2 [PC850: Multilingual]
// 		3	Page 3 [PC860: Portuguese]
// 		4	Page 4 [PC863: Canadian-French]
// 		5	Page 5 [PC865: Nordic]
// 		6	Page 6 [Hiragana]
// 		7	Page 7 [One-pass printing Kanji characters]
// 		8	Page 8 [One-pass printing Kanji characters]
// 		11	Page 11 [PC851: Greek]
// 		12	Page 12 [PC853: Turkish]
// 		13	Page 13 [PC857: Turkish]
// 		14	Page 14 [PC737: Greek]
// 		15	Page 15 [ISO8859-7: Greek]
// 		16	Page 16 [WPC1252]
// 		17	Page 17 [PC866: Cyrillic #2]
// 		18	Page 18 [PC852: Latin 2]
// 		19	Page 19 [PC858: Euro]
// 		20	Page 20 [Thai Character Code 42]
// 		21	Page 21 [Thai Character Code 11]
// 		22	Page 22 [Thai Character Code 13]
// 		23	Page 23 [Thai Character Code 14]
// 		24	Page 24 [Thai Character Code 16]
// 		25	Page 25 [Thai Character Code 17]
// 		26	Page 26 [Thai Character Code 18]
// 		30	Page 30 [TCVN-3: Vietnamese]
// 		31	Page 31 [TCVN-3: Vietnamese]
// 		32	Page 32 [PC720: Arabic]
// 		33	Page 33 [WPC775: Baltic Rim]
// 		34	Page 34 [PC855: Cyrillic]
// 		35	Page 35 [PC861: Icelandic]
// 		36	Page 36 [PC862: Hebrew]
// 		37	Page 37 [PC864: Arabic]
// 		38	Page 38 [PC869: Greek]
// 		39	Page 39 [ISO8859-2: Latin 2]
// 		40	Page 40 [ISO8859-15: Latin 9]
// 		41	Page 41 [PC1098: Farsi]
// 		42	Page 42 [PC1118: Lithuanian]
// 		43	Page 43 [PC1119: Lithuanian]
// 		44	Page 44 [PC1125: Ukrainian]
// 		45	Page 45 [WPC1250: Latin 2]
// 		46	Page 46 [WPC1251: Cyrillic]
// 		47	Page 47 [WPC1253: Greek]
// 		48	Page 48 [WPC1254: Turkish]
// 		49	Page 49 [WPC1255: Hebrew]
// 		50	Page 50 [WPC1256: Arabic]
// 		51	Page 51 [WPC1257: Baltic Rim]
// 		52	Page 52 [WPC1258: Vietnamese]
// 		53	Page 53 [KZ-1048: Kazakhstan]
// 		66	Page 66 [Devanagari]
// 		67	Page 67 [Bengali]
// 		68	Page 68 [Tamil]
// 		69	Page 69 [Telugu]
// 		70	Page 70 [Assamese]
// 		71	Page 71 [Oriya]
// 		72	Page 72 [Kannada]
// 		73	Page 73 [Malayalam]
// 		74	Page 74 [Gujarati]
// 		75	Page 75 [Punjabi]
// 		82	Page 82 [Marathi]
// 		254	Page 254
// 		255	Page 255

type charset byte

const (
	CharsetPC437 charset = 0
	CharsetPC850 charset = 2
	CharsetPC860 charset = 3
	CharsetPC863 charset = 4
	CharsetPC865 charset = 5
	// whatever...
)

func (e *escpos) Charset(c charset) string {
	switch c {
	case CharsetPC437:
		e.enc = charmap.CodePage437.NewEncoder()
	case CharsetPC850:
		e.enc = charmap.CodePage850.NewEncoder()
	case CharsetPC860:
		e.enc = charmap.CodePage860.NewEncoder()
	case CharsetPC863:
		e.enc = charmap.CodePage863.NewEncoder()
	case CharsetPC865:
		e.enc = charmap.CodePage865.NewEncoder()
	}
	return fmt.Sprintf("\x1Bt%c", c)
}
