package credential



import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)



// BeforeCreate - GORM-specific autogenerated helpers.
func (origin *OriginORM) BeforeCreate(tx *gorm.DB) (err error) {
    id, err := uuid.NewV4()
	if err != nil {
		return err
	}
    origin.Id = id.String()

    return nil
}



