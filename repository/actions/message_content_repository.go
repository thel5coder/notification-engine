package actions

import (
	"database/sql"
	"github.com/dev-legoas/notification-engine/model/sqlmodel"
	"github.com/dev-legoas/notification-engine/repository/contracts"
)

type MessageContentRepository struct{
	DB *sql.DB
}

func NewMessageContentRepository(DB *sql.DB) contracts.IMessageContentRepository{
	return MessageContentRepository{DB: DB}
}

func (repository MessageContentRepository) BrowseBy(column, value string) (data []sqlmodel.MessageContent, err error) {
	statement := `select mc.*,mp."provider_type" from "message_contents" mc 
                 inner join "message_content_providers" mcp on mcp."message_content_id"=mc."id"
                 inner join "message_providers" mp on mp."id"=mcp."message_provider_id" and mp."deleted_at" is null
                 where `+column+`=$1 and mc."deleted_at" is null`
	rows,err := repository.DB.Query(statement,value)
	if err != nil {
		return data,err
	}

	for rows.Next(){
		dataTemp := sqlmodel.MessageContent{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.MessageContentCode,
			&dataTemp.Content,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			&dataTemp.Title,
			&dataTemp.ProviderType,
			)
		if err != nil {
			return data,err
		}

		data = append(data,dataTemp)
	}

	return data,err
}

