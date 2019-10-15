package main

import (
    "log"
    "os"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/line/line-bot-sdk-go/linebot"
)

func main() {
    port := os.Getenv("PORT")

    if port == "" {
        log.Fatal("$PORT must be set")
    }

    bot, err := linebot.New(
        os.Getenv("CHANNEL_SECRET"),
        os.Getenv("CHANNEL_TOKEN"),
    )
    if err != nil {
        log.Fatal(err)
    }

    router := gin.New()
    router.Use(gin.Logger())

    router.POST("/callback", func(c *gin.Context) {
        events, err := bot.ParseRequest(c.Request)
        if err != nil {
            if err == linebot.ErrInvalidSignature {
                log.Print(err)
            }
            return
        }
        for _, event := range events {
            if event.Type == linebot.EventTypeMessage {
		switch message := event.Message.(type) {
	            case *linebot.TextMessage:

                    // reply_msg := fmt.Sprintf("Your Token: %s\nyou type: %s\nmessage id: %s\nuser_id:%s\ngroup_id:%s\nroom_id:%s", event.ReplyToken, message.Text, message.ID, event.Source.UserID, event.Source.GroupID, event.Source.RoomID)
                    // if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply_msg)).Do(); err != nil {
                    // log.Print(err)
                    // }
                    reply_msg := ""

                    switch message.Text {
                        case "珍珍"：
                        reply_msg = fmt.Sprintf("是大美女")
                        case "娜娜"：
                        reply_msg = fmt.Sprintf("是大便")
                        case "不想上班" || "我不想上班" || "好累":
                        reply_msg = fmt.Sprintf("好，不要上班，回家好嗎")
                        case "好":
                        reply_msg = fmt.Sprintf("乖")
                        case "再見":
                        reply_msg = fmt.Sprintf("不要走")
                        case "你好嗎":
                        reply_msg = fmt.Sprintf("我還好")
                        case "你走了嗎" || "你回家了嗎":
                        reply_msg = fmt.Sprintf("還沒我要工作到死")
                        case "要吃什麼" || "晚上要吃什麼":
                        reply_msg = fmt.Sprintf("大便")
                        case "快回家":
                        reply_msg = fmt.Sprintf("不要")
                        default:
                        reply_msg = fmt.Sprintf("再見")
                    }
                    if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(reply_msg)).Do(); err != nil {
                        log.Print(err)
                    }
                }
            }
        }
    })

    router.POST("/push_message", func(c *gin.Context) {
        buf  := make([]byte, 1024)  
        n, _ := c.Request.Body.Read(buf) 
        body := string(buf[0:n])

        if _, err := bot.PushMessage("R6dcb63709978fed802b24764686c3ea8", linebot.NewTextMessage(body)).Do(); err != nil {
            log.Print(err)
        }
    })

    router.Run(":" + port)
}
