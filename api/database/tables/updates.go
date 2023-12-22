package tables

type Updates struct {
	ID     int    `gorm:"column:user_id;primaryKey;autoIncrement"`
	Status string `gorm:"column:user_name;not null;type:varchar(50)"`
	Email  string `gorm:"column:user_email;unique;type:varchar(100)"`
}
