package models

type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	UserName string `json:"user_name"`
	UserUID  string `json:"user_uid"`
	BoxID    int    `json:"box_id"`
	QQ       string `json:"qq"`
}

func (User) TableName() string {
	return "users"
}
