package escpos

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
)

type Escpos struct {
	enc *encoding.Encoder
}

func NewEscpos() *Escpos {
	return &Escpos{
		enc: charmap.CodePage437.NewEncoder(),
	}
}
