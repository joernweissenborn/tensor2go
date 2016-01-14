package tensor2go

import (
	"testing"
)

func TestVectorBasics(t *testing.T) {

	v1 := NewVector(3)

	if v1.Size() != 3 {
		t.Errorf("Vectorsize is %d, expected 3", v1.Size())
	}
}
