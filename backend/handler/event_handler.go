// Handler functions for events...
package handler

import (
	"chouseisan/repository"
	"chouseisan/service"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type EventRequest struct {
	Title            string `json:"title"`
	Detail           string `json:"detail"`
	DueEdit          string `json:"due_edit"`
	DateTimeProposal string `json:"dateTimeProposal"`
}

type EventUserTimeslotRequest struct {
	Availability map[string](uint) `json:"availability"`
	Name         string            `json:"name"`
	Email        string            `json:"email"`
	Comment      string            `json:"comment"`
}

type ModifyEventUserTimeslotRequest struct {
	Availability map[string](uint) `json:"availability"`
	Name         string            `json:"name"`
	Comment      string            `json:"comment"`
	UserID       uint              `json:"user_id"`
}

type Schedule struct {
	Name       string `json:"name"`
	ID         uint   `json:"id"`
	Annotation uint   `json:"annotation"`
}

type Participant struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
	UserID  uint   `json:"user_id"`
	Result  []uint `json:"result"`
	Email   string `json:"email"`
}

type EventForm struct {
	ScheduleList []Schedule    `json:"scheduleList"`
	Participants []Participant `json:"participants"`
}

type DeleteTimeslotsRequest struct {
	TimeslotIDs []uint `json:"timeslot_ids"`
}

type EventHandler struct {
	Repo *repository.Repository
}

func NewEventHandler(repo *repository.Repository) *EventHandler {
	return &EventHandler{Repo: repo}
}

func (h *EventHandler) CreateEventHandler(c *gin.Context) {
	// Parse the JSON request
	fmt.Println("here")
	var eventRequest EventRequest
	if err := c.ShouldBindJSON(&eventRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Split proposals by "\n"
	proposals := strings.Split(eventRequest.DateTimeProposal, "\n")
	proposals = filterNotEmpty(proposals)

	// Store new information in DB
	eventID, hostToken, err := h.Repo.CreateEvent(eventRequest.Title, eventRequest.Detail, eventRequest.DueEdit, proposals)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"CreateEventHandler error": "Failed to create event"})
		return
	}

	// Set hostToken by set-cookie header field
	fmt.Println(hostToken)
	// c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie(eventID, hostToken, 3600*24, "/", "http://localhost", false, true)

	// Return the event ID as a response
	c.JSON(http.StatusOK, gin.H{"event_id": eventID})
}

func (h *EventHandler) DeleteEventHandler(c *gin.Context) {
	eventID := c.Param("eventID")

	// check cookie for host token
	tokenString, err := c.Cookie(eventID)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "permission denied, you are not the host of the event"})
		return
	}

	// get event info
	event, err := h.Repo.GetEventByID(eventID)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Event Not Found."})
		return
	}

	// check if the user is the host of the event

	if tokenString != event.HostToken {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "permission denied, you are not the host of the event"})
		return
	}

	// Delete Event from all tables
	err = h.Repo.DeleteEvent(eventID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to delete event"})
		log.Println("Gorm Error:", err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully delteted an event"})
}

func (h *EventHandler) EditTitleDetailHandler(c *gin.Context) {
	eventID := c.Param("eventID")

	// check cookie for host token
	tokenString, err := c.Cookie(eventID)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "permission denied, you are not the host of the event"})
		return
	}

	// get event info
	event, err := h.Repo.GetEventByID(eventID)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Event Not Found."})
		return
	}

	// check if the user is the host of the event

	if tokenString != event.HostToken {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "permission denied, you are not the host of the event"})
		return
	}

	// get request body
	var eventRequest EventRequest
	if err := c.ShouldBindJSON(&eventRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// edit event
	log.Println("request", eventRequest)
	if err := h.Repo.UpdateEventTitleDetail(eventID, eventRequest.Title, eventRequest.Detail, eventRequest.DueEdit); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error editing event title and detail."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully modified event title and detail"})
}

func (h *EventHandler) EditDueHandler(c *gin.Context) {
	eventID := c.Param("eventID")

	// check cookie for host token
	tokenString, err := c.Cookie(eventID)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "permission denied, you are not the host of the event"})
		return
	}

	// get event info
	event, err := h.Repo.GetEventByID(eventID)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Event Not Found."})
		return
	}

	// check if the user is the host of the event

	if tokenString != event.HostToken {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "permission denied, you are not the host of the event"})
		return
	}

	// get request body
	var eventRequest EventRequest
	if err := c.ShouldBindJSON(&eventRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// edit event
	if err := h.Repo.UpdateEventDue(eventID, eventRequest.DueEdit); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error editing edit due."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully modified event due"})
}

