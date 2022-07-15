package internal

import (
	"testing"
)

func TestNibblerIteratesNibbles(t *testing.T) {
	bytes := []uint8{0x2a}
	nibbles := NewNibbler(bytes)
	if next := nibbles.Next(); next == nil || *next != 0x02 {
		t.Fail()
	}
	if next := nibbles.Next(); next == nil || *next != 0x0a {
		t.Fail()
	}
	if next := nibbles.Next(); next != nil {
		t.Fail()
	}
}
