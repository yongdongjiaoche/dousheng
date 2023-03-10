package models

import "mime/multipart"

// ParamRegister 注册参数
type ParamRegister struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=12"`
}

// ParamLogin 登录参数
type ParamLogin struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=12"`
}

// ParamInfo 获取个人信息参数
type ParamInfo struct {
	Uid   int64  `form:"user_id" json:"user_id" binding:"required"`
	Token string `form:"token" json:"token" binding:"required"`
}

// ParamMessageChat 获取聊天信息参数
type ParamMessageChat struct {
	ToUid int64  `form:"to_user_id" json:"to_user_id" binding:"required"`
	Token string `form:"token" json:"token" binding:"required"`
	// PreMsgTime int64  `form:"pre_msg_time" json:"pre_msg_time" binding:"required"`
}

// ParamMessageChatAction 获取聊天信息参数
type ParamMessageChatAction struct {
	Token      string `form:"token" json:"token" binding:"required"`
	ToUserId   int64  `form:"to_user_id" json:"to_user_id" binding:"required"`
	ActionType int32  `form:"action_type" json:"action_type" binding:"required"`
	Content    string `form:"content" json:"content" binding:"required"`
}

// ParamFavoriteAction 用户点赞请求
type ParamFavoriteAction struct {
	Token      string `form:"token" json:"token" binding:"required"`
	VideoId    int64  `form:"video_id" json:"video_id" binding:"required"`
	ActionType int8   `form:"action_type" json:"action_type" binding:"required"`
}

// ParamFavoriteList  用户获取点赞请求列表
type ParamFavoriteList struct {
	Token  string `form:"token" json:"token" binding:"required"`
	UserId int64  `form:"user_id" json:"user_id"`
}

// ParamCommentAction 用户获取评论请求列表
type ParamCommentAction struct {
	Token       string `form:"token" json:"token" binding:"required"`
	VideoId     int64  `form:"video_id" json:"video_id" binding:"required"`
	ActionType  int8   `form:"action_type" json:"action_type" binding:"required"`
	CommentText string `form:"comment_text" json:"comment_text"`
	CommentId   int64  `form:"comment_id" json:"comment_id"`
}

// ParamCommentList  用户获取评论请求列表
type ParamCommentList struct {
	Token   string `form:"token" json:"token" binding:"required"`
	VideoId int64  `form:"video_id" json:"video_id" binding:"required"`
}

// ParamRelationAction 用户关注
type ParamRelationAction struct {
	Token      string `form:"token" json:"token" binding:"required"`
	ToUserId   int64  `form:"to_user_id" json:"to_user_id" binding:"required"`
	ActionType int8   `form:"action_type" json:"action_type" binding:"required"` // 1-关注，2-取消
}

// ParamAuth 用户Token获取
type ParamAuth struct {
	Token string `form:"token" json:"token"`
}

// ParamPublishAction  用户上传视频请求
type ParamPublishAction struct {
	Token string               `form:"token" json:"token" binding:"required"`
	Data  multipart.FileHeader `form:"data" json:"data" binding:"required"`
	Title string               `form:"title" json:"title" binding:"required"`
}

// ParamPublishList  用户发布视频列表
type ParamPublishList struct {
	Token  string `form:"token" json:"token" binding:"required"`
	UserId int64  `form:"user_id" json:"user_id" binding:"required"`
}
