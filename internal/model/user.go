package model

type User struct {
	ID    []byte `gorm:"type:binary(16);primaryKey;"`
	Email string
	BaseModel
}