func (h *EventHandler) GetTimeslotsHandler(c *gin.Context) {
	eventID := c.Param("eventID")
	// get event info
	event, event_err := h.Repo.GetEventByID(eventID)
	if event_err != nil {
		log.Println(event_err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Event Not Found."})
		return
	}

	//get all timeslots
	timeslots, err := h.Repo.GetTimeslotsByEventID(eventID)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error obtaining timeslots for this event."})
		return
	}

	// make timeslots hash dict

	//timeslotsDict := make(map[string]map[uint]string)

	timeslotsDict := make(map[uint]string)

	for _, timeslot := range timeslots {
		// if _, ok := timeslotsDict["timeslots"]; !ok {
		// 	timeslotsDict["timeslots"] = make(map[uint]string)
		// }

		timeslotsDict[timeslot.ID] = timeslot.Description
	}

	// add event title to timeslotsDict
	// timeslotsDict["title"] = event.Title

	c.IndentedJSON(http.StatusOK, gin.H{"title": event.Title, "detail": event.Detail, "timeslots": timeslotsDict})
}

func (h *EventHandler) DeleteTimeslotsHandler(c *gin.Context) {
	eventID := c.Param("eventID")

	// request body
	var req DeleteTimeslotsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check cookie for host token
	tokenString, err := c.Cookie(eventID)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "permission denied, you are not the host of the event"})
		return
	}

	// get event info
	event, err := h.Repo.GetEventByID(eventID)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Event Not Found."})
		return
	}

	// check if the user is the host of the event

	if tokenString != event.HostToken {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "permission denied, you are not the host of the event"})
		return
	}

	// delete specified events
	if err := h.Repo.DeleteEventTimeslots(eventID, req.TimeslotIDs); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error deleting events."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully deleted some timeslots"})
}

func filterNotEmpty(slice []string) []string {
	var filtered []string
	for _, s := range slice {
		if s != "" {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

func (h *EventHandler) AddTimeslotsHandler(c *gin.Context) {
	eventID := c.Param("eventID")
	var eventRequest EventRequest
	if err := c.ShouldBindJSON(&eventRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Split proposals by "\n"
	proposals := strings.Split(eventRequest.DateTimeProposal, "\n")
	// remove empty string
	proposals = filterNotEmpty(proposals)

	// check lenght of proposals
	if len(proposals) > 0 {
		err := h.Repo.InsertEventTimeslot(eventID, proposals)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"AddTimeslotsHandler error": "Failed to add timeslots"})
			return
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"event_id": eventID})
}

func (h *EventHandler) CheckEventExistsHandler(c *gin.Context) {
	eventID := c.Param("eventID")
	// get event info
	if _, err := h.Repo.GetEventByID(eventID); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Event Not Found."})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Event Found."})
}

func (h *EventHandler) IsCreatedBySelfHandler(c *gin.Context) {
	eventID := c.Param("eventID")
	// check cookie for host token
	tokenString, err := c.Cookie(eventID)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "permission denied, you are not the host of the event"})
		return
	}
	// get event info
	event, err := h.Repo.GetEventByID(eventID)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Event Not Found."})
		return
	}

	// check if the user is the host of the event

	if tokenString != event.HostToken {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "permission denied, you are not the host of the event"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "you ARE the host of the event"})
}

