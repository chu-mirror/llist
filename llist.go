// Package llist implements Lisp style list.
package llist

type any interface{}

// Unlike Lisp, LList does not accept dotted list
type cell struct {
	head any
	tail LList
}
type LList *cell

func Cons(head any, tail LList) LList {
	return &cell{head, tail};
}

func Car(l LList) any {
	return l.head
}

func Cdr(l LList) LList {
	return l.tail
}

// ToList converts a slice to LList.
func ToList(eles []any) LList {
	if len(eles) == 0 {
		return nil
	} else {
		return Cons(eles[0], ToList(eles[1:]))
	}
}

func List(eles ...any) LList {
	return ToList(eles);
}

