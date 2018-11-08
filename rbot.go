package tgbotapi

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"reflect"

	"github.com/streadway/amqp"
)

const (
	RandomStringLength = 32
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

type RequestMessageNoInt struct {
	Operation     string
	CorrelationId string

	CType     string
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
	FileID    string
	Message   Message
	Params    url.Values
	Params2   map[string]string
	Pattern   string
}

type RequestMessage struct {
	Operation     string
	CorrelationId string

	C         Chattable
	CType     string
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

func RequestMessageUnmarshal(data []byte) (*RequestMessage, error) {
	n := RequestMessage{}
	if err := json.Unmarshal(data, &n); err != nil {
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			// Type map
			typeMap := map[string]reflect.Type{
				"tgbotapi.MessageConfig":                reflect.TypeOf(MessageConfig{}),
				"tgbotapi.ForwardConfig":                reflect.TypeOf(ForwardConfig{}),
				"tgbotapi.PhotoConfig":                  reflect.TypeOf(PhotoConfig{}),
				"tgbotapi.AudioConfig":                  reflect.TypeOf(AudioConfig{}),
				"tgbotapi.DocumentConfig":               reflect.TypeOf(DocumentConfig{}),
				"tgbotapi.StickerConfig":                reflect.TypeOf(StickerConfig{}),
				"tgbotapi.VideoConfig":                  reflect.TypeOf(VideoConfig{}),
				"tgbotapi.AnimationConfig":              reflect.TypeOf(AnimationConfig{}),
				"tgbotapi.VideoNoteConfig":              reflect.TypeOf(VideoNoteConfig{}),
				"tgbotapi.VoiceConfig":                  reflect.TypeOf(VoiceConfig{}),
				"tgbotapi.MediaGroupConfig":             reflect.TypeOf(MediaGroupConfig{}),
				"tgbotapi.LocationConfig":               reflect.TypeOf(LocationConfig{}),
				"tgbotapi.VenueConfig":                  reflect.TypeOf(VenueConfig{}),
				"tgbotapi.ContactConfig":                reflect.TypeOf(ContactConfig{}),
				"tgbotapi.GameConfig":                   reflect.TypeOf(GameConfig{}),
				"tgbotapi.SetGameScoreConfig":           reflect.TypeOf(SetGameScoreConfig{}),
				"tgbotapi.GetGameHighScoresConfig":      reflect.TypeOf(GetGameHighScoresConfig{}),
				"tgbotapi.ChatActionConfig":             reflect.TypeOf(ChatActionConfig{}),
				"tgbotapi.EditMessageTextConfig":        reflect.TypeOf(EditMessageTextConfig{}),
				"tgbotapi.EditMessageCaptionConfig":     reflect.TypeOf(EditMessageCaptionConfig{}),
				"tgbotapi.EditMessageReplyMarkupConfig": reflect.TypeOf(EditMessageReplyMarkupConfig{}),
				"tgbotapi.InvoiceConfig":                reflect.TypeOf(InvoiceConfig{}),
				"tgbotapi.DeleteMessageConfig":          reflect.TypeOf(DeleteMessageConfig{}),
				"tgbotapi.PinChatMessageConfig":         reflect.TypeOf(PinChatMessageConfig{}),
				"tgbotapi.UnpinChatMessageConfig":       reflect.TypeOf(UnpinChatMessageConfig{}),
				"tgbotapi.SetChatTitleConfig":           reflect.TypeOf(SetChatTitleConfig{}),
				"tgbotapi.SetChatDescriptionConfig":     reflect.TypeOf(SetChatDescriptionConfig{}),
				"tgbotapi.SetChatPhotoConfig":           reflect.TypeOf(SetChatPhotoConfig{}),
				"tgbotapi.DeleteChatPhotoConfig":        reflect.TypeOf(DeleteChatPhotoConfig{}),
			}

			// Unmarshal raw
			m := map[string]interface{}{}
			if err := json.Unmarshal(data, &m); err != nil {
				return nil, err
			}

			// Discover type of C based on CType
			if cType, ok := m["CType"].(string); ok {
				if reflectType, found := typeMap[cType]; found {
					// Annotate type of C
					n.C = reflect.New(reflectType).Interface().(Chattable)
					n.CType = cType
				}

				// Marshal only C data
				valueBytes, err := json.Marshal(m["C"])
				if err != nil {
					return nil, err
				}

				// Unmarshal only C data
				if err = json.Unmarshal(valueBytes, &n.C); err != nil {
					return nil, err
				}
			}

			// Hard coded
			var nI RequestMessageNoInt
			if err = json.Unmarshal(data, &nI); err != nil {
				return nil, err
			}

			n.Operation = nI.Operation
			n.CorrelationId = nI.CorrelationId
			n.CType = nI.CType
			n.Config = nI.Config
			n.Config2 = nI.Config2
			n.Config3 = nI.Config3
			n.Config4 = nI.Config4
			n.Config5 = nI.Config5
			n.Config6 = nI.Config6
			n.Config7 = nI.Config7
			n.Config8 = nI.Config8
			n.Config9 = nI.Config9
			n.Config10 = nI.Config10
			n.Config11 = nI.Config11
			n.Config12 = nI.Config12
			n.Config13 = nI.Config13
			n.Config14 = nI.Config14
			n.Config15 = nI.Config15
			n.Config16 = nI.Config16
			n.Config17 = nI.Config17
			n.Config18 = nI.Config18
			n.Config19 = nI.Config19
			n.Config20 = nI.Config20
			n.Config21 = nI.Config21
			n.Config22 = nI.Config22
			n.Endpoint = nI.Endpoint
			n.Fieldname = nI.Fieldname
			n.FileID = nI.FileID
			n.Message = nI.Message
			n.Params = nI.Params
			n.Params2 = nI.Params2
			n.Pattern = nI.Pattern

			return &n, nil
		}

		return nil, err
	}

	return &n, nil
}

type ResponseMessageNoInt struct {
	Operation     string
	CorrelationId string

	R      APIResponse
	R2Type string
	R3     string
	R4     User
	R5     bool
	R6     Message
	R7     UserProfilePhotos
	R8     File
	R9     []Update
	R10    WebhookInfo
	//R11 UpdatesChannel
	R12 Chat
	R13 []ChatMember
	R14 int
	R15 ChatMember
	R16 []GameHighScore
}

type ResponseMessage struct {
	Operation     string
	CorrelationId string

	R      APIResponse
	R2     error
	R2Type string
	R3     string
	R4     User
	R5     bool
	R6     Message
	R7     UserProfilePhotos
	R8     File
	R9     []Update
	R10    WebhookInfo
	//R11 UpdatesChannel
	R12 Chat
	R13 []ChatMember
	R14 int
	R15 ChatMember
	R16 []GameHighScore
}

type ErrorString struct {
	s string
}

func (e *ErrorString) Error() string {
	return e.s
}

func convertToErrorString(e error) error {
	if e != nil {
		return &ErrorString{e.Error()}
	}

	return nil
}

type ErrorRemoteBot struct {
	s string
	i error
}

func (e *ErrorRemoteBot) Error() string {
	if e.i != nil {
		return fmt.Sprintf("%s: %s", e.s, e.i)
	}

	return e.s
}

func convertToErrorRemoteBot(s string, e error) error {
	return &ErrorRemoteBot{s, e}
}

func ResponseMessageUnmarshal(data []byte) (*ResponseMessage, error) {

	n := ResponseMessage{}
	if err := json.Unmarshal(data, &n); err != nil {
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			// Type map
			typeMap := map[string]reflect.Type{
				"*errors.errorString":   reflect.TypeOf(ErrorString{}),
				"*tgbotapi.ErrorString": reflect.TypeOf(ErrorString{}),
			}

			// Unmarshal raw
			m := map[string]interface{}{}
			if err := json.Unmarshal(data, &m); err != nil {
				return nil, err
			}

			// Discover type of R2 based on R2Type
			if r2Type, ok := m["R2Type"].(string); ok {
				if reflectType, found := typeMap[r2Type]; found {
					// Annotate type of R2
					n.R2 = reflect.New(reflectType).Interface().(error)
					n.R2Type = r2Type
				}

				// Marshal only R2 data
				valueBytes, err := json.Marshal(m["R2"])
				if err != nil {
					return nil, err
				}

				// Unmarshal only R2 data
				if err = json.Unmarshal(valueBytes, &n.R2); err != nil {
					return nil, err
				}
			}

			// Hard coded
			var nI ResponseMessageNoInt
			if err = json.Unmarshal(data, &nI); err != nil {
				return nil, err
			}

			n.Operation = nI.Operation
			n.CorrelationId = nI.CorrelationId
			n.R = nI.R
			n.R2Type = nI.R2Type
			n.R3 = nI.R3
			n.R4 = nI.R4
			n.R5 = nI.R5
			n.R6 = nI.R6
			n.R7 = nI.R7
			n.R8 = nI.R8
			n.R9 = nI.R9
			n.R10 = nI.R10
			n.R12 = nI.R12
			n.R13 = nI.R13
			n.R14 = nI.R14
			n.R15 = nI.R15
			n.R16 = nI.R16

			return &n, nil
		}

		return nil, err
	}

	return &n, nil
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
	rbot.Connection.Close()
}

