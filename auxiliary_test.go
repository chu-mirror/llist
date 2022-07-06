package llist

import "testing"

func TestMap(t *testing.T) {
	l := List(1, 2)
	l = Map(func(n any)any{ return n.(int)+1}, l)
	if Car(l) != 2 || Cadr(l) != 3 {
		t.Error("Map failed")
	}
}

func TestFilter(t *testing.T) {
	l := List(1, 2, 3)
	l = Filter(func(n any)bool{ return n.(int) > 1 }, l)
	if Car(l) != 2 || Cadr(l) != 3 {
		t.Error("Filter failed")
	}

	l = List(1, 1.1, 2)
	l = Filter(func(n any)bool{ _, ok := n.(int); return ok }, l)
	if Car(l) != 1 || Cadr(l) != 2 {
		t.Error("Filter failed")
	}
}

func TestFold(t *testing.T) {
	l := List(1, 2, 3)
	if Fold(func(a, b any)any{ return a.(int)+b.(int) }, 0, l) != 6 {
		t.Error("Fold failed")
	}
}

func TestLength(t *testing.T) {
	l := List(1, 2, 3)
	if Length(l) != 3 {
		t.Error("Length failed")
	}
}

func TestCaar(t *testing.T) {
	l := List(List(1, 2), 3)
	if Caar(l) != 1 {
		t.Error("Caar failed")
	}
}

func TestCadr(t *testing.T) {
	l := List(1, 2, 3)
	if Cadr(l) != 2 {
		t.Error("Cadr failed")
	}
}

func TestCdar(t *testing.T) {
	l := List(List(1, 2), 3)
	if Cdar(l).head != 2 {
		t.Error("Cdar failed")
	}
}

func TestCddr(t *testing.T) {
	l := List(1, 2, 3)
	if Cddr(l).head != 3 {
		t.Error("Cddr failed")
	}
}

