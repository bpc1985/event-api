package models

type Registration struct {
	ID      int64 `gorm:"primaryKey"`
	EventID int64 `gorm:"not null"`
	UserID  int64 `gorm:"not null"`
}
