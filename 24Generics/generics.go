// Starting with version 1.18, Go has added support for
// _generics_, also known as _type parameters_.

package main

import "fmt"

// As an example of a generic function, `SlicesIndex` takes
// a slice of any `comparable` type and an element of that
// type and returns the index of the first occurrence of
// v in s, or -1 if not present. The `comparable` constraint
// means that we can compare values of this type with the
// `==` and `!=` operators. For a more thorough explanation
// of this type signature, see [this blog post](https://go.dev/blog/deconstructing-type-parameters).
// Note that this function exists in the standard library
// as [slices.Index](https://pkg.go.dev/slices#Index).


// TODO: Create function SlicesIndex[S ~[]E, E comparable](s S, v E) int that returns the index of the first occurrence of v in s, or -1 if not present

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i, e := range s {
		if e == v {
			return i
		}
	}
	return -1
}

// As an example of a generic type, `List` is a
// singly-linked list with values of any type.

// TODO: Define struct List[T any] with head and tail pointers to element[T]

type List[T any] struct {
	head *element[T]
	tail *element[T]
}


// TODO: Define struct element[T any] with next pointer to element[T] and val field of type T

type element[T any] struct {
	next *element[T]
	val T
}
// We can define methods on generic types just like we
// do on regular types, but we have to keep the type
// parameters in place. The type is `List[T]`, not `List`.


// TODO: Define method Push(v T) on List[T] that pushes a value v to the list

func (l *List[T]) Push(v T) {
	if l.tail == nil {
		l.head = &element[T]{val: v}
		l.tail = l.head
	} else {
		l.tail.next = &element[T]{val: v}
		l.tail = l.tail.next
	}
}


// AllElements returns all the List elements as a slice.
// In the next example we'll see a more idiomatic way
// of iterating over all elements of custom types.


// TODO: Define method AllElements() []T on List[T] that returns all the List elements as a slice

func (l *List[T]) AllElements() []T {
	var elements []T
	for e := l.head; e != nil; e = e.next {
		elements = append(elements, e.val)
	}
	return elements
}


func main() {

	// TODO: Create var s = []string{"foo", "bar", "zoo"}
	s := []string{"foo", "bar", "zoo"}

	// When invoking generic functions, we can often rely
	// on _type inference_. Note that we don't have to
	// specify the types for `S` and `E` when
	// calling `SlicesIndex` - the compiler infers them
	// automatically.

	// TODO: Print index of zoo, zoo should be 2
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))
	// ... though we could also specify them explicitly.

	// TODO: Get index of zoo using explicit types
	fmt.Println("index of zoo:", SlicesIndex[[]string, string](s, "zoo"))
	// TODO: Create lst := List[int]{}
	lst := List[int]{}
	// TODO: Push 10, 13, 23 to lst
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	// TODO: Print list: lst.AllElements()
	fmt.Println("list:", lst.AllElements())
}