package main

import (
	"sync"
	"syscall/js"

	_ "go-webcomponents-prototype/components/mytable"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)
	js.Global().Get("window").Call("addEventListener", "beforeunload", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		wg.Done()
		return nil
	}))
	wg.Wait()
}
