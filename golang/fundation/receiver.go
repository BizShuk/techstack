package fundation

import (
	"fmt"
	"time"
)

type PointerReceiver int

func (p *PointerReceiver) Call() {}

type ValueReceiver int

func (p ValueReceiver) Call() {}

// [Notice]: [Go] pointer receiver issue with loop
func LoopEndValueWithPointerReceiver() {

	tasks := []Task{ // [Solution]: Use *Task instead of Task
		{name: "Shuk", val: 1},
		{name: "An", val: 2},
		{name: "How", val: 3},
	}

	for i, task := range tasks {
		// task := task       // [Hint]: comment out to show correct answer
		go task.Call()     // goroutine is not necessary
		go tasks[i].Call() // [Solution]

	}
	time.Sleep(time.Second * 1) // time.Sleep is not necessary
}

type Task struct {
	name string
	val  int
}

func (this *Task) Call() {
	fmt.Println(this.name, this.val)
}