type RemoteBotAPI struct {
	Connection *amqp.Connection
}

func hChannelQueueDeclare(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

func hChannelConsume(ch *amqp.Channel, name string) (<-chan amqp.Delivery, error) {
	return ch.Consume(
		name,  // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
}

func hChannelPublish(ch *amqp.Channel, requestMessage *RequestMessage, name string, request []byte) error {
	return ch.Publish(
		"",         // exchange
		"tgbotapi", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       name,
			Body:          request,
		})
}

func (rbot *RemoteBotAPI) MakeRequest(endpoint string, params url.Values) (APIResponse, error) {
	var result APIResponse

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationMakeRequest,
		CorrelationId: randomString(RandomStringLength),
		Endpoint:      endpoint,
		Params:        params,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) UploadFile(endpoint string, params2 map[string]string, fieldname string, file interface{}) (APIResponse, error) {
	var result APIResponse

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationUploadFile,
		CorrelationId: randomString(RandomStringLength),
		Endpoint:      endpoint,
		Params2:       params2,
		Fieldname:     fieldname,
		File:          file,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) GetFileDirectURL(fileID string) (string, error) {
	var result string

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationGetFileDirectURL,
		CorrelationId: randomString(RandomStringLength),
		FileID:        fileID,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R3, response.R2
}

func (rbot *RemoteBotAPI) GetMe() (User, error) {
	var result User

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationGetMe,
		CorrelationId: randomString(RandomStringLength),
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R4, response.R2
}

func (rbot *RemoteBotAPI) IsMessageToMe(message Message) bool {
	ch, err := rbot.Connection.Channel()
	if err != nil {
		//return result, convertToErrorRemoteBot("Failed to open a channel", err)
		return false
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		//return result, convertToErrorRemoteBot("Failed to declare a queue", err)
		return false
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		//return result, convertToErrorRemoteBot("Failed to register a consumer", err)
		return false
	}

	requestMessage := RequestMessage{
		Operation:     OperationIsMessageToMe,
		CorrelationId: randomString(RandomStringLength),
		Message:       message,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		//return result, convertToErrorRemoteBot("Failed to publish a message", err)
		return false
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				//return result, convertToErrorRemoteBot("Failed to convert body to response", err)
				return false
			}
			response = *rP
			break
		}
	}

	return response.R5
}

