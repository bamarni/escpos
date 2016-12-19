package escpos

import (
	"io/ioutil"
	"testing"
)

func TestCharset(t *testing.T) {
	e := NewEscpos(ioutil.Discard)

	// PC437 doesn't support Ê
	if err := e.Print("Être ou ne pas être"); err == nil {
		t.Fatalf("expected error")
	}
	// but PC850 does
	e.Charset(CharsetPC850)
	if err := e.Print("Être ou ne pas être"); err != nil {
		t.Fatalf("unexpected error : %s", err)
	}
}
