package sqlmodel

import "database/sql"

type MessageContent struct {
	ID                 string         `db:"id"`
	Title              string         `db:"title"`
	MessageContentCode string         `db:"message_content_code"`
	Content            string         `db:"content"`
	ProviderType       string         `db:"provider_type"`
	CreatedAt          string         `db:"created_at"`
	UpdatedAt          string         `db:"updated_at"`
	DeletedAt          sql.NullString `db:"deleted_at"`
}
