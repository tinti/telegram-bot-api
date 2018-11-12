package tgbotapi

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"reflect"
	"time"

	"github.com/streadway/amqp"
)

const (
	RandomStringLength = 32
	RoutingKey         = "tgbotapi"
	DefaultTimeout     = 15 * time.Second
)

const (
	FailedConnect             = "failed to connect to RabbitMQ"
	FailedConvertBodyRequest  = "failed to convert body to request"
	FailedConvertBodyResponse = "failed to convert body to response"
	FailedDeclareQueue        = "failed to declare a queue"
	FailedOpenChannel         = "failed to open a channel"
	FailedMessagePublish      = "failed to publish a message"
	FailedMessageConsume      = "failed to register a consumer"
	FailedOptionQoS           = "failed to set QoS"
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
var _ BotAPIIface = (*RemoteBotAPI)(nil)

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

type ConcreteChattable struct {
	Type string

	ValueMessageConfig                MessageConfig
	ValueForwardConfig                ForwardConfig
	ValuePhotoConfig                  PhotoConfig
	ValueAudioConfig                  AudioConfig
	ValueDocumentConfig               DocumentConfig
	ValueStickerConfig                StickerConfig
	ValueVideoConfig                  VideoConfig
	ValueAnimationConfig              AnimationConfig
	ValueVideoNoteConfig              VideoNoteConfig
	ValueVoiceConfig                  VoiceConfig
	ValueMediaGroupConfig             MediaGroupConfig
	ValueLocationConfig               LocationConfig
	ValueVenueConfig                  VenueConfig
	ValueContactConfig                ContactConfig
	ValueGameConfig                   GameConfig
	ValueSetGameScoreConfig           SetGameScoreConfig
	ValueGetGameHighScoresConfig      GetGameHighScoresConfig
	ValueChatActionConfig             ChatActionConfig
	ValueEditMessageTextConfig        EditMessageTextConfig
	ValueEditMessageCaptionConfig     EditMessageCaptionConfig
	ValueEditMessageReplyMarkupConfig EditMessageReplyMarkupConfig
	ValueInvoiceConfig                InvoiceConfig
	ValueDeleteMessageConfig          DeleteMessageConfig
	ValuePinChatMessageConfig         PinChatMessageConfig
	ValueUnpinChatMessageConfig       UnpinChatMessageConfig
	ValueSetChatTitleConfig           SetChatTitleConfig
	ValueSetChatDescriptionConfig     SetChatDescriptionConfig
	ValueDeleteChatPhotoConfig        DeleteChatPhotoConfig
	//ValueSetChatPhotoConfig           SetChatPhotoConfig
}

func hConvertToConcreteChattable(c Chattable) ConcreteChattable {
	var cC ConcreteChattable
	cC.Type = reflect.TypeOf(c).String()
	switch cC.Type {
	case reflect.TypeOf(MessageConfig{}).String():
		cC.ValueMessageConfig = c.(MessageConfig)
	case reflect.TypeOf(ForwardConfig{}).String():
		cC.ValueForwardConfig = c.(ForwardConfig)
	case reflect.TypeOf(PhotoConfig{}).String():
		cC.ValuePhotoConfig = c.(PhotoConfig)
	case reflect.TypeOf(AudioConfig{}).String():
		cC.ValueAudioConfig = c.(AudioConfig)
	case reflect.TypeOf(DocumentConfig{}).String():
		cC.ValueDocumentConfig = c.(DocumentConfig)
	case reflect.TypeOf(StickerConfig{}).String():
		cC.ValueStickerConfig = c.(StickerConfig)
	case reflect.TypeOf(VideoConfig{}).String():
		cC.ValueVideoConfig = c.(VideoConfig)
	case reflect.TypeOf(AnimationConfig{}).String():
		cC.ValueAnimationConfig = c.(AnimationConfig)
	case reflect.TypeOf(VideoNoteConfig{}).String():
		cC.ValueVideoNoteConfig = c.(VideoNoteConfig)
	case reflect.TypeOf(VoiceConfig{}).String():
		cC.ValueVoiceConfig = c.(VoiceConfig)
	case reflect.TypeOf(MediaGroupConfig{}).String():
		cC.ValueMediaGroupConfig = c.(MediaGroupConfig)
	case reflect.TypeOf(LocationConfig{}).String():
		cC.ValueLocationConfig = c.(LocationConfig)
	case reflect.TypeOf(VenueConfig{}).String():
		cC.ValueVenueConfig = c.(VenueConfig)
	case reflect.TypeOf(ContactConfig{}).String():
		cC.ValueContactConfig = c.(ContactConfig)
	case reflect.TypeOf(GameConfig{}).String():
		cC.ValueGameConfig = c.(GameConfig)
	case reflect.TypeOf(SetGameScoreConfig{}).String():
		cC.ValueSetGameScoreConfig = c.(SetGameScoreConfig)
	case reflect.TypeOf(GetGameHighScoresConfig{}).String():
		cC.ValueGetGameHighScoresConfig = c.(GetGameHighScoresConfig)
	case reflect.TypeOf(ChatActionConfig{}).String():
		cC.ValueChatActionConfig = c.(ChatActionConfig)
	case reflect.TypeOf(EditMessageTextConfig{}).String():
		cC.ValueEditMessageTextConfig = c.(EditMessageTextConfig)
	case reflect.TypeOf(EditMessageCaptionConfig{}).String():
		cC.ValueEditMessageCaptionConfig = c.(EditMessageCaptionConfig)
	case reflect.TypeOf(EditMessageReplyMarkupConfig{}).String():
		cC.ValueEditMessageReplyMarkupConfig = c.(EditMessageReplyMarkupConfig)
	case reflect.TypeOf(InvoiceConfig{}).String():
		cC.ValueInvoiceConfig = c.(InvoiceConfig)
	case reflect.TypeOf(DeleteMessageConfig{}).String():
		cC.ValueDeleteMessageConfig = c.(DeleteMessageConfig)
	case reflect.TypeOf(PinChatMessageConfig{}).String():
		cC.ValuePinChatMessageConfig = c.(PinChatMessageConfig)
	case reflect.TypeOf(UnpinChatMessageConfig{}).String():
		cC.ValueUnpinChatMessageConfig = c.(UnpinChatMessageConfig)
	case reflect.TypeOf(SetChatTitleConfig{}).String():
		cC.ValueSetChatTitleConfig = c.(SetChatTitleConfig)
	case reflect.TypeOf(SetChatDescriptionConfig{}).String():
		cC.ValueSetChatDescriptionConfig = c.(SetChatDescriptionConfig)
	case reflect.TypeOf(DeleteChatPhotoConfig{}).String():
		cC.ValueDeleteChatPhotoConfig = c.(DeleteChatPhotoConfig)
	//case reflect.TypeOf(SetChatPhotoConfig{}).String():
	//	cC.ValueSetChatPhotoConfig = c.(SetChatPhotoConfig)
	default:
		panic("can't convert chattable to concretechattable")
	}
	return cC
}

type RequestMessage struct {
	Operation     string
	CorrelationId string

	C         ConcreteChattable
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
	var n RequestMessage
	err := json.Unmarshal(data, &n)
	return &n, err
}

type ConcreteError struct {
	IsNil bool
	Value string
}

func (e *ConcreteError) toError() error {
	if e.IsNil {
		return nil
	}

	return &ErrorString{e.Value}
}

func hConvertToConcreteError(e error) ConcreteError {
	if e != nil {
		return ConcreteError{false, e.Error()}
	}

	return ConcreteError{true, ""}
}

type ResponseMessage struct {
	Operation     string
	CorrelationId string

	R   APIResponse
	R2  ConcreteError
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

func ResponseMessageUnmarshal(data []byte) (*ResponseMessage, error) {
	var n ResponseMessage
	err := json.Unmarshal(data, &n)
	return &n, err
}

type ErrorString struct {
	s string
}

func (e *ErrorString) Error() string {
	return e.s
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

func hNewErrorRemoteBot(s string, e error) error {
	return &ErrorRemoteBot{s, e}
}

func RemoteBotDial(url string) (*RemoteBotAPI, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	rbot := new(RemoteBotAPI)
	rbot.Connection = conn
	rbot.Timeout = DefaultTimeout

	return rbot, nil
}

func RemoteBotClose(rbot *RemoteBotAPI) {
	rbot.Connection.Close()
}

type RemoteBotAPI struct {
	Connection *amqp.Connection
	Timeout    time.Duration
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
		RoutingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: requestMessage.CorrelationId,
			ReplyTo:       name,
			Body:          request,
		})
}

func hCreateChannelQueueDeclareConsume(connection *amqp.Connection) (*amqp.Channel, amqp.Queue, <-chan amqp.Delivery, error) {
	ch, err := connection.Channel()
	if err != nil {
		return nil, amqp.Queue{}, make(chan amqp.Delivery), hNewErrorRemoteBot(FailedOpenChannel, err)
	}

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		ch.Close()
		return nil, amqp.Queue{}, make(chan amqp.Delivery), hNewErrorRemoteBot(FailedDeclareQueue, err)
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		ch.Close()
		return nil, amqp.Queue{}, make(chan amqp.Delivery), hNewErrorRemoteBot(FailedMessageConsume, err)
	}

	return ch, q, msgs, nil
}

