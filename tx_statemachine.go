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

// AssignStateCallback sets newState to underlying State. It implements internal.TxStateAssignor interface.
// DO NOT CALL THIS METHOD DIRECTLY.
func (e *TxStateMachine) AssignStateCallback(newState types.TxState) error {
	e.State = newState
	return nil
}

func (e *TxStateMachine) EqualSm(s types.TxState) bool {
	if e == nil {
		return false
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.Equal(s)
}

func (e *TxStateMachine) IsPendingKindSm() bool {
	if e == nil {
		return false
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsPendingKind()
}

func (e *TxStateMachine) IsPendingSm() bool {
	if e == nil {
		return false
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsPending()
}

func (e *TxStateMachine) IsModifyPendingSm() bool {
	if e == nil {
		return false
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsModifyPending()
}

func (e *TxStateMachine) IsRemovePendingSm() bool {
	if e == nil {
		return false
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsRemovePending()
}

func (e *TxStateMachine) IsActiveSm() bool {
	if e == nil {
		return false
	}

	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsActive()
}

func (e *TxStateMachine) IsCanceledSm() bool {
	if e == nil {
		return false
	}

	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsCanceled()
}

func (e *TxStateMachine) IsRemovedSm() bool {
	if e == nil {
		return false
	}

	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsRemoved()
}

// func (e *TxStateMachine) SetStateSm(newState types.TxState) error {
// 	if e == nil {
// 		return errors.New("not initialized")
// 	}
// 	if err := e.checkAndInitStateMachine(); err != nil {
// 		return err
// 	}
// 	return e.stateMachine.SetState(newState)
// }

func (e *TxStateMachine) ForceStateSm(newState types.TxState) error {
	if e == nil {
		return errors.New("not initialized")
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return err
	}
	return e.stateMachine.ForceState(newState)
}

func (e *TxStateMachine) PendingSm() error {
	if e == nil {
		return errors.New("not initialized")
	}
	if err := e.checkAndInitStateMachineWithState(types.PendingTxState); err != nil {
		return err
	}
	return e.stateMachine.SetState(types.PendingTxState)
}

func (e *TxStateMachine) ModifyPendingSm() error {
	if e == nil {
		return errors.New("not initialized")
	}
	if err := e.checkAndInitStateMachineWithState(types.ModifyPendingTxState); err != nil {
		return err
	}
	return e.stateMachine.SetState(types.ModifyPendingTxState)
}

func (e *TxStateMachine) RemovePendingSm() error {
	if e == nil {
		return errors.New("not initialized")
	}
	if err := e.checkAndInitStateMachineWithState(types.RemovePendingTxState); err != nil {
		return err
	}
	return e.stateMachine.SetState(types.RemovePendingTxState)
}

func (e *TxStateMachine) InactivePendingSm() error {
	if e == nil {
		return errors.New("not initialized")
	}
	if err := e.checkAndInitStateMachineWithState(types.InactivePendingTxState); err != nil {
		return err
	}
	return e.stateMachine.SetState(types.InactivePendingTxState)
}

func (e *TxStateMachine) ActivePendingSm() error {
	if e == nil {
		return errors.New("not initialized")
	}
	if err := e.checkAndInitStateMachineWithState(types.ActivePendingTxState); err != nil {
		return err
	}
	return e.stateMachine.SetState(types.ActivePendingTxState)
}
func (e *TxStateMachine) ApproveSm() error {
	if e == nil {
		return errors.New("not initialized")
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return err
	}
	return e.stateMachine.Approve()
}

func (e *TxStateMachine) CancelSm() error {
	if e == nil {
		return errors.New("not initialized")
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return err
	}
	return e.stateMachine.Cancel()
}

// checkAndInitStateMachine check and initialize e.stateMachine.
// It initializes e.stateMachine with PendingTxState if e.State is empty.
func (e *TxStateMachine) checkAndInitStateMachine() error {
	if e == nil {
		return errors.New("not initialized")
	}
	if e.State == "" {
		e.State = types.PendingTxState
	}
	return e.checkAndInitStateMachineWithState(e.State)
}
func (e *TxStateMachine) checkAndInitStateMachineWithState(s types.TxState) error {
	if e == nil {
		return errors.New("not initialized")
	}

	if e.stateMachine == nil {
		var err error
		e.stateMachine, err = internal.NewTxStateMachine(s, e)
		if err != nil {
			return err
		}
	}
	return nil
}
