package repository

type Event struct {
	EventID   string `gorm:"primaryKey"`
	Title     string `json:"title" gorm:"column:title"`
	Detail    string `json:"detail" gorm:"column:detail"`
	DueEdit   string `json:"dueDate gorm:column:due_edit:"`
	HostToken string `gorm:"column:host_token"`
}

type EventUser struct {
	ID       uint   `gorm:"primaryKey"`
	EventID  string `gorm:"column:event_id"`
	UserName string `gorm:"column:user_name"`
	Email    string `gorm:"column:email"`
	Comment  string `gorm:"column:comment"`
}

type EventTimeslot struct {
	ID          uint   `gorm:"primaryKey;column:id"`
	EventID     string `gorm:"column:event_id"`
	Description string `gorm:"column:description"`
}

type EventUserTimeslot struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	EventID    string `gorm:"column:event_id"`
	UserID     uint   `gorm:"column:user_id"`
	TimeslotID uint   `gorm:"column:timeslot_id"`
	Preference uint   `gorm:"column:preference"`
}
