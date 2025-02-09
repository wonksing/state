package state

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TxStateMachineClock_Tick(t *testing.T) {
	e := &TxStateMachineClock{}
	e.Tick()
	require.Equal(t, uint64(1), *e.Version)

	e = &TxStateMachineClock{}
	require.Nil(t, e.PendingSm())
	require.Equal(t, uint64(1), *e.Version)
	require.True(t, *e.VersionTicked)
}
