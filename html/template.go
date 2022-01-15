package html

import "syscall/js"

type TemplateElement struct {
	value js.Value
}

func NewTemplateElement(innerHTML string) TemplateElement {
	template := js.Global().Get("document").Call("createElement", "template")
	template.Set("innerHTML", innerHTML)
	return TemplateElement{value: template}
}

func (te TemplateElement) NodeValue() js.Value {
	return te.value
}

func (te TemplateElement) CloneNode() Node {
	node := te.value.Get("content").Call("cloneNode", true)
	return Element{value: node}
}
