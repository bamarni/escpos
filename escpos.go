package escpos

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
)

type escpos struct {
	enc *encoding.Encoder
}

func NewEscpos() *escpos {
	return &escpos{
		enc: charmap.CodePage437.NewEncoder(),
	}
}