func hChannelPublishWithTimeoutTicker(ch *amqp.Channel, requestMessage *RequestMessage, name string, request []byte, ticker *time.Ticker) error {
	errChan := make(chan error, 1)

	go func() {
		errChan <- hChannelPublish(ch, requestMessage, name, request)
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return hNewErrorRemoteBot(FailedMessagePublish, err)
		}
	case <-ticker.C:
		return hNewErrorRemoteBot(FailedMessagePublish, fmt.Errorf("timeout"))
	}

	return nil
}

func hChannelConsumeWithTimeoutTicker(correlationId string, msgs <-chan amqp.Delivery, ticker *time.Ticker) (*ResponseMessage, error) {
	var response *ResponseMessage
	errChan := make(chan error, 1)

	go func() {
		for d := range msgs {
			if correlationId == d.CorrelationId {
				r, err := ResponseMessageUnmarshal(d.Body)
				response = r
				errChan <- err
				break
			}
		}
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return nil, hNewErrorRemoteBot(FailedConvertBodyResponse, err)
		}
	case <-ticker.C:
		return nil, hNewErrorRemoteBot(FailedMessageConsume, fmt.Errorf("timeout"))
	}

	return response, nil
}

func hChannelSerializePublishConsumeWithTimeoutTicker(ch *amqp.Channel, q amqp.Queue, msgs <-chan amqp.Delivery, requestMessage *RequestMessage, ticker *time.Ticker) (*ResponseMessage, error) {
	request, err := json.Marshal(*requestMessage)
	if err != nil {
		panic(err)
	}

	remoteBotErr := hChannelPublishWithTimeoutTicker(ch, requestMessage, q.Name, request, ticker)
	if remoteBotErr != nil {
		return nil, remoteBotErr
	}

	response, remoteBotErr := hChannelConsumeWithTimeoutTicker(requestMessage.CorrelationId, msgs, ticker)
	if remoteBotErr != nil {
		return nil, remoteBotErr
	}

	return response, nil
}

