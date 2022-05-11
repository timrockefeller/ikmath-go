package basic

type F64 struct {
	raw float64
}

func (n *F64) Add(args ...F64) *F64 {
	for _, v := range args {
		n.raw += v.raw
	}
	return n
}
