package basic

type I64 struct {
	raw int64
}

func NewI64FromRaw(_raw int64) *I64 {
	return &I64{raw: _raw}
}

func (n *I64) GetRaw() int64 {
	return n.raw
}

func (n *I64) Add(args ...I64) *I64 {
	for _, v := range args {
		n.raw += v.raw
	}
	return n
}

func Add(args ...I64) *I64 {
	n := NewI64FromRaw(0)
	for _, v := range args {
		n.raw += v.raw
	}
	return n
}

func (n *I64) Sub(args ...I64) *I64 {
	for _, v := range args {
		n.raw -= v.raw
	}
	return n
}

func (n *I64) Mul(args ...I64) *I64 {
	for _, v := range args {
		n.raw *= v.raw
	}
	return n
}

func Mul(args ...I64) *I64 {
	n := NewI64FromRaw(1)
	for _, v := range args {
		n.raw *= v.raw
	}
	return n
}

func (n *I64) DivBy(arg I64) *I64 {
	n.raw /= arg.raw
	return n
}

func Pow(n *I64, times int) *I64 {
	var result int64 = 1
	for i := 0; i < times; i++ {
		result *= n.raw
	}
	return &I64{raw: result}
}

func (n *I64) Sqr() *I64 {
	return Pow(n, 2)
}