func (rbot *RemoteBotAPI) Send(c Chattable) (result Message, err error) {
	cType := reflect.TypeOf(c).String()

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationSend,
		CorrelationId: randomString(RandomStringLength),
		C:             c,
		CType:         cType,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R6, response.R2
}

func (rbot *RemoteBotAPI) GetUserProfilePhotos(config UserProfilePhotosConfig) (UserProfilePhotos, error) {
	var result UserProfilePhotos

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationGetUserProfilePhotos,
		CorrelationId: randomString(RandomStringLength),
		Config:        config,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R7, response.R2
}

func (rbot *RemoteBotAPI) GetFile(config2 FileConfig) (File, error) {
	var result File

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationGetFile,
		CorrelationId: randomString(RandomStringLength),
		Config2:       config2,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R8, response.R2
}

func (rbot *RemoteBotAPI) GetUpdates(config3 UpdateConfig) ([]Update, error) {
	var result []Update

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationGetUpdates,
		CorrelationId: randomString(RandomStringLength),
		Config3:       config3,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R9, response.R2
}

func (rbot *RemoteBotAPI) RemoveWebhook() (APIResponse, error) {
	var result APIResponse

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationRemoveWebhook,
		CorrelationId: randomString(RandomStringLength),
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) SetWebhook(config4 WebhookConfig) (APIResponse, error) {
	var result APIResponse

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationSetWebhook,
		CorrelationId: randomString(RandomStringLength),
		Config4:       config4,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) GetWebhookInfo() (WebhookInfo, error) {
	var result WebhookInfo

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationGetWebhookInfo,
		CorrelationId: randomString(RandomStringLength),
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R10, response.R2
}

