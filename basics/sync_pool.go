package main

import (
	"fmt"
	"sync"
)

type Dog struct{ name string }

func (d *Dog) Bark() { fmt.Printf("%s", d.name) }

var dogPack = sync.Pool{
	New: func() interface{} { return &Dog{} },
}

func main() {
	dog := dogPack.Get().(*Dog)
	dog.name = "billy"
	dog.Bark()
	dogPack.Put(dog)
}
