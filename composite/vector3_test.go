package composite

import (
	. "ikmath/basic"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Log(`TestAdd.Begin`)
	a := NewVector3(*NewI64FromRaw(3), *NewI64FromRaw(1), *NewI64FromRaw(2))
	if a.Maginitude2().GetRaw() != int64(14) {
		t.Error(`Vector3::Maginitude2 failed`)
	}
}
