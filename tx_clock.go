package state

import "time"

// TxClock
type TxClock struct {
	Version       *uint64 `gorm:"column:version;type:uint" json:"version,omitempty"`
	VersionTicked *bool   `gorm:"-:all" json:"-"`

	CreatedAt *time.Time `gorm:"<-:create" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"<-" json:"updated_at,omitempty"`
}

// Tick increments Version and set current time to CreatedAt and UpdatedAt.
// It returns immediately if Version is already incremented.
func (e *TxClock) Tick() {
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

func (e *TxClock) ResetTicked() {
	if e.VersionTicked == nil {
		e.VersionTicked = new(bool)
	}
	*e.VersionTicked = false
}
