package receiver

import "fmt"

func ReceiverInterface() {
	var x IFunc
	x = F1{1}
	x.Update(2)
	x.Call()
	x = &F2{1} // [Pattern]: [Go Pointer Receiver interface] No value type to pointer receiver interface
	x.Update(2)
	x.Call()
}

type IFunc interface {
	Update(x int)
	Call()
}

type F1 struct {
	val int
}

func (f F1) Update(x int) {
	f.val = x // [Pattern]: [Go Receiver] No update to value receiver properties
}
func (f F1) Call() {
	fmt.Println(f.val)
}

type F2 struct {
	val int
}

func (f *F2) Update(x int) {
	f.val = x
}
func (f *F2) Call() {
	fmt.Println(f.val)
}
