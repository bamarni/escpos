package escpos

import "testing"

func TestCharset(t *testing.T) {
	e := NewEscpos()

	var err error
	// PC437 doesn't support Ê
	if _, err := e.Print("Être ou ne pas être"); err == nil {
		t.Fatalf("expected error")
	}
	// but PC850 does
	e.Charset(CharsetPC850)
	if _, err = e.Print("Être ou ne pas être"); err != nil {
		t.Fatalf("unexpected error : %s", err)
	}
}
