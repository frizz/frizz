package frizz

import (
	"testing"
)

func TestStack_String(t *testing.T) {
	s := Stack{
		RootItem("root"),
		FieldItem("field"),
		MapItem("map"),
		ArrayItem(1),
	}
	expected := `root.field["map"][1]`
	if s.String() != expected {
		t.Fatalf("Stack string not what expected: %s", s.String())
	}
}

func TestStack_Append(t *testing.T) {
	s := Stack{}
	s1 := s.Append(RootItem("root"))
	s2 := s1.Append(FieldItem("field"))
	s3 := s2.Append(MapItem("map"))
	s4 := s3.Append(ArrayItem(1))
	if s.String() != `` {
		t.Fatalf("Stack string not what expected: %s", s.String())
	}
	if s1.String() != `root` {
		t.Fatalf("Stack string not what expected: %s", s.String())
	}
	if s2.String() != `root.field` {
		t.Fatalf("Stack string not what expected: %s", s.String())
	}
	if s3.String() != `root.field["map"]` {
		t.Fatalf("Stack string not what expected: %s", s.String())
	}
	if s4.String() != `root.field["map"][1]` {
		t.Fatalf("Stack string not what expected: %s", s.String())
	}
}
