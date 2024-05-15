package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

func (repo Repository) DeleteEventTimeslotByEventID(event_id string) error {
	result := repo.db.Where("event_id = ?", event_id).Delete(&EventTimeslot{})
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

func (repo Repository) DeleteEventTimeslots(eventID string, timeslots []uint) error {
	for _, id := range timeslots {
		if err := repo.db.Where("id = ? AND event_id = ?", id, eventID).Delete(&EventTimeslot{}).Error; err != nil {
			fmt.Println("Error deleting record:", err)
			return err
		}
	}
	return nil
}

func (repo Repository) InsertEventTimeslot(eventID string, proposals []string) error {
	for i, proposal := range proposals {
		log.Println(i)
		newEventTimeslot := EventTimeslot{
			EventID:     eventID,
			Description: proposal,
		}
		if err := repo.db.Create(&newEventTimeslot).Error; err != nil {
			log.Println("Gorm Error:", err)
			return err
		}
	}
	return nil
}

func (repo Repository) GetTimeslotsByEventID(eventID string) ([]EventTimeslot, error) {
	var eventTimeslots []EventTimeslot
	query := repo.db.Where("event_id = ?", eventID).Order("id").Find(&eventTimeslots)

	if query.Error != nil {
		if query.Error == gorm.ErrRecordNotFound {
			return eventTimeslots, fmt.Errorf("GetTimeslotsByEventID %s: error getting timeslots", eventID)
		}
		return eventTimeslots, query.Error
	}
	return eventTimeslots, nil
}

func (repo Repository) CheckTimeslotsExist(eventID string, timeslotID uint) (bool, error) {
	// type EventTimeslot struct {
	// 	TimeslotID  uint   `gorm:"primarykey;autoIncrement"`
	// 	EventID     string `gorm:"column:event_id"`
	// 	Description string `gorm:"column:description"`
	// }
	var eventTimeslot EventTimeslot

	if err := repo.db.Where("id = ?", timeslotID).First(&eventTimeslot).Error; err != nil {
		// Record not found, it does not exist
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Record does not exist")
			return false, nil
		} else {
			// Handle other errors, if any
			fmt.Println("Error:", err)
			return false, err
		}
	} else {
		// Record found, it exists
		if eventTimeslot.EventID == eventID {
			fmt.Println("Record exists:", eventTimeslot)
			return true, nil
		} else {
			fmt.Println("Record does not exist")
			return false, nil
		}
	}
}
