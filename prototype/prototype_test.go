package prototype

import "testing"

var prototypeManager *PrototypeManager

type TypeAColoneable struct {
	data string
}

func (t *TypeAColoneable) Clone() Cloneable {
	tc := *t
	return &tc
}

type TypeBColoneable struct {
	data string
}

func (t *TypeBColoneable) Clone() Cloneable {
	tc := *t
	return &tc
}

func init() {
	prototypeManager = NewPrototypeManager()
	ta := TypeAColoneable{data: "typeA"}
	prototypeManager.Set("typeA", &ta)

}

func TestClone(t *testing.T) {
	ta := prototypeManager.Get("typeA")
	taClone := ta.Clone()
	if ta == taClone {
		t.Fatal("error! clone return same obj")
	}

	taClone2 := prototypeManager.Get("typeA").Clone()

	taClone2Type := taClone2.(*TypeAColoneable)
	if taClone2Type.data != "typeA" {
		t.Fatal("error! clone return wrong type")
	}
}
