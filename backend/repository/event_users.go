package repository

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

func (repo Repository) DeleteEventUserByEventID(event_id string) error {
	result := repo.db.Where("event_id = ?", event_id).Delete(&EventUser{})
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

func (repo Repository) InsertEventUser(newEventUser *EventUser) (uint, error) {
	// returns newly created user_id
	result := repo.db.Create(newEventUser)
	if err := result.Error; err != nil {
		log.Println("Gorm Error:", err)
		return 0, err
	}
	log.Println(newEventUser.ID)
	fmt.Println(newEventUser.ID)
	return newEventUser.ID, nil
}

func (repo Repository) GetUsersByEventID(eventID string) ([]EventUser, error) {
	var eventUsers []EventUser
	query := repo.db.Where("event_id = ?", eventID).Order("id").Find(&eventUsers)

	if query.Error != nil {
		if query.Error == gorm.ErrRecordNotFound {
			return eventUsers, fmt.Errorf("GetTimeslotsByEventID %s: error getting timeslots", eventID)
		}
		return eventUsers, query.Error
	}
	return eventUsers, nil
}

func (repo Repository) CheckIfUserExists(eventID string, userID uint) (bool, error) {
	var eventUser EventUser
	result := repo.db.Where("id = ? AND event_id = ?", userID, eventID).First(&eventUser)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("Record not found")
			return false, nil
		} else {
			fmt.Println("Error occurred:", result.Error)
			return false, result.Error
		}
	}
	fmt.Println("Record found:", eventUser)
	return true, nil
}

func (repo Repository) ModifyEventUser(eventID string, userID uint, name string, comment string) error {
	condition := EventUser{
		ID:      userID,
		EventID: eventID,
	}

	// Use the FirstOrCreate method
	if err := repo.db.
		Where(condition).
		Assign(EventUser{UserName: name, Comment: comment}).
		FirstOrCreate(&EventUser{}).
		Error; err != nil {
		// Handle the error, if any
		fmt.Println("Error:", err)
		return err
	}
	// Update or create successful
	fmt.Println("Preference updated or row created")
	return nil
}
