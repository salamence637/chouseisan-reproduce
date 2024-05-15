package handler

import (
	"chouseisan/service"
	"log"
)

func (h *EventHandler) CheckDueDatesAndSendEmails() {
	// Get the current date and time
	log.Println("checking due dates!")
	events, err := h.Repo.GetAllEventsDueNow()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("N events:", len(events))
	for _, event := range events {
		eventUsers, event_err := h.Repo.GetUsersByEventID(event.EventID)
		if event_err != nil {
			log.Println(err)
			return
		}
		for _, eventUser := range eventUsers {
			err := service.SendEmail(eventUser.UserName, eventUser.Email, event.EventID, event.Title)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
