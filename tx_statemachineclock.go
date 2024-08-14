package state

import (
	"errors"
	"time"

	"github.com/wonksing/state/internal"
	"github.com/wonksing/state/types"
)

type StateClock struct {
	State        types.TxState            `gorm:"column:state;type:string;size:32;comment:state" json:"state,omitempty"`
	stateMachine *internal.TxStateMachine `gorm:"-:all" json:"-"`

	Version       uint64 `gorm:"column:version;type:uint" json:"version,omitempty"`
	VersionTicked bool   `gorm:"-:all" json:"-"`

	CreatedAt *time.Time `gorm:"<-:create;index:idx_created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"<-;index:idx_updated_at" json:"updated_at,omitempty"`
}

// AssignState sets newState to underlying State. It implements state.OnTxStateChanged function.
// DO NOT CALL this method directly.
func (e *StateClock) AssignState(newState types.TxState) error {
	e.State = newState
	return nil
}

func (e *StateClock) EqualSm(s types.TxState) bool {
	if e == nil {
		return false
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.Equal(s)
}

func (e *StateClock) IsPendingKindSm() bool {
	if e == nil {
		return false
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsPendingKind()
}

func (e *StateClock) IsPendingSm() bool {
	if e == nil {
		return false
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsPending()
}

func (e *StateClock) IsModifyPendingSm() bool {
	if e == nil {
		return false
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsModifyPending()
}

func (e *StateClock) IsRemovePendingSm() bool {
	if e == nil {
		return false
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsRemovePending()
}

func (e *StateClock) IsActiveSm() bool {
	if e == nil {
		return false
	}

	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsActive()
}

func (e *StateClock) IsCanceledSm() bool {
	if e == nil {
		return false
	}

	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsCanceled()
}

func (e *StateClock) IsRemovedSm() bool {
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

func (e *StateClock) ForceStateSm(newState types.TxState) error {
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

func (e *StateClock) PendingSm() error {
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

func (e *StateClock) ModifyPendingSm() error {
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

func (e *StateClock) RemovePendingSm() error {
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

func (e *StateClock) InactivePendingSm() error {
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

func (e *StateClock) ActivePendingSm() error {
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

func (e *StateClock) ApproveSm() error {
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

func (e *StateClock) CancelSm() error {
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
// It initialize e.stateMachine with InactiveTxState if e.State is invalid.
func (e *StateClock) checkAndInitStateMachine() error {
	return e.checkAndInitStateMachineWithState(e.State)
}
func (e *StateClock) checkAndInitStateMachineWithState(s types.TxState) error {
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
func (e *StateClock) Tick() {
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

func (e *StateClock) ResetTicked() {
	e.VersionTicked = false
}
