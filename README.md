# state

[![Go Reference](https://pkg.go.dev/badge/github.com/wonksing/state.svg)](https://pkg.go.dev/github.com/wonksing/state)
[![Go Report Card](https://goreportcard.com/badge/github.com/wonksing/state)](https://goreportcard.com/report/github.com/wonksing/state)

`state` is a package to ease the state management of services running in MSA. States include the following:

- pending
  - This is the initial state.
  - Previous state does not exist.
  - The next states could be `active` or `canceled`
- modify_pending
  - This is a state when changing an existing entity(that is in active state) in a transaction.
  - Previous state is `active`.
  - Next state is `active` when the transaction is both approved or canceled.
- active
  - This is one of the final state when the transaction is successfully approved or canceled.
  - Previous states could be `pending`, `modify_pending`, `remove_pending`, `inactive_pending` or `active_pending`.
  - Next state does not exist.
- canceled
  - This is a state after a transaction has been canceled or rolled back.
  - Previous state is `pending`.
  - Next state does not exist.
- remove_pending
  - This is a state when you try to remove an active entity.
  - Previous state is `active`.
  - Next state is `remove` or `active`.
- removed
  - This is a state indicating that an entity has been removed.
  - Previous state is `remove_pending`.
  - Next state does not exist.
- inactive_pending
  - This is a state when you try to inactivate an active entity.
  - Previous state is `active`.
  - Next states could be `inactive` or `active`.
- inactive
  - This is a state indicating that an entity has been inactivated.
  - Previous state is `inactive_pending` or `active_pending`.
  - Next state does not exist.
- active_pending
  - This is a state when you try to activate an inactive entity.
  - Previous state is `inactive`.
  - Next state is `active` or `inactive`.