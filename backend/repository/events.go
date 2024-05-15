// Basic methods to handle events table

package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

func (r Repository) GetEventByID(uuid_str string) (Event, error) {
	var event Event
	query := r.db.First(&event, "event_id = ?", uuid_str)

	if query.Error != nil {
		if query.Error == gorm.ErrRecordNotFound {
			return event, fmt.Errorf("GetEventByID %s: no such event", uuid_str)
		}
		return event, query.Error
	}
	return event, nil

	// Retrieve proposals from event_timeslot table
	// rows, err := r.db.Query("SELECT description FROM event_timeslot WHERE event_id = UUID_TO_BIN(?)", uuid)
	// if err != nil {
	// 	return event, fmt.Errorf("GetEventById %s: %v", uuid, err)
	// }
	// defer rows.Close()

	// // Iterate through rows and append each proposal to the event
	// for rows.Next() {
	// 	var proposal string
	// 	if err := rows.Scan(&proposal); err != nil {
	// 		return event, fmt.Errorf("GetEventById %s: %v", uuid, err)
	// 	}
	// 	event.Proposals = append(event.Proposals, proposal)
	// }

	// // Check for errors from iterating over rows
	// if err := rows.Err(); err != nil {
	// 	return event, fmt.Errorf("GetEventById %s: %v", uuid, err)
	// }

	// return event, nil
}

func (repo Repository) DeleteEventByID(event_id string) error {
	result := repo.db.Where("event_id = ?", event_id).Delete(&Event{})
	if result.Error != nil {
		// An error occurred during deletion
		// You can handle the error here
		fmt.Println("Error deleting record:", result.Error)
		return result.Error
	} else if result.RowsAffected == 0 {
		// No records were deleted because there were no matches for the conditions
		fmt.Println("No records matching the conditions")
	} else {
		// Deletion was successful
		fmt.Println("Record deleted successfully")
	}
	return nil
}

func (repo Repository) InsertEvent(newEvent Event) error {
	if err := repo.db.Create(&newEvent).Error; err != nil {
		log.Println("Gorm Error:", err)
		return err
	}
	return nil
}

func (repo Repository) UpdateEventTitleDetail(eventID string, title string, detail string, due string) error {
	if err := repo.db.Model(&Event{}).Where("event_id = ?", eventID).Updates(Event{Title: title, Detail: detail, DueEdit: due}).Error; err != nil {
		log.Println("Gorm Error:", err)
		return err
	}
	return nil
}

func (repo Repository) UpdateEventDue(eventID string, due string) error {
	if err := repo.db.Model(&Event{}).Where("event_id = ?", eventID).Updates(Event{DueEdit: due}).Error; err != nil {
		log.Println("Gorm Error:", err)
		return err
	}
	return nil
}

func (repo Repository) GetAllEventsDueNow() ([]Event, error) {
	currentTime := time.Now()
	utcTime := currentTime.UTC().Format("2006-01-02 15:04")
	var events []Event
	err := repo.db.Where("DATE_FORMAT(STR_TO_DATE(due_edit, '%a, %d %b %Y %H:%i:%s GMT'), '%Y-%m-%d %H:%i') = ?", utcTime).Find(&events).Error
	if err != nil {
		log.Println("Gorm Error:", err)
		return events, err
	}
	return events, nil
}
