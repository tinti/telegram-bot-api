package tgbotapi

import (
	"fmt"
	"log"
	"net/url"
	"math/rand"
	"encoding/json"

	"github.com/streadway/amqp"
)

type BotAPIIface interface {
	MakeRequest(endpoint string, params url.Values) (APIResponse, error)
	UploadFile(endpoint string, params map[string]string, fieldname string, file interface{}) (APIResponse, error)
	GetFileDirectURL(fileID string) (string, error)
	GetMe() (User, error)
	IsMessageToMe(message Message) bool
	Send(c Chattable) (Message, error)
	GetUserProfilePhotos(config UserProfilePhotosConfig) (UserProfilePhotos, error)
	GetFile(config FileConfig) (File, error)
	GetUpdates(config UpdateConfig) ([]Update, error)
	RemoveWebhook() (APIResponse, error)
	SetWebhook(config WebhookConfig) (APIResponse, error)
	GetWebhookInfo() (WebhookInfo, error)
	GetUpdatesChan(config UpdateConfig) (UpdatesChannel, error)
	ListenForWebhook(pattern string) UpdatesChannel
	AnswerInlineQuery(config InlineConfig) (APIResponse, error)
	AnswerCallbackQuery(config CallbackConfig) (APIResponse, error)
	KickChatMember(config KickChatMemberConfig) (APIResponse, error)
	LeaveChat(config ChatConfig) (APIResponse, error)
	GetChat(config ChatConfig) (Chat, error)
	GetChatAdministrators(config ChatConfig) ([]ChatMember, error)
	GetChatMembersCount(config ChatConfig) (int, error)
	GetChatMember(config ChatConfigWithUser) (ChatMember, error)
	UnbanChatMember(config ChatMemberConfig) (APIResponse, error)
	RestrictChatMember(config RestrictChatMemberConfig) (APIResponse, error)
	PromoteChatMember(config PromoteChatMemberConfig) (APIResponse, error)
	GetGameHighScores(config GetGameHighScoresConfig) ([]GameHighScore, error)
	AnswerShippingQuery(config ShippingConfig) (APIResponse, error)
	AnswerPreCheckoutQuery(config PreCheckoutConfig) (APIResponse, error)
	DeleteMessage(config DeleteMessageConfig) (APIResponse, error)
	GetInviteLink(config ChatConfig) (string, error)
	PinChatMessage(config PinChatMessageConfig) (APIResponse, error)
	UnpinChatMessage(config UnpinChatMessageConfig) (APIResponse, error)
	SetChatTitle(config SetChatTitleConfig) (APIResponse, error)
	SetChatDescription(config SetChatDescriptionConfig) (APIResponse, error)
	SetChatPhoto(config SetChatPhotoConfig) (APIResponse, error)
	DeleteChatPhoto(config DeleteChatPhotoConfig) (APIResponse, error)
}

var _ BotAPIIface = (*BotAPI)(nil)

type EchoBotAPI struct {
}

func (bot *EchoBotAPI) MakeRequest(endpoint string, params url.Values) (APIResponse, error) {
	log.Println("MakeRequest")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) UploadFile(endpoint string, params map[string]string, fieldname string, file interface{}) (APIResponse, error) {
	log.Println("UploadFile")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) GetFileDirectURL(fileID string) (string, error) {
	log.Println("GetFileDirectURL")
	return "", nil
}

func (bot *EchoBotAPI) GetMe() (User, error) {
	log.Println("GetMe")
	return User{}, nil
}

func (bot *EchoBotAPI) IsMessageToMe(message Message) bool {
	log.Println("IsMessageToMe")
	return false
}

func (bot *EchoBotAPI) Send(c Chattable) (Message, error) {
	log.Println("Send")
	return Message{}, nil
}

func (bot *EchoBotAPI) GetUserProfilePhotos(config UserProfilePhotosConfig) (UserProfilePhotos, error) {
	log.Println("GetUserProfilePhotos")
	return UserProfilePhotos{}, nil
}