func (rbot *RemoteBotAPI) GetUpdatesChan(config3 UpdateConfig) (UpdatesChannel, error) {
	var result UpdatesChannel

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationGetUpdatesChan,
		CorrelationId: randomString(RandomStringLength),
		Config3:       config3,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	c := make(UpdatesChannel)
	return c, response.R2
}

func (rbot *RemoteBotAPI) ListenForWebhook(pattern string) UpdatesChannel {
	c := make(UpdatesChannel)

	ch, err := rbot.Connection.Channel()
	if err != nil {
		//return result, convertToErrorRemoteBot("Failed to open a channel", err)
		return c
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		//return result, convertToErrorRemoteBot("Failed to declare a queue", err)
		return c
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		//return result, convertToErrorRemoteBot("Failed to register a consumer", err)
		return c
	}

	requestMessage := RequestMessage{
		Operation:     OperationListenForWebhook,
		CorrelationId: randomString(RandomStringLength),
		Pattern:       pattern,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		//return result, convertToErrorRemoteBot("Failed to publish a message", err)
		return c
	}

	//var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			//rP, err := ResponseMessageUnmarshal(d.Body)
			//if err != nil {
			//	//return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			//      return c
			//}
			//response = *rP
			break
		}
	}

	return c
}

func (rbot *RemoteBotAPI) AnswerInlineQuery(config5 InlineConfig) (APIResponse, error) {
	var result APIResponse

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationAnswerInlineQuery,
		CorrelationId: randomString(RandomStringLength),
		Config5:       config5,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) AnswerCallbackQuery(config6 CallbackConfig) (APIResponse, error) {
	var result APIResponse

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationAnswerCallbackQuery,
		CorrelationId: randomString(RandomStringLength),
		Config6:       config6,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) KickChatMember(config7 KickChatMemberConfig) (APIResponse, error) {
	var result APIResponse

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationKickChatMember,
		CorrelationId: randomString(RandomStringLength),
		Config7:       config7,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) LeaveChat(config8 ChatConfig) (APIResponse, error) {
	var result APIResponse

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationLeaveChat,
		CorrelationId: randomString(RandomStringLength),
		Config8:       config8,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) GetChat(config8 ChatConfig) (Chat, error) {
	var result Chat

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationGetChat,
		CorrelationId: randomString(RandomStringLength),
		Config8:       config8,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R12, response.R2
}

