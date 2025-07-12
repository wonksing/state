package state

import "time"

// TxClockV2
type TxClockV2 struct {
	Version       *uint64 `gorm:"column:version;type:uint" json:"version,omitempty"`
	VersionTicked *bool   `gorm:"-:all" json:"-"`

	CreatedAt *time.Time `gorm:"<-:create" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `gorm:"<-" json:"updatedAt,omitempty"`
}

// Tick increments Version and set current time to CreatedAt and UpdatedAt.
// It returns immediately if Version is already incremented.
func (e *TxClockV2) Tick() {
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

func (e *TxClockV2) ResetTicked() {
	if e.VersionTicked == nil {
		e.VersionTicked = new(bool)
	}
	*e.VersionTicked = false
}