func (bot *EchoBotAPI) GetFile(config FileConfig) (File, error) {
	log.Println("GetFile")
	return File{}, nil
}

func (bot *EchoBotAPI) GetUpdates(config UpdateConfig) ([]Update, error) {
	log.Println("GetUpdates")
	return []Update{}, nil
}

func (bot *EchoBotAPI) RemoveWebhook() (APIResponse, error) {
	log.Println("RemoveWebhook")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) SetWebhook(config WebhookConfig) (APIResponse, error) {
	log.Println("SetWebhook")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) GetWebhookInfo() (WebhookInfo, error) {
	log.Println("GetWebhookInfo")
	return WebhookInfo{}, nil
}

func (bot *EchoBotAPI) GetUpdatesChan(config UpdateConfig) (UpdatesChannel, error) {
	log.Println("GetUpdatesChan")
	ch := make(chan Update)
	return ch, nil
}

func (bot *EchoBotAPI) ListenForWebhook(pattern string) UpdatesChannel {
	log.Println("ListenForWebhook")
	ch := make(chan Update)
	return ch
}

func (bot *EchoBotAPI) AnswerInlineQuery(config InlineConfig) (APIResponse, error) {
	log.Println("AnswerInlineQuery")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) AnswerCallbackQuery(config CallbackConfig) (APIResponse, error) {
	log.Println("AnswerCallbackQuery")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) KickChatMember(config KickChatMemberConfig) (APIResponse, error) {
	log.Println("KickChatMember")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) LeaveChat(config ChatConfig) (APIResponse, error) {
	log.Println("LeaveChat")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) GetChat(config ChatConfig) (Chat, error) {
	log.Println("GetChat")
	return Chat{}, nil
}

func (bot *EchoBotAPI) GetChatAdministrators(config ChatConfig) ([]ChatMember, error) {
	log.Println("GetChatAdministrators")
	return []ChatMember{}, nil
}

func (bot *EchoBotAPI) GetChatMembersCount(config ChatConfig) (int, error) {
	log.Println("GetChatMembersCount")
	return 0, nil
}

func (bot *EchoBotAPI) GetChatMember(config ChatConfigWithUser) (ChatMember, error) {
	log.Println("GetChatMember")
	return ChatMember{}, nil
}

func (bot *EchoBotAPI) UnbanChatMember(config ChatMemberConfig) (APIResponse, error) {
	log.Println("UnbanChatMember")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) RestrictChatMember(config RestrictChatMemberConfig) (APIResponse, error) {
	log.Println("RestrictChatMember")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) PromoteChatMember(config PromoteChatMemberConfig) (APIResponse, error) {
	log.Println("PromoteChatMember")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) GetGameHighScores(config GetGameHighScoresConfig) ([]GameHighScore, error) {
	log.Println("GetGameHighScores")
	return []GameHighScore{}, nil
}

func (bot *EchoBotAPI) AnswerShippingQuery(config ShippingConfig) (APIResponse, error) {
	log.Println("AnswerShippingQuery")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) AnswerPreCheckoutQuery(config PreCheckoutConfig) (APIResponse, error) {
	log.Println("AnswerPreCheckoutQuery")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) DeleteMessage(config DeleteMessageConfig) (APIResponse, error) {
	log.Println("DeleteMessage")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) GetInviteLink(config ChatConfig) (string, error) {
	log.Println("GetInviteLink")
	return "", nil
}

func (bot *EchoBotAPI) PinChatMessage(config PinChatMessageConfig) (APIResponse, error) {
	log.Println("PinChatMessage")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) UnpinChatMessage(config UnpinChatMessageConfig) (APIResponse, error) {
	log.Println("UnpinChatMessage")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) SetChatTitle(config SetChatTitleConfig) (APIResponse, error) {
	log.Println("SetChatTitle")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) SetChatDescription(config SetChatDescriptionConfig) (APIResponse, error) {
	log.Println("SetChatDescription")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) SetChatPhoto(config SetChatPhotoConfig) (APIResponse, error) {
	log.Println("SetChatPhoto")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) DeleteChatPhoto(config DeleteChatPhotoConfig) (APIResponse, error) {
	log.Println("DeleteChatPhoto")
	return APIResponse{}, nil
}

