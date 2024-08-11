package types

type TxState string

const (
	// MUST NOT be more than 32 characters
	// pending > active(canceled)
	// active > modify_pending > active
	// active > remove_pending > removed
	// active > inactive_pending > inactive
	// inactive > active_pending > active

	PendingTxState         TxState = "pending"
	ModifyPendingTxState   TxState = "modify_pending"
	ActiveTxState          TxState = "active"
	CanceledTxState        TxState = "canceled"
	RemovePendingTxState   TxState = "remove_pending"
	RemovedTxState         TxState = "removed"
	InactivePendingTxState TxState = "inactive_pending"
	InactiveTxState        TxState = "inactive"
	ActivePendingTxState   TxState = "active_pending"
)
