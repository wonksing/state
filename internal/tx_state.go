package internal

import (
	"errors"

	"github.com/wonksing/state/types"
)

var (
	_pendingTxState         = newPendingTxState(types.ActiveTxState, types.CanceledTxState)
	_modifyPendingTxState   = newModifyPendingTxState(types.ActiveTxState, types.ActiveTxState)
	_activeTxState          = newActiveTxState(types.ActiveTxState, types.CanceledTxState)
	_canceledTxState        = newCanceledTxState(types.CanceledTxState, types.CanceledTxState)
	_removePendingTxState   = newRemovePendingTxState(types.RemovedTxState, types.ActiveTxState)
	_removedTxState         = newRemovedTxState(types.RemovedTxState, types.RemovedTxState)
	_inactivePendingTxState = newInactivePendingTxState(types.InactiveTxState, types.ActiveTxState)
	_inactiveTxState        = newInactiveTxState(types.InactiveTxState, types.InactiveTxState)
	_activePendingTxState   = newActivePendingTxState(types.ActiveTxState, types.InactiveTxState)
)

// newPendingTxState
func newPendingTxState(approved, canceled types.TxState) *pendingTxState {
	return &pendingTxState{
		approvedState: approved,
		canceledState: canceled,
	}
}

type pendingTxState struct {
	approvedState types.TxState
	canceledState types.TxState
}

func (s *pendingTxState) Approve() (next types.TxState, err error) {
	next = s.approvedState
	return
}

// Cancel changes pending to canceled state
func (s *pendingTxState) Cancel() (next types.TxState, err error) {
	next = s.canceledState
	return
}

// newTxPendingState
func newModifyPendingTxState(approved, canceled types.TxState) *modifyPendingTxState {
	return &modifyPendingTxState{
		approvedState: approved,
		canceledState: canceled,
	}
}

type modifyPendingTxState struct {
	approvedState types.TxState
	canceledState types.TxState
}

func (s *modifyPendingTxState) Approve() (next types.TxState, err error) {
	next = s.approvedState
	return
}

func (s *modifyPendingTxState) Cancel() (next types.TxState, err error) {
	next = s.canceledState
	return
}

// newActiveTxState
func newActiveTxState(approved, canceled types.TxState) *activeTxState {
	return &activeTxState{
		approvedState: approved,
		canceledState: canceled,
	}
}

type activeTxState struct {
	approvedState types.TxState
	canceledState types.TxState
}

func (s *activeTxState) Approve() (next types.TxState, err error) {
	err = errors.New("already active")
	return
}

// Cancel always returns error when state is active
func (s *activeTxState) Cancel() (next types.TxState, err error) {
	err = errors.New("cannot cancel active state")
	// next = s.canceledState
	return
}

// newCanceledTxState
func newCanceledTxState(approved, canceled types.TxState) *canceledTxState {
	return &canceledTxState{
		approvedState: approved,
		canceledState: canceled,
	}
}

type canceledTxState struct {
	approvedState types.TxState
	canceledState types.TxState
}

func (s *canceledTxState) Approve() (next types.TxState, err error) {
	err = errors.New("already canceled")
	return
}

func (s *canceledTxState) Cancel() (next types.TxState, err error) {
	err = errors.New("cannot cancel canceled state")
	return
}

// newRemovePendingTxState
func newRemovePendingTxState(approved, canceled types.TxState) *removePendingTxState {
	return &removePendingTxState{
		approvedState: approved,
		canceledState: canceled,
	}
}

type removePendingTxState struct {
	approvedState types.TxState
	canceledState types.TxState
}

func (s *removePendingTxState) Approve() (next types.TxState, err error) {
	next = s.approvedState
	return
}

func (s *removePendingTxState) Cancel() (next types.TxState, err error) {
	next = s.canceledState
	return
}

// newRemovedTxState
func newRemovedTxState(approved, canceled types.TxState) *removedTxState {
	return &removedTxState{
		approvedState: approved,
		canceledState: canceled,
	}
}

type removedTxState struct {
	approvedState types.TxState
	canceledState types.TxState
}

func (s *removedTxState) Approve() (next types.TxState, err error) {
	err = errors.New("already removed")
	return
}

func (s *removedTxState) Cancel() (next types.TxState, err error) {
	err = errors.New("already removed")
	return
}

// newInactivePendingTxState
func newInactivePendingTxState(approved, canceled types.TxState) *inactivePendingTxState {
	return &inactivePendingTxState{
		approvedState: approved,
		canceledState: canceled,
	}
}

type inactivePendingTxState struct {
	approvedState types.TxState
	canceledState types.TxState
}

func (s *inactivePendingTxState) Approve() (next types.TxState, err error) {
	next = s.approvedState
	return
}

func (s *inactivePendingTxState) Cancel() (next types.TxState, err error) {
	next = s.canceledState
	return
}

// newInactiveTxState
func newInactiveTxState(approved, canceled types.TxState) *inactiveTxState {
	return &inactiveTxState{
		approvedState: approved,
		canceledState: canceled,
	}
}

type inactiveTxState struct {
	approvedState types.TxState
	canceledState types.TxState
}

func (s *inactiveTxState) Approve() (next types.TxState, err error) {
	err = errors.New("already inactive")
	return
}

func (s *inactiveTxState) Cancel() (next types.TxState, err error) {
	err = errors.New("already inactive")
	return
}

// newActivePendingTxState
func newActivePendingTxState(approved, canceled types.TxState) *activePendingTxState {
	return &activePendingTxState{
		approvedState: approved,
		canceledState: canceled,
	}
}

type activePendingTxState struct {
	approvedState types.TxState
	canceledState types.TxState
}

func (s *activePendingTxState) Approve() (next types.TxState, err error) {
	next = s.approvedState
	return
}

func (s *activePendingTxState) Cancel() (next types.TxState, err error) {
	next = s.canceledState
	return
}
