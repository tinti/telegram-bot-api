package tgbotapi

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"

	"github.com/streadway/amqp"
)

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
		log.Printf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

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

const (
	OperationMakeRequest            = "MakeRequest"
	OperationUploadFile             = "UploadFile"
	OperationGetFileDirectURL       = "GetFileDirectURL"
	OperationGetMe                  = "GetMe"
	OperationIsMessageToMe          = "IsMessageToMe"
	OperationSend                   = "Send"
	OperationGetUserProfilePhotos   = "GetUserProfilePhotos"
	OperationGetFile                = "GetFile"
	OperationGetUpdates             = "GetUpdates"
	OperationRemoveWebhook          = "RemoveWebhook"
	OperationSetWebhook             = "SetWebhook"
	OperationGetWebhookInfo         = "GetWebhookInfo"
	OperationGetUpdatesChan         = "GetUpdatesChan"
	OperationListenForWebhook       = "ListenForWebhook"
	OperationAnswerInlineQuery      = "AnswerInlineQuery"
	OperationAnswerCallbackQuery    = "AnswerCallbackQuery"
	OperationKickChatMember         = "KickChatMember"
	OperationLeaveChat              = "LeaveChat"
	OperationGetChat                = "GetChat"
	OperationGetChatAdministrators  = "GetChatAdministrators"
	OperationGetChatMembersCount    = "GetChatMembersCount"
	OperationGetChatMember          = "GetChatMember"
	OperationUnbanChatMember        = "UnbanChatMember"
	OperationRestrictChatMember     = "RestrictChatMember"
	OperationPromoteChatMember      = "PromoteChatMember"
	OperationGetGameHighScores      = "GetGameHighScores"
	OperationAnswerShippingQuery    = "AnswerShippingQuery"
	OperationAnswerPreCheckoutQuery = "AnswerPreCheckoutQuery"
	OperationDeleteMessage          = "DeleteMessage"
	OperationGetInviteLink          = "GetInviteLink"
	OperationPinChatMessage         = "PinChatMessage"
	OperationUnpinChatMessage       = "UnpinChatMessage"
	OperationSetChatTitle           = "SetChatTitle"
	OperationSetChatDescription     = "SetChatDescription"
	OperationSetChatPhoto           = "SetChatPhoto"
	OperationDeleteChatPhoto        = "DeleteChatPhoto"
)

type RequestMessage struct {
	Operation     string
	CorrelationId string

	C         Chattable
	Config    UserProfilePhotosConfig
	Config2   FileConfig
	Config3   UpdateConfig
	Config4   WebhookConfig
	Config5   InlineConfig
	Config6   CallbackConfig
	Config7   KickChatMemberConfig
	Config8   ChatConfig
	Config9   ChatConfigWithUser
	Config10  ChatMemberConfig
	Config11  RestrictChatMemberConfig
	Config12  PromoteChatMemberConfig
	Config13  GetGameHighScoresConfig
	Config14  ShippingConfig
	Config15  PreCheckoutConfig
	Config16  DeleteMessageConfig
	Config17  PinChatMessageConfig
	Config18  UnpinChatMessageConfig
	Config19  SetChatTitleConfig
	Config20  SetChatDescriptionConfig
	Config21  SetChatPhotoConfig
	Config22  DeleteChatPhotoConfig
	Endpoint  string
	Fieldname string
	File      interface{}
	FileID    string
	Message   Message
	Params    url.Values
	Params2   map[string]string
	Pattern   string
}

type ResponseMessage struct {
	Operation     string
	CorrelationId string

	R   APIResponse
	R2  error
	R3  string
	R4  User
	R5  bool
	R6  Message
	R7  UserProfilePhotos
	R8  File
	R9  []Update
	R10 WebhookInfo
	//R11 UpdatesChannel
	R12 Chat
	R13 []ChatMember
	R14 int
	R15 ChatMember
	R16 []GameHighScore
}

func RemoteBotDial(url string) (*RemoteBotAPI, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	rbot := new(RemoteBotAPI)
	rbot.Connection = conn

	return rbot, nil
}