func (rbot *RemoteBotAPI) GetChatAdministrators(config8 ChatConfig) ([]ChatMember, error) {
	var result []ChatMember

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationGetChatAdministrators,
		CorrelationId: randomString(RandomStringLength),
		Config8:       config8,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R13, response.R2
}

func (rbot *RemoteBotAPI) GetChatMembersCount(config8 ChatConfig) (int, error) {
	var result int

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationGetChatMembersCount,
		CorrelationId: randomString(RandomStringLength),
		Config8:       config8,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R14, response.R2
}

func (rbot *RemoteBotAPI) GetChatMember(config9 ChatConfigWithUser) (ChatMember, error) {
	var result ChatMember

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationGetChatMember,
		CorrelationId: randomString(RandomStringLength),
		Config9:       config9,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R15, response.R2
}

func (rbot *RemoteBotAPI) UnbanChatMember(config10 ChatMemberConfig) (APIResponse, error) {
	var result APIResponse

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationUnbanChatMember,
		CorrelationId: randomString(RandomStringLength),
		Config10:      config10,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) RestrictChatMember(config11 RestrictChatMemberConfig) (APIResponse, error) {
	var result APIResponse

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationRestrictChatMember,
		CorrelationId: randomString(RandomStringLength),
		Config11:      config11,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) PromoteChatMember(config12 PromoteChatMemberConfig) (APIResponse, error) {
	var result APIResponse

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationPromoteChatMember,
		CorrelationId: randomString(RandomStringLength),
		Config12:      config12,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) GetGameHighScores(config13 GetGameHighScoresConfig) ([]GameHighScore, error) {
	var result []GameHighScore

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationGetGameHighScores,
		CorrelationId: randomString(RandomStringLength),
		Config13:      config13,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R16, response.R2
}

func (rbot *RemoteBotAPI) AnswerShippingQuery(config14 ShippingConfig) (APIResponse, error) {
	var result APIResponse

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationAnswerShippingQuery,
		CorrelationId: randomString(RandomStringLength),
		Config14:      config14,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) AnswerPreCheckoutQuery(config15 PreCheckoutConfig) (APIResponse, error) {
	var result APIResponse

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationAnswerPreCheckoutQuery,
		CorrelationId: randomString(RandomStringLength),
		Config15:      config15,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) DeleteMessage(config16 DeleteMessageConfig) (APIResponse, error) {
	var result APIResponse

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationDeleteMessage,
		CorrelationId: randomString(RandomStringLength),
		Config16:      config16,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) GetInviteLink(config8 ChatConfig) (string, error) {
	var result string

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationGetInviteLink,
		CorrelationId: randomString(RandomStringLength),
		Config8:       config8,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R3, response.R2
}

func (rbot *RemoteBotAPI) PinChatMessage(config17 PinChatMessageConfig) (APIResponse, error) {
	var result APIResponse

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationPinChatMessage,
		CorrelationId: randomString(RandomStringLength),
		Config17:      config17,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) UnpinChatMessage(config18 UnpinChatMessageConfig) (APIResponse, error) {
	var result APIResponse

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationUnpinChatMessage,
		CorrelationId: randomString(RandomStringLength),
		Config18:      config18,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) SetChatTitle(config19 SetChatTitleConfig) (APIResponse, error) {
	var result APIResponse

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationSetChatTitle,
		CorrelationId: randomString(RandomStringLength),
		Config19:      config19,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) SetChatDescription(config20 SetChatDescriptionConfig) (APIResponse, error) {
	var result APIResponse

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationSetChatDescription,
		CorrelationId: randomString(RandomStringLength),
		Config20:      config20,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) SetChatPhoto(config21 SetChatPhotoConfig) (APIResponse, error) {
	var result APIResponse

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationSetChatPhoto,
		CorrelationId: randomString(RandomStringLength),
		Config21:      config21,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R, response.R2
}

