package model

import "time"

type Booking struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    string    `gorm:"type:varchar(50);not null" json:"user_id"`
	TicketID  string    `gorm:"type:varchar(50);not null" json:"ticket_id"`
	Status    string    `gorm:"type:varchar(20);default:'SUCCESS'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
