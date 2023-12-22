package tables

import "time"

type Updates struct {
	ID             int       `gorm:"column:id;primaryKey;autoIncrement"`
	SponsorshipsID uint      `gorm:"column:sponsorships_id"`
	Status         string    `gorm:"column:status;not null;type:varchar(50)"`
	Email          string    `gorm:"column:email;unique;type:varchar(100)"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}
