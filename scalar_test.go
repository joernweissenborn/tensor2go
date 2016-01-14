package tensor2go

import "testing"

func TestScalarMultiply(t *testing.T) {
	va := ScalarInt(2)
	vb := ScalarInt(21)
	a := NewScalar(va)
	b := NewScalar(vb)
	
	if a.Val().(ScalarInt) != va {
		t.Error("a has wrong value:", a.Val())
	}
	
	if a.Multiply(b).(ScalarInt) != ScalarInt(42) {
		t.Error("The answer is wrong",a.Multiply(b))
	}
}
