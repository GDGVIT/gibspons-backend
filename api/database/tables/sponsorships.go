package tables

type Sponsorships struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement"`
	PocID       int    `gorm:"column:poc_id;not null"`
	Email       string `gorm:"column:email;type:varchar(100)"`
	Status      string `gorm:"column:status;not null;type:varchar(100)"`
	Description string `gorm:"column:description;type:varchar(500)"`
}
