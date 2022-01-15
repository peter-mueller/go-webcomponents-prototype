package html

import (
	"syscall/js"
)

type Node interface {
	NodeValue() js.Value
}

type Element struct {
	value js.Value
}

func CreateElement(tagName string) Element {
	el := js.Global().Get("document").Call("createElement", tagName)
	return Element{value: el}
}

func NewElement(element js.Value) Element {
	return Element{value: element}
}
func (el Element) NodeValue() js.Value {
	return el.value
}

func (el Element) Children() []Element {
	array := el.value.Get("children")
	elements := make([]Element, array.Length())
	for i := range elements {
		elements[i] = NewElement(array.Index(i))
	}
	return elements
}

func (el Element) Append(node Node) {
	el.value.Call("append", node.NodeValue())
}
