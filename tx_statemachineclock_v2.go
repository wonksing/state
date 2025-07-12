package state

import (
	"errors"
	"time"

	"github.com/wonksing/state/internal"
	"github.com/wonksing/state/types"
)

type TxStateMachineClockV2 struct {
	State        types.TxState            `gorm:"column:state;type:string;size:32;comment:state" json:"state,omitempty"`
	stateMachine *internal.TxStateMachine `gorm:"-:all" json:"-"`

	Version       *uint64 `gorm:"column:version;type:uint" json:"version,omitempty"`
	VersionTicked *bool   `gorm:"-:all" json:"-"`

	CreatedAt *time.Time `gorm:"<-:create" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `gorm:"<-" json:"updatedAt,omitempty"`
}

// AssignStateCallback sets newState to underlying State. It implements internal.TxStateAssignor interface.
// DO NOT CALL THIS METHOD DIRECTLY.
func (e *TxStateMachineClockV2) AssignStateCallback(newState types.TxState) error {
	e.State = newState
	return nil
}

func (e *TxStateMachineClockV2) EqualSm(s types.TxState) bool {
	if e == nil {
		return false
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.Equal(s)
}

func (e *TxStateMachineClockV2) IsPendingKindSm() bool {
	if e == nil {
		return false
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsPendingKind()
}

func (e *TxStateMachineClockV2) IsPendingSm() bool {
	if e == nil {
		return false
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsPending()
}

func (e *TxStateMachineClockV2) IsModifyPendingSm() bool {
	if e == nil {
		return false
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsModifyPending()
}

func (e *TxStateMachineClockV2) IsRemovePendingSm() bool {
	if e == nil {
		return false
	}
	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsRemovePending()
}

func (e *TxStateMachineClockV2) IsActiveSm() bool {
	if e == nil {
		return false
	}

	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsActive()
}

func (e *TxStateMachineClockV2) IsCanceledSm() bool {
	if e == nil {
		return false
	}

	if err := e.checkAndInitStateMachine(); err != nil {
		return false
	}
	return e.stateMachine.IsCanceled()
}

func (e *TxStateMachineClockV2) IsRemovedSm() bool {
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

func (e *TxStateMachineClockV2) ForceStateSm(newState types.TxState) error {
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

func (e *TxStateMachineClockV2) PendingSm() error {
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

func (e *TxStateMachineClockV2) ModifyPendingSm() error {
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

func (e *TxStateMachineClockV2) RemovePendingSm() error {
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

func (e *TxStateMachineClockV2) InactivePendingSm() error {
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

func (e *TxStateMachineClockV2) ActivePendingSm() error {
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

func (e *TxStateMachineClockV2) ApproveSm() error {
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

func (e *TxStateMachineClockV2) CancelSm() error {
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
func (e *TxStateMachineClockV2) checkAndInitStateMachine() error {
	if e == nil {
		return errors.New("not initialized")
	}
	if e.State == "" {
		e.State = types.PendingTxState
	}
	return e.checkAndInitStateMachineWithState(e.State)
}
func (e *TxStateMachineClockV2) checkAndInitStateMachineWithState(s types.TxState) error {
	if e == nil {
		return errors.New("not initialized")
	}

	if e.stateMachine == nil {
		if e.State != "" {
			s = e.State
		}

		var err error
		e.stateMachine, err = internal.NewTxStateMachine(s, e)
		if err != nil {
			return err
		}
	}

	if e.VersionTicked == nil {
		e.VersionTicked = new(bool)
	}
	if e.Version == nil {
		e.Version = new(uint64)
	}

	return nil
}

// Tick increments Version and set current time to CreatedAt and UpdatedAt.
// It returns immediately if Version is already incremented.
func (e *TxStateMachineClockV2) Tick() {
	if e.VersionTicked == nil {
		e.VersionTicked = new(bool)
	} else if *e.VersionTicked {
		return
	}
	*e.VersionTicked = true

	if e.Version == nil {
		e.Version = new(uint64)
	}
	*e.Version++

	now := time.Now().Round(time.Microsecond)
	if e.CreatedAt == nil {
		e.CreatedAt = &now
	}
	e.UpdatedAt = &now
}

func (e *TxStateMachineClockV2) ResetTicked() {
	if e.VersionTicked == nil {
		e.VersionTicked = new(bool)
	}
	*e.VersionTicked = false
}
