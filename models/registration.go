package models

import "example.com/rest-api/db"

type Registration struct {
	ID      int64 `gorm:"primaryKey"`
	EventID int64 `gorm:"not null"`
	UserID  int64 `gorm:"not null"`
}

func GetAttendeesByEventID(eventID int64) ([]User, error) {
	var users []User
	result := db.DB.
		Select("users.id, users.firstname, users.lastname, users.email").
		Joins("JOIN registrations ON users.id = registrations.user_id").
		Where("registrations.event_id = ?", eventID).
		Find(&users)
	return users, result.Error
}

func GetEventsByAttendeeID(userID int64) ([]Event, error) {
	var events []Event
	result := db.DB.
		Joins("JOIN registrations ON events.id = registrations.event_id").
		Where("registrations.user_id = ?", userID).
		Find(&events)
	return events, result.Error
}