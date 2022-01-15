package html

import (
	"reflect"
	"syscall/js"
)

type ShadowRootInit struct {
	Mode ShadowRootMode
}

func (init ShadowRootInit) JSValue() js.Value {
	return js.ValueOf(map[string]interface{}{
		"mode": init.Mode,
	})
}

type ShadowRootMode int

const (
	ShadowRootModeOpen ShadowRootMode = iota
)

func (m ShadowRootMode) JSValue() js.Value {
	switch m {
	case ShadowRootModeOpen:
		return js.ValueOf("open")
	}
	return js.ValueOf("")
}

func (el *Element) AttachShadow(init ShadowRootInit) {
	el.value.Call("attachShadow", init)
}

type ShadowRoot struct {
	value js.Value
}

func (el Element) ShadowRoot() ShadowRoot {
	root := el.value.Get("shadowRoot")
	return ShadowRoot{value: root}
}

func (s ShadowRoot) AppendChild(node Node) {
	s.value.Call("appendChild", node.NodeValue())
}

func (s ShadowRoot) GetElementById(id string) Element {
	return NewElement(s.value.Call("getElementById", id))
}

type CustomElement interface {
	ConnectedCallback(element Element)
}

func Define(tagName string, constructor func() CustomElement, templateString string) {
	template := NewTemplateElement(templateString)
	init := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		el := NewElement(this)
		el.AttachShadow(ShadowRootInit{Mode: ShadowRootModeOpen})
		shadowroot := el.ShadowRoot()

		element := constructor()
		this.Set("connectedFunc", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			t := reflect.TypeOf(element).Elem()
			v := reflect.ValueOf(element).Elem()

			for i := 0; i < t.NumField(); i++ {
				f := t.Field(i)
				id := f.Tag.Get("id")
				if id != "" {
					v.Field(i).Set(reflect.ValueOf(shadowroot.GetElementById(id)))
				}
			}

			element.ConnectedCallback(NewElement(this))
			return nil
		}))

		shadowroot.AppendChild(template.CloneNode())

		return nil
	})
	js.Global().Call("makeComponent", tagName, init)
}