func RemoteBotClose(rbot *RemoteBotAPI) {
	defer rbot.Connection.Close()
}

type RemoteBotAPI struct {
	Connection *amqp.Connection
}

func (rbot *RemoteBotAPI) MakeRequest(endpoint string, params url.Values) (APIResponse, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationMakeRequest,
		CorrelationId: randomString(32),
		Endpoint:      endpoint,
		Params:        params,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) UploadFile(endpoint string, params2 map[string]string, fieldname string, file interface{}) (APIResponse, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationUploadFile,
		CorrelationId: randomString(32),
		Endpoint:      endpoint,
		Params2:       params2,
		Fieldname:     fieldname,
		File:          file,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) GetFileDirectURL(fileID string) (string, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationGetFileDirectURL,
		CorrelationId: randomString(32),
		FileID:        fileID,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R3, response.R2
}

func (rbot *RemoteBotAPI) GetMe() (User, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationGetMe,
		CorrelationId: randomString(32),
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R4, response.R2
}

func (rbot *RemoteBotAPI) IsMessageToMe(message Message) bool {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationIsMessageToMe,
		CorrelationId: randomString(32),
		Message:       message,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R5
}

func (rbot *RemoteBotAPI) Send(c Chattable) (Message, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationSend,
		CorrelationId: randomString(32),
		C:             c,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R6, response.R2
}

func (rbot *RemoteBotAPI) GetUserProfilePhotos(config UserProfilePhotosConfig) (UserProfilePhotos, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationGetUserProfilePhotos,
		CorrelationId: randomString(32),
		Config:        config,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R7, response.R2
}

func (rbot *RemoteBotAPI) GetFile(config2 FileConfig) (File, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationGetFile,
		CorrelationId: randomString(32),
		Config2:       config2,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R8, response.R2
}

func (rbot *RemoteBotAPI) GetUpdates(config3 UpdateConfig) ([]Update, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationGetUpdates,
		CorrelationId: randomString(32),
		Config3:       config3,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R9, response.R2
}

func (rbot *RemoteBotAPI) RemoveWebhook() (APIResponse, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationRemoveWebhook,
		CorrelationId: randomString(32),
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) SetWebhook(config4 WebhookConfig) (APIResponse, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationSetWebhook,
		CorrelationId: randomString(32),
		Config4:       config4,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) GetWebhookInfo() (WebhookInfo, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationGetWebhookInfo,
		CorrelationId: randomString(32),
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R10, response.R2
}

func (rbot *RemoteBotAPI) GetUpdatesChan(config3 UpdateConfig) (UpdatesChannel, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationGetUpdatesChan,
		CorrelationId: randomString(32),
		Config3:       config3,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	c := make(UpdatesChannel)
	return c, response.R2
}

func (rbot *RemoteBotAPI) ListenForWebhook(pattern string) UpdatesChannel {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationListenForWebhook,
		CorrelationId: randomString(32),
		Pattern:       pattern,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	c := make(UpdatesChannel)
	return c
}

func (rbot *RemoteBotAPI) AnswerInlineQuery(config5 InlineConfig) (APIResponse, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationAnswerInlineQuery,
		CorrelationId: randomString(32),
		Config5:       config5,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) AnswerCallbackQuery(config6 CallbackConfig) (APIResponse, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationAnswerCallbackQuery,
		CorrelationId: randomString(32),
		Config6:       config6,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) KickChatMember(config7 KickChatMemberConfig) (APIResponse, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationKickChatMember,
		CorrelationId: randomString(32),
		Config7:       config7,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) LeaveChat(config8 ChatConfig) (APIResponse, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationLeaveChat,
		CorrelationId: randomString(32),
		Config8:       config8,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) GetChat(config8 ChatConfig) (Chat, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationGetChat,
		CorrelationId: randomString(32),
		Config8:       config8,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R12, response.R2
}

func (rbot *RemoteBotAPI) GetChatAdministrators(config8 ChatConfig) ([]ChatMember, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationGetChatAdministrators,
		CorrelationId: randomString(32),
		Config8:       config8,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R13, response.R2
}

