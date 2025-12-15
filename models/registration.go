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
		Joins("JOIN registrations ON users.id = registrations.user_id").
		Where("registrations.event_id = ?", eventID).
		Find(&users)
	return users, result.Error
}

func GetEventsByAttendeeID(userID int64) ([]Event, error) {
	var events []Event
	result := db.DB.Table("events e").
		Select("e.id, e.user_id, e.name, e.description, e.dateTime, e.location").
		Joins("JOIN registrations a ON e.id = a.event_id").
		Where("a.user_id = ?", userID).
		Find(&events)
	return events, result.Error
}