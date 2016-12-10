package escpos

func (e *Escpos) Print(s string) ([]byte, error) {
	return e.enc.Bytes([]byte(s))
}