func (rbot *RemoteBotAPI) MakeRequest(endpoint string, params url.Values) (APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result APIResponse

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationMakeRequest,
		CorrelationId: randomString(RandomStringLength),
		Endpoint:      endpoint,
		Params:        params,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.toError()
}

func (rbot *RemoteBotAPI) UploadFile(endpoint string, params2 map[string]string, fieldname string, file interface{}) (APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result APIResponse

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationUploadFile,
		CorrelationId: randomString(RandomStringLength),
		Endpoint:      endpoint,
		Params2:       params2,
		Fieldname:     fieldname,
		File:          file,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.toError()
}

func (rbot *RemoteBotAPI) GetFileDirectURL(fileID string) (string, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result string

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetFileDirectURL,
		CorrelationId: randomString(RandomStringLength),
		FileID:        fileID,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R3, response.R2.toError()
}

func (rbot *RemoteBotAPI) GetMe() (User, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result User

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetMe,
		CorrelationId: randomString(RandomStringLength),
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R4, response.R2.toError()
}

func (rbot *RemoteBotAPI) IsMessageToMe(message Message) bool {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	ch, err := rbot.Connection.Channel()
	if err != nil {
		//return result, hNewErrorRemoteBot(FailedOpenChannel, err)
		return false
	}
	defer ch.Close()

	q, err := hChannelQueueDeclare(ch)
	if err != nil {
		//return result, hNewErrorRemoteBot(FailedDeclareQueue, err)
		return false
	}

	msgs, err := hChannelConsume(ch, q.Name)
	if err != nil {
		//return result, hNewErrorRemoteBot(FailedMessageConsume, err)
		return false
	}

	requestMessage := RequestMessage{
		Operation:     OperationIsMessageToMe,
		CorrelationId: randomString(RandomStringLength),
		Message:       message,
	}
	request, err := json.Marshal(requestMessage)
	if err != nil {
		panic(err)
	}

	err = hChannelPublish(ch, &requestMessage, q.Name, request)
	if err != nil {
		//return result, hNewErrorRemoteBot(FailedMessagePublish, err)
		return false
	}

	var response ResponseMessage
	for d := range msgs {
		if requestMessage.CorrelationId == d.CorrelationId {
			rP, err := ResponseMessageUnmarshal(d.Body)
			if err != nil {
				//return result, hNewErrorRemoteBot(FailedConvertBodyResponse, err)
				return false
			}
			response = *rP
			break
		}
	}

	return response.R5
}

