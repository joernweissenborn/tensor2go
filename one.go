package tensor2go

type One struct {}

func (One) Multiply(in interface{}) interface{}{
	return in
}

