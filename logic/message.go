package logic

import (
	"douyin/dao"
	"douyin/models"
	"time"
)

// InsertMessage 将聊天消息插入数据库
func InsertMessage(p *models.ParamMessageChatAction) (*models.Message, error) {
	message := &models.Message{
		ToUserId:   p.ToUserId,
		FromUserId: GetUserByToken(p.Token).Id,
		Content:    p.Content,
		CreateTime: time.Now().Unix(),
	}
	//数据库操作
	if err := dao.InsertMessage(message); err != nil {
		return nil, err
	}
	return message, nil
}

func GetMessageList(p *models.ParamMessageChat) ([]models.Message, error) {
	uid := GetUserByToken(p.Token).Id
	uidB := p.ToUid
	return dao.GetMessageList(uid, uidB)
}
