package credential



import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)



// BeforeCreate - GORM-specific autogenerated helpers.
func (public *PublicORM) BeforeCreate(tx *gorm.DB) (err error) {
    id, err := uuid.NewV4()
	if err != nil {
		return err
	}
    public.Id = id.String()

    return nil
}



