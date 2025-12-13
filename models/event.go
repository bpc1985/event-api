package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64     `gorm:"primaryKey"`
	Name        string    `gorm:"not null" binding:"required"`
	Description string    `gorm:"not null" binding:"required"`
	Location    string    `gorm:"not null" binding:"required"`
	DateTime    time.Time `gorm:"not null" binding:"required"`
	UserID      int64     `gorm:"not null"`
}

func (e *Event) Save() error {
	result := db.DB.Create(e)
	return result.Error
}

func GetAllEvents() ([]Event, error) {
	var events []Event
	result := db.DB.Find(&events)
	return events, result.Error
}

func GetEventByID(id int64) (*Event, error) {
	var event Event
	result := db.DB.First(&event, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &event, nil
}

func (event *Event) Update() error {
	result := db.DB.Model(event).Updates(Event{
		Name:        event.Name,
		Description: event.Description,
		Location:    event.Location,
		DateTime:    event.DateTime,
	})
	return result.Error
}

func (event *Event) Delete() error {
	result := db.DB.Delete(event)
	return result.Error
}

func (event *Event) Register(userID int64) error {
	registration := Registration{
		EventID: event.ID,
		UserID:  userID,
	}
	result := db.DB.Create(&registration)
	return result.Error
}

func (event *Event) Cancel(userID int64) error {
	result := db.DB.Where("event_id = ? AND user_id = ?", event.ID, userID).Delete(&Registration{})
	return result.Error
}
