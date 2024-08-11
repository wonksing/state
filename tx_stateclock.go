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

// SetState sets newState to underlying State. It implements state.OnTxStateChanged function.
// Avoid calling this method directly.
func (e *StateClock) SetState(newState types.TxState) error {
	e.State = newState
	return nil
}

func (e *StateClock) EqualStateSm(s types.TxState) bool {
	if e == nil {
		return false
	}
	e.checkAndInitStateMachine()
	return e.stateMachine.State == s
}

func (e *StateClock) IsPendingKindSm() bool {
	if e == nil {
		return false
	}
	e.checkAndInitStateMachine()
	return e.stateMachine.IsPendingKind()
}

func (e *StateClock) IsPendingSm() bool {
	if e == nil {
		return false
	}
	e.checkAndInitStateMachine()
	return e.stateMachine.IsPending()
}

func (e *StateClock) IsModifyPendingSm() bool {
	if e == nil {
		return false
	}
	e.checkAndInitStateMachine()
	return e.stateMachine.IsModifyPending()
}

func (e *StateClock) IsRemovePendingSm() bool {
	if e == nil {
		return false
	}
	e.checkAndInitStateMachine()
	return e.stateMachine.IsRemovePending()
}

func (e *StateClock) IsActiveSm() bool {
	if e == nil {
		return false
	}

	e.checkAndInitStateMachine()
	return e.stateMachine.IsActive()
}

func (e *StateClock) IsCanceledSm() bool {
	if e == nil {
		return false
	}

	e.checkAndInitStateMachine()
	return e.stateMachine.IsCanceled()
}

func (e *StateClock) IsRemovedSm() bool {
	if e == nil {
		return false
	}

	e.checkAndInitStateMachine()
	return e.stateMachine.IsRemoved()
}

func (e *StateClock) SetStateSm(newState types.TxState) error {
	if e == nil {
		return errors.New("not initialized")
	}
	e.checkAndInitStateMachine()
	err := e.stateMachine.SetState(newState, e)
	if err != nil {
		return err
	}
	e.Tick()
	return nil
}

func (e *StateClock) ForceStateSm(newState types.TxState) error {
	if e == nil {
		return errors.New("not initialized")
	}
	e.checkAndInitStateMachine()
	err := e.stateMachine.ForceState(newState, e)
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
	e.checkAndInitStateMachine()
	err := e.stateMachine.SetState(types.PendingTxState, e)
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
	e.checkAndInitStateMachine()
	err := e.stateMachine.SetState(types.ModifyPendingTxState, e)
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
	e.checkAndInitStateMachine()
	err := e.stateMachine.SetState(types.RemovePendingTxState, e)
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
	e.checkAndInitStateMachine()
	err := e.stateMachine.SetState(types.InactivePendingTxState, e)
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
	e.checkAndInitStateMachine()
	err := e.stateMachine.SetState(types.ActivePendingTxState, e)
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
	e.checkAndInitStateMachine()
	err := e.stateMachine.Approve(e)
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
	e.checkAndInitStateMachine()
	err := e.stateMachine.Cancel(e)
	if err != nil {
		return err
	}
	e.Tick()
	return nil
}

// checkAndInitStateMachine check and initialize e.stateMachine.
// It initialize e.stateMachine with InactiveTxState if e.State is invalid.
func (e *StateClock) checkAndInitStateMachine() {
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

func (e *StateClock) ResetTick() {
	e.VersionTicked = false
}
