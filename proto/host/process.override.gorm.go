package host

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// BeforeCreate - GORM-specific autogenerated helpers.
func (process *ProcessORM) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	process.Id = id.String()

	// Find any process on this host with the same PID and delete
	existing := &ProcessORM{}
	err = tx.Where(&ProcessORM{
		Pid:    process.Pid,
		HostId: process.HostId,
	}).Delete(existing).Error

	return err
}