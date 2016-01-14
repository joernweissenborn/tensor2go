package tensor2go

type Vector struct {
	*Tensor
}

func NewVector(size int) (v *Vector){
	v = &Vector{NewTensor(size)}
	return
}

func (v *Vector) Size() (s int) {
	s = v.DimSize(0)
	return
}
