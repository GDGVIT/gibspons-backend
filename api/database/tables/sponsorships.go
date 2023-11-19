package tables

type Sponsorship struct {
	ID          int    `gorm:"column:user_id;primaryKey;autoIncrement"`
	POCName     string `gorm:"column:poc_name;not null;type:varchar(100)"`
	Email       string `gorm:"column:email;unique;type:varchar(100)"`
	Status      string `gorm:"column:status;not null;type:varchar(100)"`
	Description string `gorm:"column:desc;type:varchar(500)"`
}
