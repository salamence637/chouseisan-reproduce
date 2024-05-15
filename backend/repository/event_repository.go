// Repository methods related to Event handling.
package repository

import (
	"log"

	"github.com/google/uuid"
)

// Define methods for reading data from the database

// CreateEventInDB creates a new event in the database
func (repo Repository) CreateEvent(title, detail string, due_edit string, proposals []string) (string, string, error) {
	// issue new event id and hostToken (both are uuid)
	newEventID := uuid.New().String()
	hostToken := uuid.New().String()

	// Create new event instance
	newEvent := Event{
		EventID:   newEventID,
		Title:     title,
		Detail:    detail,
		DueEdit:   due_edit,
		HostToken: hostToken,
	}

	// Insert new event info to db
	err := repo.InsertEvent(newEvent)
	if err != nil {
		log.Println("Gorm Error:", err)
		return "", "", err
	}

	// Insert each proposal into event_timeslot
	err = repo.InsertEventTimeslot(newEventID, proposals)
	if err != nil {
		log.Println("Gorm Error:", err)
		return "", "", err
	}

	return newEventID, hostToken, nil
}

func (repo Repository) DeleteEvent(event_id string) error {
	if err := repo.DeleteEventByID(event_id); err != nil {
		log.Println("error deleting event from events table:", err)
		return err
	}
	if err := repo.DeleteEventUserByEventID(event_id); err != nil {
		log.Println("error deleting event from event_users table:", err)
		return err
	}
	if err := repo.DeleteEventTimeslotByEventID(event_id); err != nil {
		log.Println("error deleting event from event_timeslots table:", err)
		return err
	}
	if err := repo.DeleteEventUserTimeslotByEventID(event_id); err != nil {
		log.Println("error deleting event from event_user_timeslots table:", err)
		return err
	}
	return nil
}

func (repo Repository) AddAttendance(eventID string, availability map[uint](uint), name string, comment string, email string) error {
	// first add user
	newUser := EventUser{EventID: eventID, UserName: name, Email: email, Comment: comment}
	newUserID, err := repo.InsertEventUser(&newUser)
	if err != nil {
		log.Println("error creating new user:", err)
		return err
	}
	// For each timeslot, add record to event_user_timeslots
	for timeslot_id, preference := range availability {
		// first check if timeslot exists
		exist, err := repo.CheckTimeslotsExist(eventID, timeslot_id)
		if err != nil {
			log.Println("error checking timeslot existence")
			return err
		}
		if exist {
			newEventUserTimeslot := EventUserTimeslot{EventID: eventID, UserID: newUserID, TimeslotID: timeslot_id, Preference: preference}
			if err := repo.InsertEventUserTimeslot(newEventUserTimeslot); err != nil {
				log.Println("error inserting availability")
				return err
			}
		} //else: do nothing
	}
	return nil
}

func (repo Repository) ModifyAttendance(eventID string, availability map[uint](uint), name string, comment string, userID uint) error {
	// For each timeslot, add record to event_user_timeslots
	for timeslot_id, preference := range availability {
		// first check if timeslot exists
		exist, err := repo.CheckTimeslotsExist(eventID, timeslot_id)
		if err != nil {
			log.Println("error checking timeslot existence")
			return err
		}
		if exist {
			// Specify the condition to identify the row
			if err := repo.ModifyEventUserTimeslot(eventID, timeslot_id, userID, preference); err != nil {
				log.Println("error modifying attendance", err)
				return err
			}
		} //else: do nothing
	}
	return nil
}

func (repo Repository) GetAllPreferences(eventID string, eventUsers []EventUser, eventTimeslots []EventTimeslot) (map[uint](map[uint](uint)), error) {
	userAvailability := make(map[uint](map[uint](uint)))
	for _, user := range eventUsers {
		userMap := make(map[uint](uint))
		for _, timeslot := range eventTimeslots {
			preference, err := repo.GetPreference(eventID, user.ID, timeslot.ID)
			if err != nil {
				log.Println("error obtaining preference")
				return userAvailability, err
			}
			userMap[timeslot.ID] = preference
		}
		userAvailability[user.ID] = userMap
	}
	return userAvailability, nil
}

// CreateEventInDB creates a new event in the database
// func (repo Repository) CreateEvent(title, detail string, proposals []string) (string, error) {
// 	// Start a database transaction
// 	tx, err := repo.db.Begin()
// 	if err != nil {
// 		log.Println("Error starting transaction:", err)
// 		return "", err
// 	}

// 	newUUID := uuid.New().String()
// 	hostToken := uuid.New().String()

// 	fmt.Println(newUUID)

// 	// Insert event information into the events table
// 	_, err = tx.Exec("INSERT INTO events (event_id, title, detail, host_token) VALUES (?,?,?,?);", newUUID, title, detail, hostToken)
// 	if err != nil {
// 		// Rollback the transaction in case of an error
// 		tx.Rollback()
// 		log.Println("Error inserting event:", err)
// 		return "", err
// 	}

// 	// Insert timeslot information into the event-timeslot table
// 	stmt, err := tx.Prepare("INSERT INTO event_timeslots (event_id, description) VALUES (?, ?)")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()

// 	// Insert each proposal into event_timeslot
// 	for proposal := range proposals {
// 		_, err := stmt.Exec(newUUID, proposal)
// 		fmt.Println(proposal)
// 		if err != nil {
// 			// Rollback the transaction in case of an error
// 			tx.Rollback()
// 			log.Fatal(err)
// 		}
// 	}

// 	// Commit the transaction
// 	if err := tx.Commit(); err != nil {
// 		log.Println("Error committing transaction:", err)
// 		return "", err
// 	}

// 	return newUUID, nil
// }
