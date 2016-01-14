package tensor2go

type Scalar struct {
	*Tensor
}

func NewScalar(val interface{}) (v *Scalar){
	v = &Scalar{NewTensor()}
	v.Tensor.tensorRoot = v.Tensor.graph.CreateVertex("root",val)
	return
}

func (s *Scalar) Val() interface{} {
	return s.Tensor.tensorRoot.Properties()
}

func (s *Scalar) Multiply(val interface{})(r interface{}){
	m, ok := s.Val().(Multiplier)
	if !ok {
		panic("Value is not a multiplier")
	}
	if v, ok := val.(*Scalar); ok {
		r = m.Multiply(v.Val())
	}

	return
}
