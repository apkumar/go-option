package option

import (
	"testing"
)

func TestOption(t *testing.T) {
	noneOption := None[int]()

	if noneOption.IsSome() {
		t.Errorf("None option should not be Some.")
	}

	val, exists := noneOption.Get()
	if exists {
		t.Errorf("None option should not exist.")
	}

	val = noneOption.UnwrapOr(100)
	if val != 100 {
		t.Errorf("UnwrapOr did not return default value.")
	}

	noneOption.Set(10)
	if !noneOption.IsSome() {
		t.Errorf("After Set, option should be some.")
	}

	val, exists = noneOption.Get()
	if val != 10 {
		t.Errorf("After Set, val should be 10.")
	}
}