func (rbot *RemoteBotAPI) GetChatMembersCount(config8 ChatConfig) (int, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationGetChatMembersCount,
		CorrelationId: randomString(32),
		Config8:       config8,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R14, response.R2
}

func (rbot *RemoteBotAPI) GetChatMember(config9 ChatConfigWithUser) (ChatMember, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationGetChatMember,
		CorrelationId: randomString(32),
		Config9:       config9,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R15, response.R2
}

func (rbot *RemoteBotAPI) UnbanChatMember(config10 ChatMemberConfig) (APIResponse, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationUnbanChatMember,
		CorrelationId: randomString(32),
		Config10:      config10,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) RestrictChatMember(config11 RestrictChatMemberConfig) (APIResponse, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationRestrictChatMember,
		CorrelationId: randomString(32),
		Config11:      config11,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) PromoteChatMember(config12 PromoteChatMemberConfig) (APIResponse, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationPromoteChatMember,
		CorrelationId: randomString(32),
		Config12:      config12,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) GetGameHighScores(config13 GetGameHighScoresConfig) ([]GameHighScore, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationGetGameHighScores,
		CorrelationId: randomString(32),
		Config13:      config13,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R16, response.R2
}

func (rbot *RemoteBotAPI) AnswerShippingQuery(config14 ShippingConfig) (APIResponse, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationAnswerShippingQuery,
		CorrelationId: randomString(32),
		Config14:      config14,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) AnswerPreCheckoutQuery(config15 PreCheckoutConfig) (APIResponse, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationAnswerPreCheckoutQuery,
		CorrelationId: randomString(32),
		Config15:      config15,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) DeleteMessage(config16 DeleteMessageConfig) (APIResponse, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationDeleteMessage,
		CorrelationId: randomString(32),
		Config16:      config16,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) GetInviteLink(config8 ChatConfig) (string, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationGetInviteLink,
		CorrelationId: randomString(32),
		Config8:       config8,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R3, response.R2
}

func (rbot *RemoteBotAPI) PinChatMessage(config17 PinChatMessageConfig) (APIResponse, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationPinChatMessage,
		CorrelationId: randomString(32),
		Config17:      config17,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) UnpinChatMessage(config18 UnpinChatMessageConfig) (APIResponse, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationUnpinChatMessage,
		CorrelationId: randomString(32),
		Config18:      config18,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) SetChatTitle(config19 SetChatTitleConfig) (APIResponse, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationSetChatTitle,
		CorrelationId: randomString(32),
		Config19:      config19,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) SetChatDescription(config20 SetChatDescriptionConfig) (APIResponse, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationSetChatDescription,
		CorrelationId: randomString(32),
		Config20:      config20,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) SetChatPhoto(config21 SetChatPhotoConfig) (APIResponse, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationSetChatPhoto,
		CorrelationId: randomString(32),
		Config21:      config21,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) DeleteChatPhoto(config22 DeleteChatPhotoConfig) (APIResponse, error) {
	ch, err := rbot.Connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
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

	requestMessage := RequestMessage{
		Operation:     OperationDeleteChatPhoto,
		CorrelationId: randomString(32),
		Config22:      config22,
	}
	request, _ := json.Marshal(requestMessage)

	err = ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       q.Name,
			Body:          []byte(request),
		})
	failOnError(err, "Failed to publish a message")

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			err = json.Unmarshal(d.Body, &response)
			failOnError(err, "Failed to convert body to response")
			break
		}
	}

	return response.R, response.R2
}

