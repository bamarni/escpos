package escpos

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"io"
)

type Escpos struct {
	enc *encoding.Encoder
	w   io.Writer
}

func NewEscpos(w io.Writer) *Escpos {
	return &Escpos{
		enc: charmap.CodePage437.NewEncoder(),
		w:   w,
	}
}

func (e *Escpos) write(b []byte) error {
	_, err := e.w.Write(b)
	return err
}
