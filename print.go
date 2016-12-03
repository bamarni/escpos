package escpos

func (e *escpos) Print(s string) (string, error) {
	return e.enc.String(s)
}
