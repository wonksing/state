package internal

import (
	"errors"

	"github.com/wonksing/state/types"
)

type TxStateTransitioner interface {
	Approve() (next types.TxState, err error)
	Cancel() (next types.TxState, err error)
}

type TxStateSetter interface {
	AssignState(s types.TxState) error
}

type OnTxStateChanged func(state types.TxState)

// NewTxStateMachine
func NewTxStateMachine(initState types.TxState, setter TxStateSetter) (*TxStateMachine, error) {
	err := validateTxState(initState)
	if err != nil {
		return nil, err
	}

	machine := &TxStateMachine{
		State:  initState,
		setter: setter,
	}

	machine.m = make(map[types.TxState]TxStateTransitioner)
	machine.m[types.PendingTxState] = _pendingTxState
	machine.m[types.ModifyPendingTxState] = _modifyPendingTxState
	machine.m[types.ActiveTxState] = _activeTxState
	machine.m[types.CanceledTxState] = _canceledTxState
	machine.m[types.RemovePendingTxState] = _removePendingTxState
	machine.m[types.RemovedTxState] = _removedTxState
	machine.m[types.InactivePendingTxState] = _inactivePendingTxState
	machine.m[types.InactiveTxState] = _inactiveTxState
	machine.m[types.ActivePendingTxState] = _activePendingTxState

	err = machine.SetState(initState)
	return machine, err
}

type TxStateMachine struct {
	State  types.TxState
	m      map[types.TxState]TxStateTransitioner
	setter TxStateSetter
}

func (m TxStateMachine) Current() types.TxState {
	return m.State
}

func (m TxStateMachine) Equal(v types.TxState) bool {
	return m.State == v
}

func (m TxStateMachine) IsActive() bool {
	return m.State == types.ActiveTxState
}

func (m TxStateMachine) IsModifyPending() bool {
	return m.State == types.ModifyPendingTxState
}

func (m TxStateMachine) IsRemoved() bool {
	return m.State == types.RemovedTxState
}

func (m TxStateMachine) IsCanceled() bool {
	return m.State == types.CanceledTxState
}

func (m TxStateMachine) IsInactive() bool {
	return m.State == types.InactiveTxState
}

func (m TxStateMachine) IsPending() bool {
	return m.State == types.PendingTxState
}

func (m TxStateMachine) IsRemovePending() bool {
	return m.State == types.RemovePendingTxState
}

func (m TxStateMachine) IsInactivePending() bool {
	return m.State == types.InactivePendingTxState
}

func (m TxStateMachine) IsActivePending() bool {
	return m.State == types.ActivePendingTxState
}

func (m TxStateMachine) IsPendingKind() bool {
	if m.State == types.PendingTxState || m.State == types.ActivePendingTxState ||
		m.State == types.ModifyPendingTxState || m.State == types.RemovePendingTxState || m.State == types.InactivePendingTxState {
		return true
	}
	return false
}

// SetState sets newState to m.State.
// If newState is one of types.ModifyPendingTxState, types.RemovePendingTxState or types.InactivePendingTxState,
// m.State should be types.ActiveTxState.
// If newState is equal to m.State, it returns nil.
func (m *TxStateMachine) SetState(newState types.TxState) error {
	err := validateTxState(newState)
	if err != nil {
		return err
	}

	if newState == m.State {
		if m.setter != nil {
			return m.setter.AssignState(newState)
		}
		return nil
	}

	switch newState {
	case types.ModifyPendingTxState:
		switch m.State {
		case types.ActiveTxState:
		case types.PendingTxState:
			return nil
		default:
			return errors.New("unable to set state")
		}
	case types.RemovePendingTxState:
		switch m.State {
		case types.ActiveTxState:
		default:
			return errors.New("unable to set state")
		}
	case types.InactivePendingTxState:
		switch m.State {
		case types.ActiveTxState:
		default:
			return errors.New("unable to set state")
		}
	case types.ActivePendingTxState:
		switch m.State {
		case types.InactiveTxState:
		default:
			return errors.New("unable to set state")
		}
	case types.PendingTxState:
		if string(m.State) != "" {
			return errors.New("unable to set state")
		}
	default:
		return errors.New("unable to set state")
	}

	m.State = newState
	if m.setter != nil {
		return m.setter.AssignState(newState)
	}
	return nil
}

func (m *TxStateMachine) ForceState(newState types.TxState) error {
	err := validateTxState(newState)
	if err != nil {
		return err
	}

	m.State = newState
	if m.setter != nil {
		return m.setter.AssignState(newState)
	}
	return nil
}

func (m *TxStateMachine) Approve() error {
	if v, ok := m.m[m.State]; ok {
		next, err := v.Approve()
		if err != nil {
			return err
		}
		m.State = next
		if m.setter != nil {
			return m.setter.AssignState(m.State)
		}
		return nil
	}

	return errors.New("current state was not initialized")
}

func (m *TxStateMachine) Cancel() error {
	if v, ok := m.m[m.State]; ok {
		next, err := v.Cancel()
		if err != nil {
			return err
		}
		m.State = next
		if m.setter != nil {
			return m.setter.AssignState(m.State)
		}
		return nil
	}

	return errors.New("current state was not initialized")
}

func validateTxState(state types.TxState) (err error) {
	switch state {
	case types.PendingTxState:
	case types.ModifyPendingTxState:
	case types.ActiveTxState:
	case types.CanceledTxState:
	case types.RemovePendingTxState:
	case types.RemovedTxState:
	case types.InactivePendingTxState:
	case types.InactiveTxState:
	case types.ActivePendingTxState:
	default:
		return errors.New("state is invalid")
	}
	return nil
}
