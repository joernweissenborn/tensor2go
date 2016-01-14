package tensor2go

type Zero struct {}

func (Zero) Add(in interface{}) interface{}{
	return in
}

func (Zero) Multiply(interface{}) interface{}{
	return nil
}
