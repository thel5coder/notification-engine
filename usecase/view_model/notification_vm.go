package view_model

type NotificationVm struct {
	ID                 string `json:"id"`
	Title              string `json:"title"`
	MessageContentCode string `json:"message_content_code"`
	Content            string `json:"content"`
	ProviderType       string `json:"provider_type"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	DeletedAt          string `json:"deleted_at"`
}
