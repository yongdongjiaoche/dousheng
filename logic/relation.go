package logic

import (
	"douyin/dao"
	"douyin/global"
	"douyin/models"
	"fmt"

	"go.uber.org/zap"
)

// GetFreindList 获取用户好友列表
func GetFriendList(p *models.ParamInfo) ([]models.FriendUser, error) {
	var friendList []models.FriendUser
	followList, err := GetFollowListLogic(p.Uid)
	if err != nil {
		global.Logger.Error("获取关注用户失败", zap.Error(err))
		return friendList, err
	}
	friendList_raw, err := GetFriendListLogic(p.Uid, followList)
	if err != nil {
		global.Logger.Error("获取好友列表失败", zap.Error(err))
		return friendList, err
	}
	for i := range friendList_raw {
		friend := friendList_raw[i]
		friend_new := models.FriendUser{}
		friend_new.Id = friend.Id
		friend_new.Name = friend.Name
		friend_new.PassWord = friend.PassWord
		friend_new.FollowCount = friend.FollowCount
		friend_new.FollowerCount = friend.FollowerCount
		friend_new.IsFollow = friend.IsFollow
		friend_new.Salt = friend.Salt
		friend_new.Token = friend.Token
		msg, _ := dao.GetMessageList(p.Uid, friend.Id)
		content := "我们已经是好友啦，快开始聊天吧！"
		if len(msg) > 0 {
			content = msg[len(msg)-1].Content
		}
		friend_new.Message = content
		friend_new.MsgType = 0
		friendList = append(friendList, friend_new)
	}
	return friendList, nil
}

// GetFollowerListLogic logic层获取用户好友列表
func GetFriendListLogic(uid int64, users []models.User) ([]models.User, error) {
	friendList := []models.User{}
	for i := range users {
		follow_id := users[i].Id
		boolRes, err := dao.IsFollower(follow_id, uid)
		if err != nil {
			global.Logger.Error("Error in Getting FriendList", zap.Error(err))
			return friendList, err
		}
		if boolRes {
			friend, err := dao.GetUserByID(follow_id)
			if err != nil {
				global.Logger.Error("Error in Getting User", zap.Error(err))
				return friendList, err
			}
			friendList = append(friendList, *friend)
		}
	}
	return friendList, nil
}

// GetFollowListLogic logic层获取关注用户列表
func GetFollowListLogic(uid int64) ([]models.User, error) {
	return dao.GetFollowList(uid)
}

// GetFollowerListLogic logic层获取粉丝列表
func GetFollowerListLogic(uid int64) ([]models.User, error) {
	return dao.GetFollowerList(uid)
}

func RelationActionLogic(uid int64, toUid int64, actionType int8) error {
	// 判断关注与被关注用户是否存在
	err1 := dao.CheckUserExistById(uid)
	if err1 != nil {
		global.Logger.Error("Error in setting relation action", zap.Error(err1))
		return err1
	}
	err2 := dao.CheckUserExistById(toUid)
	if err1 != nil {
		global.Logger.Error("Error in setting relation action", zap.Error(err2))
		return err2
	}

	fmt.Println("uid:", uid, ",toUid:", toUid)
	// 执行关注/取关操作
	switch actionType {
	case 1:
		if err := dao.DoActionFollow(uid, toUid); err != nil {
			return err
		}
	case 2:
		if err := dao.DoActionUnfollow(uid, toUid); err != nil {
			return err
		}
	default:
		global.Logger.Error("Error: invalid action type")
	}
	return nil
}
