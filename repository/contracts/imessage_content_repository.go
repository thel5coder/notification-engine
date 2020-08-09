package contracts

import "github.com/dev-legoas/notification-engine/model/sqlmodel"

type IMessageContentRepository interface {
	BrowseBy(column,value string) (data []sqlmodel.MessageContent,err error)
}
