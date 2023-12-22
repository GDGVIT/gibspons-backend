package tables

type Updates struct {
	ID     int    `gorm:"column:id;primaryKey;autoIncrement"`
	Status string `gorm:"column:status;not null;type:varchar(50)"`
	Email  string `gorm:"column:email;unique;type:varchar(100)"`
}