func SimpleServer(url string, bot *BotAPI) {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"tgbotapi", // name
		false,      // durable
		false,      // delete when usused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
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
			failOnError(err, "Failed to convert body to request")

			var r ResponseMessage
			switch n.Operation {
			case OperationMakeRequest:
				apiResponse, err := bot.MakeRequest(n.Endpoint, n.Params)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationUploadFile:
				apiResponse, err := bot.UploadFile(n.Endpoint, n.Params2, n.Fieldname, n.File)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetFileDirectURL:
				fileID, err := bot.GetFileDirectURL(n.FileID)

				r = ResponseMessage{}
				r.R3 = fileID
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetMe:
				user, err := bot.GetMe()

				r = ResponseMessage{}
				r.R4 = user
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationIsMessageToMe:
				yes := bot.IsMessageToMe(n.Message)

				r = ResponseMessage{}
				r.R5 = yes

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationSend:
				message, err := bot.Send(n.C)

				r = ResponseMessage{}
				r.R6 = message
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetUserProfilePhotos:
				userProfilePhotos, err := bot.GetUserProfilePhotos(n.Config)

				r = ResponseMessage{}
				r.R7 = userProfilePhotos
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetFile:
				file, err := bot.GetFile(n.Config2)

				r = ResponseMessage{}
				r.R8 = file
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetUpdates:
				updates, err := bot.GetUpdates(n.Config3)

				r = ResponseMessage{}
				r.R9 = updates
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationRemoveWebhook:
				apiResponse, err := bot.RemoveWebhook()

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationSetWebhook:
				apiResponse, err := bot.SetWebhook(n.Config4)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetWebhookInfo:
				webhookInfo, err := bot.GetWebhookInfo()

				r = ResponseMessage{}
				r.R10 = webhookInfo
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetUpdatesChan:
				r = ResponseMessage{}
				r.R2 = fmt.Errorf("not implemented")

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationListenForWebhook:
				r = ResponseMessage{}
				r.R2 = fmt.Errorf("not implemented")

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationAnswerInlineQuery:
				apiResponse, err := bot.AnswerInlineQuery(n.Config5)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationAnswerCallbackQuery:
				apiResponse, err := bot.AnswerCallbackQuery(n.Config6)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationKickChatMember:
				apiResponse, err := bot.KickChatMember(n.Config7)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationLeaveChat:
				apiResponse, err := bot.LeaveChat(n.Config8)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetChat:
				chat, err := bot.GetChat(n.Config8)

				r = ResponseMessage{}
				r.R12 = chat
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetChatAdministrators:
				chatMembers, err := bot.GetChatAdministrators(n.Config8)

				r = ResponseMessage{}
				r.R13 = chatMembers
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetChatMembersCount:
				count, err := bot.GetChatMembersCount(n.Config8)

				r = ResponseMessage{}
				r.R14 = count
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetChatMember:
				chatMember, err := bot.GetChatMember(n.Config9)

				r = ResponseMessage{}
				r.R15 = chatMember
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationUnbanChatMember:
				apiResponse, err := bot.UnbanChatMember(n.Config10)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationRestrictChatMember:
				apiResponse, err := bot.RestrictChatMember(n.Config11)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationPromoteChatMember:
				apiResponse, err := bot.PromoteChatMember(n.Config12)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetGameHighScores:
				gameHighScores, err := bot.GetGameHighScores(n.Config13)

				r = ResponseMessage{}
				r.R16 = gameHighScores
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationAnswerShippingQuery:
				apiResponse, err := bot.AnswerShippingQuery(n.Config14)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationAnswerPreCheckoutQuery:
				apiResponse, err := bot.AnswerPreCheckoutQuery(n.Config15)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationDeleteMessage:
				apiResponse, err := bot.DeleteMessage(n.Config16)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetInviteLink:
				link, err := bot.GetInviteLink(n.Config8)

				r = ResponseMessage{}
				r.R3 = link
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationPinChatMessage:
				apiResponse, err := bot.PinChatMessage(n.Config17)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationUnpinChatMessage:
				apiResponse, err := bot.UnpinChatMessage(n.Config18)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationSetChatTitle:
				apiResponse, err := bot.SetChatTitle(n.Config19)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationSetChatDescription:
				apiResponse, err := bot.SetChatDescription(n.Config20)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationSetChatPhoto:
				apiResponse, err := bot.SetChatPhoto(n.Config21)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationDeleteChatPhoto:
				apiResponse, err := bot.DeleteChatPhoto(n.Config22)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = err

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			default:
				r = ResponseMessage{}
				r.R2 = fmt.Errorf("not implemented")

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
