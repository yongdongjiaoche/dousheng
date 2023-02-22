package models

type Message struct {
	Id         int64  `json:"id,omitempty" gorm:"column:message_id; primaryKey"`
	ToUserId   int64  `json:"to_user_id"`
	FromUserId int64  `json:"from_user_id"`
	Content    string `json:"content,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
}
