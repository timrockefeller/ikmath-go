package composite

import (
	"ikmath/basic"
)

type Vector3 struct {
	X basic.I64
	Y basic.I64
	Z basic.I64
}

func NewVector3(x, y, z basic.I64) *Vector3 {
	return &Vector3{X: x, Y: y, Z: z}
}

func NewVector3FromRaw(x_raw, y_raw, z_raw int64) *Vector3 {
	return &Vector3{
		X: *basic.NewI64FromRaw(x_raw),
		Y: *basic.NewI64FromRaw(y_raw),
		Z: *basic.NewI64FromRaw(z_raw),
	}
}

func (v *Vector3) Copy() *Vector3 {
	return &Vector3{X: v.X, Y: v.Y, Z: v.Z}
}

func (v *Vector3) Add(a Vector3) *Vector3 {
	v.X.Add(a.X)
	v.Y.Add(a.Y)
	v.Z.Add(a.Z)
	return v
}

func (v *Vector3) Maginitude2() *basic.I64 {
	return basic.Add(*v.X.Sqr(), *v.Y.Sqr(), *v.Z.Sqr())
}