func (rbot *RemoteBotAPI) Send(c Chattable) (result Message, err error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	cC := hConvertToConcreteChattable(c)

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationSend,
		CorrelationId: randomString(RandomStringLength),
		C:             cC,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R6, response.R2.toError()
}

func (rbot *RemoteBotAPI) GetUserProfilePhotos(config UserProfilePhotosConfig) (UserProfilePhotos, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result UserProfilePhotos

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetUserProfilePhotos,
		CorrelationId: randomString(RandomStringLength),
		Config:        config,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R7, response.R2.toError()
}

func (rbot *RemoteBotAPI) GetFile(config2 FileConfig) (File, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result File

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetFile,
		CorrelationId: randomString(RandomStringLength),
		Config2:       config2,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R8, response.R2.toError()
}

func (rbot *RemoteBotAPI) GetUpdates(config3 UpdateConfig) ([]Update, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result []Update

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetUpdates,
		CorrelationId: randomString(RandomStringLength),
		Config3:       config3,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R9, response.R2.toError()
}

func (rbot *RemoteBotAPI) RemoveWebhook() (APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result APIResponse

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationRemoveWebhook,
		CorrelationId: randomString(RandomStringLength),
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.toError()
}

func (rbot *RemoteBotAPI) SetWebhook(config4 WebhookConfig) (APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result APIResponse

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationSetWebhook,
		CorrelationId: randomString(RandomStringLength),
		Config4:       config4,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.toError()
}

func (rbot *RemoteBotAPI) GetWebhookInfo() (WebhookInfo, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result WebhookInfo

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetWebhookInfo,
		CorrelationId: randomString(RandomStringLength),
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R10, response.R2.toError()
}

func (rbot *RemoteBotAPI) GetUpdatesChan(config3 UpdateConfig) (UpdatesChannel, error) {
	// TODO(tinti) not implemented
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result UpdatesChannel

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetUpdatesChan,
		CorrelationId: randomString(RandomStringLength),
		Config3:       config3,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	result = make(UpdatesChannel)
	return result, response.R2.toError()
}

