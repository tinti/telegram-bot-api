package tgbotapi

import (
	"fmt"
	"net/url"

	_ "github.com/streadway/amqp"
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
	fmt.Println("MakeRequest")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) UploadFile(endpoint string, params map[string]string, fieldname string, file interface{}) (APIResponse, error) {
	fmt.Println("UploadFile")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) GetFileDirectURL(fileID string) (string, error) {
	fmt.Println("GetFileDirectURL")
	return "", nil
}

func (bot *EchoBotAPI) GetMe() (User, error) {
	fmt.Println("GetMe")
	return User{}, nil
}

func (bot *EchoBotAPI) IsMessageToMe(message Message) bool {
	fmt.Println("IsMessageToMe")
	return false
}

func (bot *EchoBotAPI) Send(c Chattable) (Message, error) {
	fmt.Println("Send")
	return Message{}, nil
}

func (bot *EchoBotAPI) GetUserProfilePhotos(config UserProfilePhotosConfig) (UserProfilePhotos, error) {
	fmt.Println("GetUserProfilePhotos")
	return UserProfilePhotos{}, nil
}

func (bot *EchoBotAPI) GetFile(config FileConfig) (File, error) {
	fmt.Println("GetFile")
	return File{}, nil
}

func (bot *EchoBotAPI) GetUpdates(config UpdateConfig) ([]Update, error) {
	fmt.Println("GetUpdates")
	return []Update{}, nil
}

func (bot *EchoBotAPI) RemoveWebhook() (APIResponse, error) {
	fmt.Println("RemoveWebhook")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) SetWebhook(config WebhookConfig) (APIResponse, error) {
	fmt.Println("SetWebhook")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) GetWebhookInfo() (WebhookInfo, error) {
	fmt.Println("GetWebhookInfo")
	return WebhookInfo{}, nil
}

func (bot *EchoBotAPI) GetUpdatesChan(config UpdateConfig) (UpdatesChannel, error) {
	fmt.Println("GetUpdatesChan")
	ch := make(chan Update)
	return ch, nil
}

func (bot *EchoBotAPI) ListenForWebhook(pattern string) UpdatesChannel {
	fmt.Println("ListenForWebhook")
	ch := make(chan Update)
	return ch
}

func (bot *EchoBotAPI) AnswerInlineQuery(config InlineConfig) (APIResponse, error) {
	fmt.Println("AnswerInlineQuery")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) AnswerCallbackQuery(config CallbackConfig) (APIResponse, error) {
	fmt.Println("AnswerCallbackQuery")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) KickChatMember(config KickChatMemberConfig) (APIResponse, error) {
	fmt.Println("KickChatMember")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) LeaveChat(config ChatConfig) (APIResponse, error) {
	fmt.Println("LeaveChat")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) GetChat(config ChatConfig) (Chat, error) {
	fmt.Println("GetChat")
	return Chat{}, nil
}

func (bot *EchoBotAPI) GetChatAdministrators(config ChatConfig) ([]ChatMember, error) {
	fmt.Println("GetChatAdministrators")
	return []ChatMember{}, nil
}

func (bot *EchoBotAPI) GetChatMembersCount(config ChatConfig) (int, error) {
	fmt.Println("GetChatMembersCount")
	return 0, nil
}

func (bot *EchoBotAPI) GetChatMember(config ChatConfigWithUser) (ChatMember, error) {
	fmt.Println("GetChatMember")
	return ChatMember{}, nil
}

func (bot *EchoBotAPI) UnbanChatMember(config ChatMemberConfig) (APIResponse, error) {
	fmt.Println("UnbanChatMember")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) RestrictChatMember(config RestrictChatMemberConfig) (APIResponse, error) {
	fmt.Println("RestrictChatMember")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) PromoteChatMember(config PromoteChatMemberConfig) (APIResponse, error) {
	fmt.Println("PromoteChatMember")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) GetGameHighScores(config GetGameHighScoresConfig) ([]GameHighScore, error) {
	fmt.Println("GetGameHighScores")
	return []GameHighScore{}, nil
}

func (bot *EchoBotAPI) AnswerShippingQuery(config ShippingConfig) (APIResponse, error) {
	fmt.Println("AnswerShippingQuery")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) AnswerPreCheckoutQuery(config PreCheckoutConfig) (APIResponse, error) {
	fmt.Println("AnswerPreCheckoutQuery")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) DeleteMessage(config DeleteMessageConfig) (APIResponse, error) {
	fmt.Println("DeleteMessage")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) GetInviteLink(config ChatConfig) (string, error) {
	fmt.Println("GetInviteLink")
	return "", nil
}

func (bot *EchoBotAPI) PinChatMessage(config PinChatMessageConfig) (APIResponse, error) {
	fmt.Println("PinChatMessage")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) UnpinChatMessage(config UnpinChatMessageConfig) (APIResponse, error) {
	fmt.Println("UnpinChatMessage")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) SetChatTitle(config SetChatTitleConfig) (APIResponse, error) {
	fmt.Println("SetChatTitle")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) SetChatDescription(config SetChatDescriptionConfig) (APIResponse, error) {
	fmt.Println("SetChatDescription")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) SetChatPhoto(config SetChatPhotoConfig) (APIResponse, error) {
	fmt.Println("SetChatPhoto")
	return APIResponse{}, nil
}

func (bot *EchoBotAPI) DeleteChatPhoto(config DeleteChatPhotoConfig) (APIResponse, error) {
	fmt.Println("DeleteChatPhoto")
	return APIResponse{}, nil
}

var _ BotAPIIface = (*EchoBotAPI)(nil)
