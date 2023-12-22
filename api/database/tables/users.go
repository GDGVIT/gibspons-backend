package tables

type Users struct {
	ID       int    `gorm:"column:user_id;primaryKey;autoIncrement"`
	Name     string `gorm:"column:user_name;not null;type:varchar(100)"`
	Email    string `gorm:"column:user_email;unique;type:varchar(100)"`
	Phone    string `gorm:"column:user_phone;not null;type:varchar(10)"`
	Password []byte `gorm:"column:user_password;not null"`
}