func (rbot *RemoteBotAPI) ListenForWebhook(pattern string) UpdatesChannel {
	// TODO(tinti) not implemented
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result UpdatesChannel

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationListenForWebhook,
		CorrelationId: randomString(RandomStringLength),
		Pattern:       pattern,
	}
	request, err := json.Marshal(requestMessage)
	if err != nil {
		panic(err)
	}

	remoteBotErr = hChannelPublishWithTimeoutTicker(ch, &requestMessage, q.Name, request, ticker)
	if remoteBotErr != nil {
		return result
	}

	response, remoteBotErr := hChannelConsumeWithTimeoutTicker(requestMessage.CorrelationId, msgs, ticker)
	if remoteBotErr != nil {
		return result
	}

	// TODO(tinti) remove this
	// This is just for refactoring
	_ = response
	return result
}

func (rbot *RemoteBotAPI) AnswerInlineQuery(config5 InlineConfig) (APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result APIResponse

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationAnswerInlineQuery,
		CorrelationId: randomString(RandomStringLength),
		Config5:       config5,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.toError()
}

func (rbot *RemoteBotAPI) AnswerCallbackQuery(config6 CallbackConfig) (APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result APIResponse

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationAnswerCallbackQuery,
		CorrelationId: randomString(RandomStringLength),
		Config6:       config6,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.toError()
}

func (rbot *RemoteBotAPI) KickChatMember(config7 KickChatMemberConfig) (APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result APIResponse

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationKickChatMember,
		CorrelationId: randomString(RandomStringLength),
		Config7:       config7,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.toError()
}

func (rbot *RemoteBotAPI) LeaveChat(config8 ChatConfig) (APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result APIResponse

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationLeaveChat,
		CorrelationId: randomString(RandomStringLength),
		Config8:       config8,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.toError()
}

func (rbot *RemoteBotAPI) GetChat(config8 ChatConfig) (Chat, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result Chat

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetChat,
		CorrelationId: randomString(RandomStringLength),
		Config8:       config8,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R12, response.R2.toError()
}

func (rbot *RemoteBotAPI) GetChatAdministrators(config8 ChatConfig) ([]ChatMember, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result []ChatMember

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetChatAdministrators,
		CorrelationId: randomString(RandomStringLength),
		Config8:       config8,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R13, response.R2.toError()
}

func (rbot *RemoteBotAPI) GetChatMembersCount(config8 ChatConfig) (int, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result int

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetChatMembersCount,
		CorrelationId: randomString(RandomStringLength),
		Config8:       config8,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R14, response.R2.toError()
}

func (rbot *RemoteBotAPI) GetChatMember(config9 ChatConfigWithUser) (ChatMember, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result ChatMember

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetChatMember,
		CorrelationId: randomString(RandomStringLength),
		Config9:       config9,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R15, response.R2.toError()
}

func (rbot *RemoteBotAPI) UnbanChatMember(config10 ChatMemberConfig) (APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result APIResponse

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationUnbanChatMember,
		CorrelationId: randomString(RandomStringLength),
		Config10:      config10,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.toError()
}

func (rbot *RemoteBotAPI) RestrictChatMember(config11 RestrictChatMemberConfig) (APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result APIResponse

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationRestrictChatMember,
		CorrelationId: randomString(RandomStringLength),
		Config11:      config11,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.toError()
}

func (rbot *RemoteBotAPI) PromoteChatMember(config12 PromoteChatMemberConfig) (APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result APIResponse

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationPromoteChatMember,
		CorrelationId: randomString(RandomStringLength),
		Config12:      config12,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.toError()
}

func (rbot *RemoteBotAPI) GetGameHighScores(config13 GetGameHighScoresConfig) ([]GameHighScore, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result []GameHighScore

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetGameHighScores,
		CorrelationId: randomString(RandomStringLength),
		Config13:      config13,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R16, response.R2.toError()
}

