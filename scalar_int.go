package tensor2go

import "fmt"

type ScalarInt int

func (i ScalarInt) Multiply(val interface{}) (r interface{}) {
	switch val.(type) {
	case int:
		r = i * ScalarInt(val.(int))
	case int8:
		r = i * ScalarInt(val.(int8))
	case int16:
		r = i * ScalarInt(val.(int16))
	case int32:
		r = i * ScalarInt(val.(int32))
	case int64:
		r = i * ScalarInt(val.(int64))
	case ScalarInt:
		r = i * val.(ScalarInt)
	default:
		panic(fmt.Sprintf("ScalarInt multipling unsupported type %T",val))
	}
	return
}

func (i ScalarInt) Add(val interface{}) (r interface{}) {
	switch val.(type) {
	case int:
		r = i + ScalarInt(val.(int))
	case int8:
		r = i + ScalarInt(val.(int8))
	case int16:
		r = i + ScalarInt(val.(int16))
	case int32:
		r = i + ScalarInt(val.(int32))
	case int64:
		r = i + ScalarInt(val.(int64))
	case ScalarInt:
		r = i + val.(ScalarInt)
	default:
		panic(fmt.Sprintf("ScalarInt adding unsupported type %T",val))
	}
	return
}
