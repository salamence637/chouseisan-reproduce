package repository

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

func (repo Repository) DeleteEventUserTimeslotByEventID(event_id string) error {
	result := repo.db.Where("event_id = ?", event_id).Delete(&EventUserTimeslot{})
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

func (repo Repository) InsertEventUserTimeslot(newEventUserTimeslot EventUserTimeslot) error {
	if err := repo.db.Create(&newEventUserTimeslot).Error; err != nil {
		log.Println("Gorm Error:", err)
		return err
	}
	return nil
}

func (repo Repository) ModifyEventUserTimeslot(eventID string, timeslot_id uint, userID uint, preference uint) error {
	condition := EventUserTimeslot{
		EventID:    eventID,
		TimeslotID: timeslot_id, // Replace with the specific timeslot_id
		UserID:     userID,      // Replace with the specific user_id
	}

	// Use the FirstOrCreate method
	if err := repo.db.
		Where(condition).
		Assign(EventUserTimeslot{Preference: preference}).
		FirstOrCreate(&EventUserTimeslot{}).
		Error; err != nil {
		// Handle the error, if any
		fmt.Println("Error:", err)
		return err
	}
	// Update or create successful
	fmt.Println("Preference updated or row created")
	return nil
}

func (repo Repository) GetPreference(eventID string, userID uint, timeslotID uint) (uint, error) {
	var eventUserTimeslot EventUserTimeslot
	query := repo.db.First(&eventUserTimeslot, "event_id = ? AND user_id = ? AND timeslot_id = ?", eventID, userID, timeslotID)

	if query.Error != nil {
		if query.Error == gorm.ErrRecordNotFound {
			fmt.Println("Record not found")
			fmt.Println(userID)
			return 0, nil
		}
		return 0, query.Error
	}
	return eventUserTimeslot.Preference, nil
}