func (rbot *RemoteBotAPI) DeleteChatPhoto(config22 DeleteChatPhotoConfig) (APIResponse, error) {
	var result APIResponse

	ch, err := rbot.Connection.Channel()
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to open a channel", err)
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to declare a queue", err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to register a consumer", err)
	}

	requestMessage := RequestMessage{
		Operation:     OperationDeleteChatPhoto,
		CorrelationId: randomString(RandomStringLength),
		Config22:      config22,
	}
	request, _ := json.Marshal(requestMessage)

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		return result, convertToErrorRemoteBot("Failed to publish a message", err)
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				return result, convertToErrorRemoteBot("Failed to convert body to response", err)
			}
			response = *rP
			break
		}
	}

	return response.R, response.R2
}

func SimpleServer(url string, bot *BotAPI) {
	conn, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}
	//failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	//failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"tgbotapi", // name
		false,      // durable
		false,      // delete when usused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	//failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	//failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	//failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			nP, err := RequestMessageUnmarshal(d.Body)
			//failOnError(err, "Failed to convert body to request")
			if err != nil {
				panic(err)
			}
			n := *nP

			var r ResponseMessage
			switch n.Operation {
			case OperationMakeRequest:
				apiResponse, err := bot.MakeRequest(n.Endpoint, n.Params)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationUploadFile:
				apiResponse, err := bot.UploadFile(n.Endpoint, n.Params2, n.Fieldname, n.File)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetFileDirectURL:
				fileID, err := bot.GetFileDirectURL(n.FileID)

				r = ResponseMessage{}
				r.R3 = fileID
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetMe:
				user, err := bot.GetMe()

				r = ResponseMessage{}
				r.R4 = user
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationIsMessageToMe:
				yes := bot.IsMessageToMe(n.Message)

				r = ResponseMessage{}
				r.R5 = yes

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationSend:
				switch n.CType {
				case reflect.TypeOf(MessageConfig{}).String():
					m := n.C.(*MessageConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(ForwardConfig{}).String():
					m := n.C.(*ForwardConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(PhotoConfig{}).String():
					m := n.C.(*PhotoConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(AudioConfig{}).String():
					m := n.C.(*AudioConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(DocumentConfig{}).String():
					m := n.C.(*DocumentConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(StickerConfig{}).String():
					m := n.C.(*StickerConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(VideoConfig{}).String():
					m := n.C.(*VideoConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(AnimationConfig{}).String():
					m := n.C.(*AnimationConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(VideoNoteConfig{}).String():
					m := n.C.(*VideoNoteConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(VoiceConfig{}).String():
					m := n.C.(*VoiceConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(MediaGroupConfig{}).String():
					m := n.C.(*MediaGroupConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(LocationConfig{}).String():
					m := n.C.(*LocationConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(VenueConfig{}).String():
					m := n.C.(*VenueConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(ContactConfig{}).String():
					m := n.C.(*ContactConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(GameConfig{}).String():
					m := n.C.(*GameConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(SetGameScoreConfig{}).String():
					m := n.C.(*SetGameScoreConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(GetGameHighScoresConfig{}).String():
					m := n.C.(*GetGameHighScoresConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(ChatActionConfig{}).String():
					m := n.C.(*ChatActionConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(EditMessageTextConfig{}).String():
					m := n.C.(*EditMessageTextConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(EditMessageCaptionConfig{}).String():
					m := n.C.(*EditMessageCaptionConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(EditMessageReplyMarkupConfig{}).String():
					m := n.C.(*EditMessageReplyMarkupConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(InvoiceConfig{}).String():
					m := n.C.(*InvoiceConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(DeleteMessageConfig{}).String():
					m := n.C.(*DeleteMessageConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(PinChatMessageConfig{}).String():
					m := n.C.(*PinChatMessageConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(UnpinChatMessageConfig{}).String():
					m := n.C.(*UnpinChatMessageConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(SetChatTitleConfig{}).String():
					m := n.C.(*SetChatTitleConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(SetChatDescriptionConfig{}).String():
					m := n.C.(*SetChatDescriptionConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(SetChatPhotoConfig{}).String():
					m := n.C.(*SetChatPhotoConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				case reflect.TypeOf(DeleteChatPhotoConfig{}).String():
					m := n.C.(*DeleteChatPhotoConfig)
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = convertToErrorString(err)
				default:
					r = ResponseMessage{}
					r.R2 = fmt.Errorf("not implemented")
				}

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetUserProfilePhotos:
				userProfilePhotos, err := bot.GetUserProfilePhotos(n.Config)

				r = ResponseMessage{}
				r.R7 = userProfilePhotos
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetFile:
				file, err := bot.GetFile(n.Config2)

				r = ResponseMessage{}
				r.R8 = file
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetUpdates:
				updates, err := bot.GetUpdates(n.Config3)

				r = ResponseMessage{}
				r.R9 = updates
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationRemoveWebhook:
				apiResponse, err := bot.RemoveWebhook()

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationSetWebhook:
				apiResponse, err := bot.SetWebhook(n.Config4)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetWebhookInfo:
				webhookInfo, err := bot.GetWebhookInfo()

				r = ResponseMessage{}
				r.R10 = webhookInfo
				r.R2 = convertToErrorString(err)

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
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationAnswerCallbackQuery:
				apiResponse, err := bot.AnswerCallbackQuery(n.Config6)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationKickChatMember:
				apiResponse, err := bot.KickChatMember(n.Config7)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationLeaveChat:
				apiResponse, err := bot.LeaveChat(n.Config8)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetChat:
				chat, err := bot.GetChat(n.Config8)

				r = ResponseMessage{}
				r.R12 = chat
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetChatAdministrators:
				chatMembers, err := bot.GetChatAdministrators(n.Config8)

				r = ResponseMessage{}
				r.R13 = chatMembers
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetChatMembersCount:
				count, err := bot.GetChatMembersCount(n.Config8)

				r = ResponseMessage{}
				r.R14 = count
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetChatMember:
				chatMember, err := bot.GetChatMember(n.Config9)

				r = ResponseMessage{}
				r.R15 = chatMember
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationUnbanChatMember:
				apiResponse, err := bot.UnbanChatMember(n.Config10)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationRestrictChatMember:
				apiResponse, err := bot.RestrictChatMember(n.Config11)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationPromoteChatMember:
				apiResponse, err := bot.PromoteChatMember(n.Config12)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetGameHighScores:
				gameHighScores, err := bot.GetGameHighScores(n.Config13)

				r = ResponseMessage{}
				r.R16 = gameHighScores
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationAnswerShippingQuery:
				apiResponse, err := bot.AnswerShippingQuery(n.Config14)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationAnswerPreCheckoutQuery:
				apiResponse, err := bot.AnswerPreCheckoutQuery(n.Config15)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationDeleteMessage:
				apiResponse, err := bot.DeleteMessage(n.Config16)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetInviteLink:
				link, err := bot.GetInviteLink(n.Config8)

				r = ResponseMessage{}
				r.R3 = link
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationPinChatMessage:
				apiResponse, err := bot.PinChatMessage(n.Config17)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationUnpinChatMessage:
				apiResponse, err := bot.UnpinChatMessage(n.Config18)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationSetChatTitle:
				apiResponse, err := bot.SetChatTitle(n.Config19)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationSetChatDescription:
				apiResponse, err := bot.SetChatDescription(n.Config20)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationSetChatPhoto:
				apiResponse, err := bot.SetChatPhoto(n.Config21)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationDeleteChatPhoto:
				apiResponse, err := bot.DeleteChatPhoto(n.Config22)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = convertToErrorString(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			default:
				r = ResponseMessage{}
				r.R2 = fmt.Errorf("not implemented")

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			}

			// Annotate R2Type if needed
			if r.R2 != nil {
				r.R2Type = reflect.TypeOf(r.R2).String()
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
					Body:          response,
				})
			//failOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
	}()

	<-forever
}