var _ BotAPIIface = (*EchoBotAPI)(nil)

type RemoteBotAPI struct {
}

type RequestMessage struct {
	Operation string
	CorrelationId string

	Endpoint string
	Params   url.Values
}

type ResponseMessage struct {
	Operation string
	CorrelationId string

	ApiResponse APIResponse
	Err         error
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func (bot *RemoteBotAPI) MakeRequest(endpoint string, params url.Values) (APIResponse, error) {
	log.Println("-> MakeRequest")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/guest")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	corrId := randomString(32)

	request, _ := json.Marshal(RequestMessage{"MakeRequest", corrId, endpoint, params})

	err = ch.Publish(
		"",          // exchange
		"rpc_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if corrId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.ApiResponse, response.Err
}

func (bot *RemoteBotAPI) UploadFile(endpoint string, params map[string]string, fieldname string, file interface{}) (APIResponse, error) {
	log.Println("UploadFile")
	return APIResponse{}, nil
}

func (bot *RemoteBotAPI) GetFileDirectURL(fileID string) (string, error) {
	log.Println("GetFileDirectURL")
	return "", nil
}

func (bot *RemoteBotAPI) GetMe() (User, error) {
	log.Println("GetMe")
	return User{}, nil
}

func (bot *RemoteBotAPI) IsMessageToMe(message Message) bool {
	log.Println("IsMessageToMe")
	return false
}

func (bot *RemoteBotAPI) Send(c Chattable) (Message, error) {
	log.Println("Send")
	return Message{}, nil
}

func (bot *RemoteBotAPI) GetUserProfilePhotos(config UserProfilePhotosConfig) (UserProfilePhotos, error) {
	log.Println("GetUserProfilePhotos")
	return UserProfilePhotos{}, nil
}

func (bot *RemoteBotAPI) GetFile(config FileConfig) (File, error) {
	log.Println("GetFile")
	return File{}, nil
}

func (bot *RemoteBotAPI) GetUpdates(config UpdateConfig) ([]Update, error) {
	log.Println("GetUpdates")
	return []Update{}, nil
}

func (bot *RemoteBotAPI) RemoveWebhook() (APIResponse, error) {
	log.Println("RemoveWebhook")
	return APIResponse{}, nil
}

func (bot *RemoteBotAPI) SetWebhook(config WebhookConfig) (APIResponse, error) {
	log.Println("SetWebhook")
	return APIResponse{}, nil
}

func (bot *RemoteBotAPI) GetWebhookInfo() (WebhookInfo, error) {
	log.Println("GetWebhookInfo")
	return WebhookInfo{}, nil
}

func (bot *RemoteBotAPI) GetUpdatesChan(config UpdateConfig) (UpdatesChannel, error) {
	log.Println("GetUpdatesChan")
	ch := make(chan Update)
	return ch, nil
}

func (bot *RemoteBotAPI) ListenForWebhook(pattern string) UpdatesChannel {
	log.Println("ListenForWebhook")
	ch := make(chan Update)
	return ch
}

func (bot *RemoteBotAPI) AnswerInlineQuery(config InlineConfig) (APIResponse, error) {
	log.Println("AnswerInlineQuery")
	return APIResponse{}, nil
}

func (bot *RemoteBotAPI) AnswerCallbackQuery(config CallbackConfig) (APIResponse, error) {
	log.Println("AnswerCallbackQuery")
	return APIResponse{}, nil
}

func (bot *RemoteBotAPI) KickChatMember(config KickChatMemberConfig) (APIResponse, error) {
	log.Println("KickChatMember")
	return APIResponse{}, nil
}

func (bot *RemoteBotAPI) LeaveChat(config ChatConfig) (APIResponse, error) {
	log.Println("LeaveChat")
	return APIResponse{}, nil
}

func (bot *RemoteBotAPI) GetChat(config ChatConfig) (Chat, error) {
	log.Println("GetChat")
	return Chat{}, nil
}

func (bot *RemoteBotAPI) GetChatAdministrators(config ChatConfig) ([]ChatMember, error) {
	log.Println("GetChatAdministrators")
	return []ChatMember{}, nil
}

func (bot *RemoteBotAPI) GetChatMembersCount(config ChatConfig) (int, error) {
	log.Println("GetChatMembersCount")
	return 0, nil
}

func (bot *RemoteBotAPI) GetChatMember(config ChatConfigWithUser) (ChatMember, error) {
	log.Println("GetChatMember")
	return ChatMember{}, nil
}

func (bot *RemoteBotAPI) UnbanChatMember(config ChatMemberConfig) (APIResponse, error) {
	log.Println("UnbanChatMember")
	return APIResponse{}, nil
}

func (bot *RemoteBotAPI) RestrictChatMember(config RestrictChatMemberConfig) (APIResponse, error) {
	log.Println("RestrictChatMember")
	return APIResponse{}, nil
}

func (bot *RemoteBotAPI) PromoteChatMember(config PromoteChatMemberConfig) (APIResponse, error) {
	log.Println("PromoteChatMember")
	return APIResponse{}, nil
}

func (bot *RemoteBotAPI) GetGameHighScores(config GetGameHighScoresConfig) ([]GameHighScore, error) {
	log.Println("GetGameHighScores")
	return []GameHighScore{}, nil
}

func (bot *RemoteBotAPI) AnswerShippingQuery(config ShippingConfig) (APIResponse, error) {
	log.Println("AnswerShippingQuery")
	return APIResponse{}, nil
}

func (bot *RemoteBotAPI) AnswerPreCheckoutQuery(config PreCheckoutConfig) (APIResponse, error) {
	log.Println("AnswerPreCheckoutQuery")
	return APIResponse{}, nil
}

func (bot *RemoteBotAPI) DeleteMessage(config DeleteMessageConfig) (APIResponse, error) {
	log.Println("DeleteMessage")
	return APIResponse{}, nil
}

func (bot *RemoteBotAPI) GetInviteLink(config ChatConfig) (string, error) {
	log.Println("GetInviteLink")
	return "", nil
}

func (bot *RemoteBotAPI) PinChatMessage(config PinChatMessageConfig) (APIResponse, error) {
	log.Println("PinChatMessage")
	return APIResponse{}, nil
}

func (bot *RemoteBotAPI) UnpinChatMessage(config UnpinChatMessageConfig) (APIResponse, error) {
	log.Println("UnpinChatMessage")
	return APIResponse{}, nil
}

func (bot *RemoteBotAPI) SetChatTitle(config SetChatTitleConfig) (APIResponse, error) {
	log.Println("SetChatTitle")
	return APIResponse{}, nil
}

func (bot *RemoteBotAPI) SetChatDescription(config SetChatDescriptionConfig) (APIResponse, error) {
	log.Println("SetChatDescription")
	return APIResponse{}, nil
}

func (bot *RemoteBotAPI) SetChatPhoto(config SetChatPhotoConfig) (APIResponse, error) {
	log.Println("SetChatPhoto")
	return APIResponse{}, nil
}

func (bot *RemoteBotAPI) DeleteChatPhoto(config DeleteChatPhotoConfig) (APIResponse, error) {
	log.Println("DeleteChatPhoto")
	return APIResponse{}, nil
}

var _ BotAPIIface = (*RemoteBotAPI)(nil)

func DummyServer() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/guest")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when usused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			n := RequestMessage{}
			err := json.Unmarshal(d.Body, &n)
			log.Printf("%#v", n)
			failOnError(err, "Failed to convert body to request")

			var r ResponseMessage
			switch n.Operation {
			case "MakeRequest":
				log.Println("<- MakeRequest")
				r = ResponseMessage{}
				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			default:
				log.Printf("?")
				r = ResponseMessage{}
				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			}

			response, _ := json.Marshal(r)

			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(response),
				})
			failOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
	}()

	<-forever
}