func (h *EventHandler) AddAttendanceHandler(c *gin.Context) {
	eventID := c.Param("eventID")
	// get request body
	var req EventUserTimeslotRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	// get event info
	event, event_err := h.Repo.GetEventByID(eventID)
	if event_err != nil {
		log.Println(event_err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Event Not Found."})
		return
	}

	// convert availability map's key to integer
	uintAvailability := make(map[uint](uint))
	for strKey, value := range req.Availability {
		uintKey, err := strconv.ParseUint(strKey, 10, 64)
		if err != nil {
			// Handle the error if the conversion fails
			fmt.Printf("Error converting key %s: %v\n", strKey, err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error converting availability"})
			return
		}
		uintAvailability[uint(uintKey)] = value
	}

	err := h.Repo.AddAttendance(eventID, uintAvailability, req.Name, req.Comment, req.Email)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error storing preferences"})
		return
	}

	email_err := service.SendEmailAdd(req.Name, req.Email, eventID, event.Title, event.DueEdit)
	if email_err != nil {
		log.Println("email error")
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Successfully stored preferences"})
}

func createEventForm(users []repository.EventUser, timeslots []repository.EventTimeslot, userAvailability map[uint]map[uint]uint, maxList, minList, optimalList []uint) EventForm {
	//Create ScheduleList from EventTimeslot
	scheduleList := make([]Schedule, len(timeslots))
	annotationMap := make(map[uint](uint))
	for _, item := range maxList {
		annotationMap[item] = 1
	}
	for _, item := range minList {
		annotationMap[item] = 2
	}
	for _, item := range optimalList {
		annotationMap[item] = 3
	}
	timeslotIDToIndex := make(map[uint]int)
	for i, timeslot := range timeslots {
		annotation, ok := annotationMap[timeslot.ID]
		if !ok {
			annotation = 0
		}
		// Iterate over the slice and check if the element exists
		scheduleList[i] = Schedule{
			Name:       timeslot.Description,
			ID:         timeslot.ID,
			Annotation: annotation, // Set this accordingly
		}
		timeslotIDToIndex[timeslot.ID] = i
	}

	// //edit annotation
	// for _, timeslot := range maxList {
	// 	scheduleList[timeslot].Annotation = 1
	// }
	// for _, timeslot := range minList {
	// 	scheduleList[timeslot].Annotation = 2
	// }
	// for _, timeslot := range optimalList {
	// 	scheduleList[timeslot].Annotation = 3
	// }

	// Create patricipants

	participants := make([]Participant, len(users))
	for i, user := range users {
		result := make([]uint, len(timeslots))
		if availability, ok := userAvailability[user.ID]; ok {
			for timeslotID, preference := range availability {
				if index, exists := timeslotIDToIndex[timeslotID]; exists {
					result[index] = preference
				}
			}
		}

		participants[i] = Participant{
			Name:    user.UserName,
			Comment: user.Comment,
			UserID:  user.ID,
			Result:  result,
			Email:   user.Email,
		}
	}

	return EventForm{
		ScheduleList: scheduleList,
		Participants: participants,
	}
}

func (h *EventHandler) GetAttendanceHandler(c *gin.Context) {
	eventID := c.Param("eventID")

	event, event_err := h.Repo.GetEventByID(eventID)
	if event_err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error obtaining event"})
		return
	}

	// get list of users
	eventUsers, users_err := h.Repo.GetUsersByEventID(eventID)
	if users_err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error obtaining list of users"})
		return
	}
	// get list of timeslots
	eventTimeslots, ts_err := h.Repo.GetTimeslotsByEventID(eventID)
	if ts_err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error obtaining list of timeslots"})
		return
	}

	preferences, err := h.Repo.GetAllPreferences(eventID, eventUsers, eventTimeslots)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error obtaining preferences"})
		return
	}

	// m

	// find optimal ones
	maxList, minList, optimalList := service.FindOptimals(preferences, len(eventUsers), eventTimeslots)

	eventForm := createEventForm(eventUsers, eventTimeslots, preferences, maxList, minList, optimalList)

	// c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully got all preferences",
	// 	"userAvailability": preferences,
	// 	"users":            eventUsers,
	// 	"timeslots":        eventTimeslots})
	c.IndentedJSON(http.StatusOK, gin.H{"scheduleList": eventForm.ScheduleList, "participants": eventForm.Participants, "due_edit": event.DueEdit})
}

func (h *EventHandler) GetEventBasicHandler(c *gin.Context) {
	eventID := c.Param("eventID")

	event, err := h.Repo.GetEventByID(eventID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error obtaining event"})
		return
	}
	users, err := h.Repo.GetUsersByEventID(eventID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error obtaining users"})
		return
	}
	num_users := len(users)
	c.IndentedJSON(http.StatusOK, gin.H{"title": event.Title, "detail": event.Detail, "num_users": num_users})

}

func (h *EventHandler) ModifyAttendanceHandler(c *gin.Context) {
	eventID := c.Param("eventID")
	// get request body
	var req ModifyEventUserTimeslotRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	// get event info
	if _, err := h.Repo.GetEventByID(eventID); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Event Not Found."})
		return
	}

	// the following checking part and modifying part could be integrated (will be more efficient), but for now we do it separately
	// check if user exists
	if _, err := h.Repo.CheckIfUserExists(eventID, req.UserID); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User Not Found."})
		return
	}

	// modify user info
	if err := h.Repo.ModifyEventUser(eventID, req.UserID, req.Name, req.Comment); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error updating user information"})
		return
	}

	// convert availability map's key to integer
	uintAvailability := make(map[uint](uint))
	for strKey, value := range req.Availability {
		uintKey, err := strconv.ParseUint(strKey, 10, 64)
		if err != nil {
			// Handle the error if the conversion fails
			fmt.Printf("Error converting key %s: %v\n", strKey, err)
			continue
		}
		uintAvailability[uint(uintKey)] = value
	}

	err := h.Repo.ModifyAttendance(eventID, uintAvailability, req.Name, req.Comment, req.UserID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error updating preferences"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully modified preferences"})

}
