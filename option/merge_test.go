package option

import (
	"testing"
)

func TestMerge(t *testing.T) {
	type Options struct {
		String string
		Nullable *string
		AnonymousStruct struct{foo string}
		unexported string
	}
	bar := "bar"

	opt := Options {
		"foo",
		nil,
		struct{foo string}{foo: "hello"},
		"baz",
	}
	
	optTwo := Options {
		"",
		&bar,
		struct{foo string}{},
		"quux",
	}

	out := Merge(opt, optTwo)

	if out.String != "foo" {
		t.Errorf("String was %s", out.String)
	}

	if *out.Nullable !=  "bar" {
		t.Errorf("Nullable was %s", *out.Nullable)
	}

	if out.AnonymousStruct.foo != "hello" {
		t.Errorf("AnonStruct was %s", out.AnonymousStruct.foo)
	}

	if out.unexported != "baz" {
		t.Errorf("unexported was %s", out.unexported)
	}
}
