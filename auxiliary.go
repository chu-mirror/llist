package llist

// Map create a new list by mapping a function on elements.
func Map(f func(any)any, l LList) LList {
	if l == nil {
		return nil
	} else {
		return Cons(f(Car(l)), Map(f, Cdr(l)))
	}
}

// Filter selects elements by a prediction, and create a new list using these
// elements.
func Filter(p func(any)bool, l LList) LList {
	if l == nil {
		return nil
	} else {
		if p(Car(l)) {
			return Cons(Car(l), Filter(p, Cdr(l)))
		} else {
			return Filter(p, Cdr(l))
		}
	}
}

// Foreach execute a function on each element.
func Foreach(e func(any), l LList) {
	if l == nil {
		return
	} else {
		e(Car(l))
		Foreach(e, Cdr(l))

	}
}

// Fold sums up all elements to a single value.
func Fold(s func(any, any)any, init any, l LList) any {
	if l == nil {
		return init
	} else {
		return Fold(s, s(Car(l), init), Cdr(l))
	}
}

// Length returns the length of a list.
func Length(l LList) int {
	if l == nil {
		return 0
	} else {
		return Length(Cdr(l)) + 1
	}
}

func Caar(l LList) any {
	h := Car(l).(LList)
	return Car(h)
}

func Cadr(l LList) any {
	t := Cdr(l)
	return Car(t)
}

func Cdar(l LList) LList {
	h := Car(l).(LList)
	return Cdr(h)
}

func Cddr(l LList) LList {
	t := Cdr(l)
	return Cdr(t)
}
