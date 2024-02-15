package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/JMjirapat/qrthrough-api/infrastructure"
	"github.com/JMjirapat/qrthrough-api/internal/core/domain"
	"github.com/JMjirapat/qrthrough-api/pkg/errors"
	"github.com/JMjirapat/qrthrough-api/pkg/rest"
	"github.com/JMjirapat/qrthrough-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

type lineHandler struct {
	serv domain.LineService
}

func NewLineHandler(serv domain.LineService) *lineHandler {
	return &lineHandler{
		serv: serv,
	}
}

func (h lineHandler) Webhook(c *fiber.Ctx) error {
	httpRequest := new(http.Request)
	if err := fasthttpadaptor.ConvertRequest(c.Context(), httpRequest, true); err != nil {
		log.Panic(err)
		return rest.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	events, err := infrastructure.LineBot.ParseRequest(httpRequest)
	if err != nil {
		log.Panic(err)
		return rest.ResponseBadRequest(c)
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if message.Text == "Get QR" {
					qrString := "https://api.qrserver.com/v1/create-qr-code/?size=320x320&margin=10&data=" + "QRTHROUGH:" + message.ID
					// Handle text messages
					messageId, err := strconv.ParseInt(message.ID, 10, 64)
					if err != nil {
						if _, err = infrastructure.LineBot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(errors.LW_PARSE_INT_FAILED_DESC)).Do(); err != nil {
							log.Panic(err)
						}
						//ถ้าถึงจุดนี้ให้หยุดและไป process event อันใหม่
						continue
					}

					if err = h.serv.CreateQR(messageId, event.Source.UserID); err != nil {
						if _, err = infrastructure.LineBot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(err.Error())).Do(); err != nil {
							log.Panic(err)
						}
						//ถ้าถึงจุดนี้ให้หยุดและไป process event อันใหม่
						continue
					}

					var messages []linebot.SendingMessage
					qrCode := linebot.NewImageMessage(qrString, qrString)
					suggest := linebot.NewTextMessage(utils.QRCODE_RESPONSE_TEXT)
					messages = append(messages, qrCode, suggest)
					if _, err = infrastructure.LineBot.ReplyMessage(event.ReplyToken, messages...).Do(); err != nil {
						log.Panic(err)
					}
				}
			}
		}
	}
	return rest.ResponseOK(c, nil)
}

// type LINEGetQRCode struct {
// 	UserID string
// }

// func (h lineHandler) ManualQR(c *fiber.Ctx) error {
// 	var body LINEGetQRCode
// 	if err := c.BodyParser(&body); err != nil {
// 		return rest.ResponseUnprocessableEntity(c)
// 	}
// 	qrcode := rand.Int63n(9999)
// 	if err := h.serv.CreateQR(qrcode, body.UserID); err != nil {
// 		log.Panicf("%v", err)
// 	}
// 	return rest.ResponseOK(c, fmt.Sprintf("QRTHROUGH:%v", qrcode))
// }
