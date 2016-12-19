package escpos

import (
	"bytes"
	"reflect"
	"testing"
)

func TestTabPositions(t *testing.T) {
	var b bytes.Buffer
	e := NewEscpos(&b)
	if err := e.TabPositions(make([]byte, 33)...); err == nil {
		t.Fatalf("expected error")
	}
	if err := e.TabPositions(8, 20); err != nil || !reflect.DeepEqual(b.Bytes(), []byte{27, 68, 8, 20, 0}) {
		t.Fatal("tab positions failed")
	}
}