func (rbot *RemoteBotAPI) AnswerShippingQuery(config14 ShippingConfig) (APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result APIResponse

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationAnswerShippingQuery,
		CorrelationId: randomString(RandomStringLength),
		Config14:      config14,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.toError()
}

func (rbot *RemoteBotAPI) AnswerPreCheckoutQuery(config15 PreCheckoutConfig) (APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result APIResponse

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationAnswerPreCheckoutQuery,
		CorrelationId: randomString(RandomStringLength),
		Config15:      config15,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.toError()
}

func (rbot *RemoteBotAPI) DeleteMessage(config16 DeleteMessageConfig) (APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result APIResponse

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationDeleteMessage,
		CorrelationId: randomString(RandomStringLength),
		Config16:      config16,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.toError()
}

func (rbot *RemoteBotAPI) GetInviteLink(config8 ChatConfig) (string, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result string

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationGetInviteLink,
		CorrelationId: randomString(RandomStringLength),
		Config8:       config8,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R3, response.R2.toError()
}

func (rbot *RemoteBotAPI) PinChatMessage(config17 PinChatMessageConfig) (APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result APIResponse

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationPinChatMessage,
		CorrelationId: randomString(RandomStringLength),
		Config17:      config17,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.toError()
}

func (rbot *RemoteBotAPI) UnpinChatMessage(config18 UnpinChatMessageConfig) (APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result APIResponse

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationUnpinChatMessage,
		CorrelationId: randomString(RandomStringLength),
		Config18:      config18,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.toError()
}

func (rbot *RemoteBotAPI) SetChatTitle(config19 SetChatTitleConfig) (APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result APIResponse

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationSetChatTitle,
		CorrelationId: randomString(RandomStringLength),
		Config19:      config19,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.toError()
}

func (rbot *RemoteBotAPI) SetChatDescription(config20 SetChatDescriptionConfig) (APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result APIResponse

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationSetChatDescription,
		CorrelationId: randomString(RandomStringLength),
		Config20:      config20,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.toError()
}

func (rbot *RemoteBotAPI) SetChatPhoto(config21 SetChatPhotoConfig) (APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result APIResponse

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationSetChatPhoto,
		CorrelationId: randomString(RandomStringLength),
		Config21:      config21,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.toError()
}

func (rbot *RemoteBotAPI) DeleteChatPhoto(config22 DeleteChatPhotoConfig) (APIResponse, error) {
	ticker := time.NewTicker(rbot.Timeout)
	defer ticker.Stop()

	var result APIResponse

	ch, q, msgs, remoteBotErr := hCreateChannelQueueDeclareConsume(rbot.Connection)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}
	defer ch.Close()

	requestMessage := RequestMessage{
		Operation:     OperationDeleteChatPhoto,
		CorrelationId: randomString(RandomStringLength),
		Config22:      config22,
	}

	response, remoteBotErr := hChannelSerializePublishConsumeWithTimeoutTicker(ch, q, msgs, &requestMessage, ticker)
	if remoteBotErr != nil {
		return result, remoteBotErr
	}

	return response.R, response.R2.toError()
}

func SimpleServerDefault(url string, bot *BotAPI) {
	SimpleServer(url, bot, func(error) {})
}

