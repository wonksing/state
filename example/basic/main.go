package main

import (
	"fmt"

	"github.com/wonksing/state"
	"github.com/wonksing/state/types"
)

func main() {
	fmt.Println("Hello, playground")
	p := &Person{
		Name: "John",
	}
	p.PendingSm()
	fmt.Println(p.Name, p.State)
	p.ApproveSm()
	fmt.Println(p.Name, p.State)

	p.ForceStateSm(types.PendingTxState)
	fmt.Println(p.Name, p.State)
	p.ApproveSm()
	fmt.Println(p.Name, p.State)

	p.ForceStateSm(types.ModifyPendingTxState)
	fmt.Println(p.Name, p.State)
	p.ApproveSm()
	fmt.Println(p.Name, p.State)

	p.ForceStateSm(types.PendingTxState)
	fmt.Println(p.Name, p.State)
	p.CancelSm()
	fmt.Println(p.Name, p.State)
}

type Person struct {
	Name string `json:"name,omitempty"`
	state.TxStateMachineClock
}
