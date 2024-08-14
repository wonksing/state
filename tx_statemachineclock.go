package state

import (
	"errors"
	"time"

	"github.com/wonksing/state/internal"
	"github.com/wonksing/state/types"
)

type TxStateMachineClock struct {
	State        types.TxState            `gorm:"column:state;type:string;size:32;comment:state" json:"state,omitempty"`
	stateMachine *internal.TxStateMachine `gorm:"-:all" json:"-"`

	Version       uint64 `gorm:"column:version;type:uint" json:"version,omitempty"`
	VersionTicked bool   `gorm:"-:all" json:"-"`

	CreatedAt *time.Time `gorm:"<-:create;index:idx_created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"<-;index:idx_updated_at" json:"updated_at,omitempty"`
}

// AssignStateCallback sets newState to underlying State. It implements internal.TxStateAssignor interface.
// DO NOT CALL THIS METHOD DIRECTLY.
func (e *TxStateMachineClock) AssignStateCallback(newState types.TxState) error {
	e.State = newState
	return nil
}

func (e *TxStateMachineClock) EqualSm(s types.TxState) bool {
	if e == nil {
		return false
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.Equal(s)
}

func (e *TxStateMachineClock) IsPendingKindSm() bool {
	if e == nil {
		return false
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsPendingKind()
}

func (e *TxStateMachineClock) IsPendingSm() bool {
	if e == nil {
		return false
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsPending()
}

func (e *TxStateMachineClock) IsModifyPendingSm() bool {
	if e == nil {
		return false
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsModifyPending()
}

func (e *TxStateMachineClock) IsRemovePendingSm() bool {
	if e == nil {
		return false
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsRemovePending()
}

func (e *TxStateMachineClock) IsActiveSm() bool {
	if e == nil {
		return false
	}

	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsActive()
}

func (e *TxStateMachineClock) IsCanceledSm() bool {
	if e == nil {
		return false
	}

	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsCanceled()
}

func (e *TxStateMachineClock) IsRemovedSm() bool {
	if e == nil {
		return false
	}

	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsRemoved()
}

// func (e *StateClock) SetStateSm(newState types.TxState) error {
// 	if e == nil {
// 		return errors.New("not initialized")
// 	}
// 	err := e.checkAndInitStateMachine()
// 	if err != nil {
// 		return err
// 	}
// 	err = e.stateMachine.SetState(newState)
// 	if err != nil {
// 		return err
// 	}
// 	e.Tick()
// 	return nil
// }

func (e *TxStateMachineClock) ForceStateSm(newState types.TxState) error {
	if e == nil {
		return errors.New("not initialized")
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return err
	}
	err := e.stateMachine.ForceState(newState)
	if err != nil {
		return err
	}
	e.Tick()
	return nil
}

func (e *TxStateMachineClock) PendingSm() error {
	if e == nil {
		return errors.New("not initialized")
	}
	if err := e.checkAndInitStateMachineWithState(types.PendingTxState); err != nil {
		return err
	}
	err := e.stateMachine.SetState(types.PendingTxState)
	if err != nil {
		return err
	}
	e.Tick()
	return nil
}

func (e *TxStateMachineClock) ModifyPendingSm() error {
	if e == nil {
		return errors.New("not initialized")
	}
	if err := e.checkAndInitStateMachineWithState(types.ModifyPendingTxState); err != nil {
		return err
	}
	err := e.stateMachine.SetState(types.ModifyPendingTxState)
	if err != nil {
		return err
	}
	e.Tick()
	return nil
}

func (e *TxStateMachineClock) RemovePendingSm() error {
	if e == nil {
		return errors.New("not initialized")
	}
	if err := e.checkAndInitStateMachineWithState(types.RemovePendingTxState); err != nil {
		return err
	}
	err := e.stateMachine.SetState(types.RemovePendingTxState)
	if err != nil {
		return err
	}
	e.Tick()
	return nil
}

func (e *TxStateMachineClock) InactivePendingSm() error {
	if e == nil {
		return errors.New("not initialized")
	}
	if err := e.checkAndInitStateMachineWithState(types.InactivePendingTxState); err != nil {
		return err
	}
	err := e.stateMachine.SetState(types.InactivePendingTxState)
	if err != nil {
		return err
	}
	e.Tick()
	return nil
}

func (e *TxStateMachineClock) ActivePendingSm() error {
	if e == nil {
		return errors.New("not initialized")
	}
	if err := e.checkAndInitStateMachineWithState(types.ActivePendingTxState); err != nil {
		return err
	}
	err := e.stateMachine.SetState(types.ActivePendingTxState)
	if err != nil {
		return err
	}
	e.Tick()
	return nil
}

func (e *TxStateMachineClock) ApproveSm() error {
	if e == nil {
		return errors.New("not initialized")
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return err
	}
	err := e.stateMachine.Approve()
	if err != nil {
		return err
	}
	e.Tick()
	return nil
}

func (e *TxStateMachineClock) CancelSm() error {
	if e == nil {
		return errors.New("not initialized")
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return err
	}
	err := e.stateMachine.Cancel()
	if err != nil {
		return err
	}
	e.Tick()
	return nil
}

// checkAndInitStateMachine check and initialize e.stateMachine.
// It initializes e.stateMachine with PendingTxState if e.State is empty.
func (e *TxStateMachineClock) checkAndInitStateMachine() error {
	if e == nil {
		return errors.New("not initialized")
	}
	if e.State == "" {
		e.State = types.PendingTxState
	}
	return e.checkAndInitStateMachineWithState(e.State)
}
func (e *TxStateMachineClock) checkAndInitStateMachineWithState(s types.TxState) error {
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

// Tick increments Version and set current time to CreatedAt and UpdatedAt.
// It returns immediately if Version is already incremented.
func (e *TxStateMachineClock) Tick() {
	if e.VersionTicked {
		return
	}
	e.VersionTicked = true
	e.Version++

	now := time.Now()
	if e.CreatedAt == nil {
		e.CreatedAt = &now
	}
	e.UpdatedAt = &now
}

func (e *TxStateMachineClock) ResetTicked() {
	e.VersionTicked = false
}
