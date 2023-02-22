package controller

import (
	"douyin/global"
	"douyin/logic"
	"douyin/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var tempChat = map[string][]models.Message{}

type ChatResponse struct {
	Response
	MessageList []models.Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	p := new(models.ParamMessageChatAction)

	if err := c.ShouldBind(p); err != nil {
		global.Logger.Error("Error in Message Sending!")
		c.JSON(http.StatusOK, Response{StatusCode: 1})
	}
	msg, err := logic.InsertMessage(p)
	if err != nil {
		global.Logger.Error("Error in Message Sending!")
		c.JSON(http.StatusOK, Response{StatusCode: 1})
	}
	fmt.Println("msg sending:", msg)

	c.JSON(http.StatusOK, Response{StatusCode: 0})
}

var chatBox = map[int]int{}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	p := new(models.ParamMessageChat)
	if err := c.ShouldBind(p); err != nil {
		global.Logger.Error("Error in Chat!")
		c.JSON(http.StatusOK, ChatResponse{
			Response: Response{
				StatusCode: 1,
			},
			MessageList: []models.Message{},
		})
	}
	// fmt.Println("premsgtime", p.PreMsgTime)
	msgList, err := logic.GetMessageList(p)
	if len(msgList) > 0 {
		msgList = msgList[len(msgList)-1:]
		latestTime := msgList[0].CreateTime
		if _, ok := chatBox[int(latestTime)]; ok {
			msgList = msgList[:0]
		}
		chatBox[int(latestTime)] = 1
	}

	if err != nil {
		global.Logger.Error("Error in Chat!")
		c.JSON(http.StatusOK, ChatResponse{
			Response: Response{
				StatusCode: 1,
			},
			MessageList: []models.Message{},
		})
	}
	fmt.Println("msgList:", msgList)

	c.JSON(http.StatusOK, ChatResponse{
		Response: Response{
			StatusCode: 0,
		},
		MessageList: msgList,
	})
}
