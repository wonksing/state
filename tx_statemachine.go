package state

import (
	"errors"

	"github.com/wonksing/state/internal"
	"github.com/wonksing/state/types"
)

type TxStateMachine struct {
	State        types.TxState            `gorm:"column:state;type:string;size:32;comment:state" json:"state,omitempty"`
	stateMachine *internal.TxStateMachine `gorm:"-:all" json:"-"`
}

// SetState sets newState to underlying State. It implements state.OnTxStateChanged function.
// Avoid calling this method directly.
func (e *TxStateMachine) SetState(newState types.TxState) error {
	e.State = newState
	return nil
}

func (e *TxStateMachine) SmIsState(s types.TxState) bool {
	if e == nil {
		return false
	}
	e.checkAndInitStateMachine()
	return e.stateMachine.State == s
}

func (e *TxStateMachine) SmIsPendingKind() bool {
	if e == nil {
		return false
	}
	e.checkAndInitStateMachine()
	return e.stateMachine.IsPendingKind()
}

func (e *TxStateMachine) SmIsPending() bool {
	if e == nil {
		return false
	}
	e.checkAndInitStateMachine()
	return e.stateMachine.IsPending()
}

func (e *TxStateMachine) SmIsModifyPending() bool {
	if e == nil {
		return false
	}
	e.checkAndInitStateMachine()
	return e.stateMachine.IsModifyPending()
}

func (e *TxStateMachine) SmIsRemovePending() bool {
	if e == nil {
		return false
	}
	e.checkAndInitStateMachine()
	return e.stateMachine.IsRemovePending()
}

func (e *TxStateMachine) SmIsActive() bool {
	if e == nil {
		return false
	}

	e.checkAndInitStateMachine()
	return e.stateMachine.IsActive()
}

func (e *TxStateMachine) SmIsCanceled() bool {
	if e == nil {
		return false
	}

	e.checkAndInitStateMachine()
	return e.stateMachine.IsCanceled()
}

func (e *TxStateMachine) SmIsRemoved() bool {
	if e == nil {
		return false
	}

	e.checkAndInitStateMachine()
	return e.stateMachine.IsRemoved()
}

func (e *TxStateMachine) SmSetState(newState types.TxState) error {
	if e == nil {
		return errors.New("not initialized")
	}
	e.checkAndInitStateMachine()
	return e.stateMachine.SetState(newState, e)
}

func (e *TxStateMachine) SmForceState(newState types.TxState) error {
	if e == nil {
		return errors.New("not initialized")
	}
	e.checkAndInitStateMachine()
	return e.stateMachine.ForceState(newState, e)
}

func (e *TxStateMachine) SmApprove() error {
	if e == nil {
		return errors.New("not initialized")
	}
	e.checkAndInitStateMachine()
	return e.stateMachine.Approve(e)
}

func (e *TxStateMachine) SmCancel() error {
	if e == nil {
		return errors.New("not initialized")
	}
	e.checkAndInitStateMachine()
	return e.stateMachine.Cancel(e)
}

// checkAndInitStateMachine check and initialize e.stateMachine.
// It initialize e.stateMachine with InactiveTxState if e.State is invalid.
func (e *TxStateMachine) checkAndInitStateMachine() {
	if e == nil {
		return
	}

	if e.stateMachine == nil {
		var err error
		e.stateMachine, err = internal.NewTxStateMachine(e.State, e)
		if err != nil {
			e.State = types.InactiveTxState
			e.stateMachine, _ = internal.NewTxStateMachine(types.InactiveTxState, e)
			return
		}
	}
}