func SimpleServer(url string, bot *BotAPI, errorHandler func(error)) error {
	conn, err := amqp.Dial(url)
	if err != nil {
		errorHandler(err)
	}

	return hNewErrorRemoteBot(FailedConnect, err)
	defer conn.Close()

	ch, err := conn.Channel()
	return hNewErrorRemoteBot(FailedOpenChannel, err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		RoutingKey, // name
		false,      // durable
		false,      // delete when usused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	return hNewErrorRemoteBot(FailedDeclareQueue, err)

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	return hNewErrorRemoteBot(FailedOptionQoS, err)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	return hNewErrorRemoteBot(FailedMessageConsume, err)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			nP, err := RequestMessageUnmarshal(d.Body)
			if err != nil {
				errorHandler(err)
			}
			n := *nP

			var r ResponseMessage
			switch n.Operation {
			case OperationMakeRequest:
				apiResponse, err := bot.MakeRequest(n.Endpoint, n.Params)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationUploadFile:
				apiResponse, err := bot.UploadFile(n.Endpoint, n.Params2, n.Fieldname, n.File)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetFileDirectURL:
				fileID, err := bot.GetFileDirectURL(n.FileID)

				r = ResponseMessage{}
				r.R3 = fileID
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetMe:
				user, err := bot.GetMe()

				r = ResponseMessage{}
				r.R4 = user
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationIsMessageToMe:
				yes := bot.IsMessageToMe(n.Message)

				r = ResponseMessage{}
				r.R5 = yes

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationSend:
				switch n.C.Type {
				case reflect.TypeOf(MessageConfig{}).String():
					m := n.C.ValueMessageConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(ForwardConfig{}).String():
					m := n.C.ValueForwardConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(PhotoConfig{}).String():
					m := n.C.ValuePhotoConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(AudioConfig{}).String():
					m := n.C.ValueAudioConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(DocumentConfig{}).String():
					m := n.C.ValueDocumentConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(StickerConfig{}).String():
					m := n.C.ValueStickerConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(VideoConfig{}).String():
					m := n.C.ValueVideoConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(AnimationConfig{}).String():
					m := n.C.ValueAnimationConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(VideoNoteConfig{}).String():
					m := n.C.ValueVideoNoteConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(VoiceConfig{}).String():
					m := n.C.ValueVoiceConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(MediaGroupConfig{}).String():
					m := n.C.ValueMediaGroupConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(LocationConfig{}).String():
					m := n.C.ValueLocationConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(VenueConfig{}).String():
					m := n.C.ValueVenueConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(ContactConfig{}).String():
					m := n.C.ValueContactConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(GameConfig{}).String():
					m := n.C.ValueGameConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(SetGameScoreConfig{}).String():
					m := n.C.ValueSetGameScoreConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(GetGameHighScoresConfig{}).String():
					m := n.C.ValueGetGameHighScoresConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(ChatActionConfig{}).String():
					m := n.C.ValueChatActionConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(EditMessageTextConfig{}).String():
					m := n.C.ValueEditMessageTextConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(EditMessageCaptionConfig{}).String():
					m := n.C.ValueEditMessageCaptionConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(EditMessageReplyMarkupConfig{}).String():
					m := n.C.ValueEditMessageReplyMarkupConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(InvoiceConfig{}).String():
					m := n.C.ValueInvoiceConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(DeleteMessageConfig{}).String():
					m := n.C.ValueDeleteMessageConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(PinChatMessageConfig{}).String():
					m := n.C.ValuePinChatMessageConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(UnpinChatMessageConfig{}).String():
					m := n.C.ValueUnpinChatMessageConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(SetChatTitleConfig{}).String():
					m := n.C.ValueSetChatTitleConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(SetChatDescriptionConfig{}).String():
					m := n.C.ValueSetChatDescriptionConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				case reflect.TypeOf(DeleteChatPhotoConfig{}).String():
					m := n.C.ValueDeleteChatPhotoConfig
					message, err := bot.Send(m)

					r = ResponseMessage{}
					r.R6 = message
					r.R2 = hConvertToConcreteError(err)
				//case reflect.TypeOf(SetChatPhotoConfig{}).String():
				//	m := n.C.ValueSetChatPhotoConfig
				//	message, err := bot.Send(m)
				//
				//	r = ResponseMessage{}
				//	r.R6 = message
				//	r.R2 = hConvertToConcreteError(err)
				default:
					r = ResponseMessage{}
					r.R2 = hConvertToConcreteError(fmt.Errorf("not implemented"))
				}

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetUserProfilePhotos:
				userProfilePhotos, err := bot.GetUserProfilePhotos(n.Config)

				r = ResponseMessage{}
				r.R7 = userProfilePhotos
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetFile:
				file, err := bot.GetFile(n.Config2)

				r = ResponseMessage{}
				r.R8 = file
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetUpdates:
				updates, err := bot.GetUpdates(n.Config3)

				r = ResponseMessage{}
				r.R9 = updates
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationRemoveWebhook:
				apiResponse, err := bot.RemoveWebhook()

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationSetWebhook:
				apiResponse, err := bot.SetWebhook(n.Config4)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetWebhookInfo:
				webhookInfo, err := bot.GetWebhookInfo()

				r = ResponseMessage{}
				r.R10 = webhookInfo
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetUpdatesChan:
				r = ResponseMessage{}
				r.R2 = hConvertToConcreteError(fmt.Errorf("not implemented"))

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationListenForWebhook:
				r = ResponseMessage{}
				r.R2 = hConvertToConcreteError(fmt.Errorf("not implemented"))

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationAnswerInlineQuery:
				apiResponse, err := bot.AnswerInlineQuery(n.Config5)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationAnswerCallbackQuery:
				apiResponse, err := bot.AnswerCallbackQuery(n.Config6)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationKickChatMember:
				apiResponse, err := bot.KickChatMember(n.Config7)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationLeaveChat:
				apiResponse, err := bot.LeaveChat(n.Config8)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetChat:
				chat, err := bot.GetChat(n.Config8)

				r = ResponseMessage{}
				r.R12 = chat
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetChatAdministrators:
				chatMembers, err := bot.GetChatAdministrators(n.Config8)

				r = ResponseMessage{}
				r.R13 = chatMembers
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetChatMembersCount:
				count, err := bot.GetChatMembersCount(n.Config8)

				r = ResponseMessage{}
				r.R14 = count
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetChatMember:
				chatMember, err := bot.GetChatMember(n.Config9)

				r = ResponseMessage{}
				r.R15 = chatMember
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationUnbanChatMember:
				apiResponse, err := bot.UnbanChatMember(n.Config10)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationRestrictChatMember:
				apiResponse, err := bot.RestrictChatMember(n.Config11)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationPromoteChatMember:
				apiResponse, err := bot.PromoteChatMember(n.Config12)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetGameHighScores:
				gameHighScores, err := bot.GetGameHighScores(n.Config13)

				r = ResponseMessage{}
				r.R16 = gameHighScores
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationAnswerShippingQuery:
				apiResponse, err := bot.AnswerShippingQuery(n.Config14)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationAnswerPreCheckoutQuery:
				apiResponse, err := bot.AnswerPreCheckoutQuery(n.Config15)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationDeleteMessage:
				apiResponse, err := bot.DeleteMessage(n.Config16)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationGetInviteLink:
				link, err := bot.GetInviteLink(n.Config8)

				r = ResponseMessage{}
				r.R3 = link
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationPinChatMessage:
				apiResponse, err := bot.PinChatMessage(n.Config17)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationUnpinChatMessage:
				apiResponse, err := bot.UnpinChatMessage(n.Config18)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationSetChatTitle:
				apiResponse, err := bot.SetChatTitle(n.Config19)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationSetChatDescription:
				apiResponse, err := bot.SetChatDescription(n.Config20)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationSetChatPhoto:
				apiResponse, err := bot.SetChatPhoto(n.Config21)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			case OperationDeleteChatPhoto:
				apiResponse, err := bot.DeleteChatPhoto(n.Config22)

				r = ResponseMessage{}
				r.R = apiResponse
				r.R2 = hConvertToConcreteError(err)

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			default:
				r = ResponseMessage{}
				r.R2 = hConvertToConcreteError(fmt.Errorf("not implemented"))

				r.Operation, r.CorrelationId = n.Operation, n.CorrelationId
			}

			response, err := json.Marshal(r)
			if err != nil {
				errorHandler(err)
			}

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
			if err != nil {
				errorHandler(hNewErrorRemoteBot(FailedMessagePublish, err))
			}

			d.Ack(false)
		}
	}()

	<-forever

	return nil
}
