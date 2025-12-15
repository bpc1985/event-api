package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

// @Summary Register for an event
// @Description Register authenticated user for a specific event
// @Tags registrations
// @Produce json
// @Param id path int true "Event ID"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /events/{id}/register [post]
func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for the event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Successfully registered for the event."})
}

// @Summary Cancel event registration
// @Description Cancel authenticated user's registration for a specific event
// @Tags registrations
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /events/{id}/register [delete]
func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.Cancel(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration for the event. Try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully canceled registration for the event."})
}

// @Summary Get event attendees
// @Description Get list of users registered for a specific event
// @Tags registrations
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {array} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /events/{id}/attendees [get]
func getAttendeesForEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	attendees, err := models.GetAttendeesByEventID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch attendees for the event. Try again later."})
		return
	}

	context.JSON(http.StatusOK, attendees)
}

// @Summary Get events by attendee
// @Description Get list of events that a specific user has registered for
// @Tags registrations
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} models.Event
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /attendees/{id}/events [get]
func getEventsByAttendee(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user id."})
		return
	}

	events, err := models.GetEventsByAttendeeID(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events for the user. Try again later."})
		return
	}

	context.JSON(http.StatusOK, events)
}