package llist

import "testing"

func TestCons(t *testing.T) {
	l := List(1, 2)
	if l.head != 1 || l.tail.head != 2 {
		t.Error("Cons failed")
	}
}

func TestCar(t *testing.T) {
	l := List(1)
	if Car(l) != 1 {
		t.Error("Car failed")
	}
}

func TestCdr(t *testing.T) {
	l := List(1, 2)
	if Cdr(l).head != 2 {
		t.Error("Cdr failed")
	}
}

