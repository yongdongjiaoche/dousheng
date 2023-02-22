package dao

import (
	"douyin/global"
	"douyin/models"
)

// InsertUser 插入新用户
func InsertMessage(message *models.Message) error {
	//执行sql语句入库
	if err := global.MysqlEngine.Create(message).Error; err != nil {
		return err
	}
	return nil
}

// GetMessageList 从数据库查询对方给当前用户发送的聊天信息,从uidB到uid
func GetMessageList(uid int64, uidB int64) ([]models.Message, error) {
	var messageList []models.Message
	err := global.MysqlEngine.Where("to_user_id=? and from_user_id=?", uid, uidB).Find(&messageList).Error
	return messageList, err
}
