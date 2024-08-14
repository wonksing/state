package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wonksing/state/types"
)

func Test_TxState_SetState_when_empty(t *testing.T) {
	m, _ := NewTxStateMachine(types.ActiveTxState, nil)
	m.State = ""
	err := m.SetState(types.PendingTxState)
	require.Nil(t, err)
}

func Test_TxStateValidate(t *testing.T) {
	_, err := NewTxStateMachine(types.TxState("a"), nil)
	require.NotNil(t, err)
}

func Test_TxStateMachine(t *testing.T) {
	m, err := NewTxStateMachine(types.PendingTxState, nil)
	require.Nil(t, err)

	err = m.SetState(types.ModifyPendingTxState)
	require.Nil(t, err)
	require.EqualValues(t, types.PendingTxState, m.State)

	err = m.SetState(types.RemovePendingTxState)
	require.NotNil(t, err)

	err = m.SetState(types.InactivePendingTxState)
	require.NotNil(t, err)

	err = m.ForceState(types.InactivePendingTxState)
	require.Nil(t, err)
	require.EqualValues(t, types.InactivePendingTxState, m.State)
}

func Test_DefaultState(t *testing.T) {
	e := newTestEntity()
	require.EqualValues(t, types.PendingTxState, e.State)
	require.EqualValues(t, e.State, e.StateMachine.State)

	err := e.Approve()
	require.Nil(t, err)
	require.EqualValues(t, types.ActiveTxState, e.State)
	require.EqualValues(t, e.State, e.StateMachine.State)

	err = e.Cancel()
	require.NotNil(t, err)
	require.EqualValues(t, types.ActiveTxState, e.State)
	require.EqualValues(t, e.State, e.StateMachine.State)

	e.SetMachineState(types.ModifyPendingTxState)
	require.EqualValues(t, types.ModifyPendingTxState, e.State)
	require.EqualValues(t, e.State, e.StateMachine.State)
	err = e.Cancel()
	require.Nil(t, err)
	require.EqualValues(t, types.ActiveTxState, e.State)
	require.EqualValues(t, e.State, e.StateMachine.State)

	e.StateMachine.ForceState(types.PendingTxState)
	e.SetMachineState(types.PendingTxState)
	require.EqualValues(t, types.PendingTxState, e.State)
	require.EqualValues(t, e.State, e.StateMachine.State)
	err = e.Cancel()
	require.Nil(t, err)
	require.EqualValues(t, types.CanceledTxState, e.State)
	require.EqualValues(t, e.State, e.StateMachine.State)

	e = newTestEntity()
	e.Approve()
	e.SetMachineState(types.RemovePendingTxState)
	require.EqualValues(t, types.RemovePendingTxState, e.State)
	require.EqualValues(t, e.State, e.StateMachine.State)
	err = e.Approve()
	require.Nil(t, err)
	require.EqualValues(t, types.RemovedTxState, e.State)
	require.EqualValues(t, e.State, e.StateMachine.State)

	e = newTestEntity()
	e.Approve()
	e.SetMachineState(types.RemovePendingTxState)
	require.EqualValues(t, types.RemovePendingTxState, e.State)
	require.EqualValues(t, e.State, e.StateMachine.State)
	err = e.Cancel()
	require.Nil(t, err)
	require.EqualValues(t, types.ActiveTxState, e.State)
	require.EqualValues(t, e.State, e.StateMachine.State)

	e.SetMachineState(types.ActiveTxState)
	require.EqualValues(t, types.ActiveTxState, e.State)
	require.EqualValues(t, e.State, e.StateMachine.State)
	err = e.Cancel()
	require.NotNil(t, err)
	require.EqualValues(t, types.ActiveTxState, e.State)
	require.EqualValues(t, e.State, e.StateMachine.State)
}

type testEntity struct {
	State        types.TxState
	StateMachine *TxStateMachine
}

func newTestEntity() *testEntity {
	e := testEntity{}
	e.StateMachine, _ = NewTxStateMachine(types.PendingTxState, &e)
	return &e
}

func (e *testEntity) String() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *testEntity) AssignStateCallback(s types.TxState) error {
	e.State = s
	// e.StateMachine.SetState(s)
	return nil
}

func (e *testEntity) SetMachineState(s types.TxState) {
	e.StateMachine.SetState(s)
	// e.SetState(s)
}

func (e *testEntity) Approve() error {
	return e.StateMachine.Approve()
}

func (e *testEntity) Cancel() error {
	return e.StateMachine.Cancel()
}
