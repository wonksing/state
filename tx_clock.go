package state

import "time"

// TxClock
type TxClock struct {
	Version       uint64 `gorm:"column:version;type:uint" json:"version,omitempty"`
	VersionTicked bool   `gorm:"-:all" json:"-"`

	CreatedAt *time.Time `gorm:"<-:create;index:idx_created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"<-;index:idx_updated_at" json:"updated_at,omitempty"`
}

// Tick increments Version and set current time to CreatedAt and UpdatedAt.
// It returns immediately if Version is already incremented.
func (e *TxClock) Tick() {
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

func (e *TxClock) ResetTick() {
	e.VersionTicked = false
}
