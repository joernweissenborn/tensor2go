package tensor2go

type tensorProperties struct {
	dims int
}

func (tp *tensorProperties) raiseDims() {
	tp.dims++
}
