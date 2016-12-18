package escpos

import (
	"reflect"
	"testing"
)

func TestTabPositions(t *testing.T) {
	e := NewEscpos()
	if _, err := e.TabPositions(make([]byte, 33)...); err == nil {
		t.Fatalf("expected error")
	}
	if b, err := e.TabPositions(8, 20); err != nil || !reflect.DeepEqual(b, []byte{27, 68, 8, 20, 0}) {
		t.Fatal("tab positions failed")
	}
}
