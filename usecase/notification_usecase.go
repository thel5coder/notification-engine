package usecase

import (
	"github.com/dev-legoas/notification-engine/repository/actions"
	"github.com/dev-legoas/notification-engine/usecase/view_model"
)

type NotificationUseCase struct {
	*Contract
}

func (uc NotificationUseCase) BrowseBy(column, value string) (res []view_model.NotificationVm, err error) {
	repository := actions.NewMessageContentRepository(uc.DB)
	messages, err := repository.BrowseBy(column, value)
	if err != nil {
		return res, err
	}

	for _, message := range messages {
		res = append(res, view_model.NotificationVm{
			ID:                 message.ID,
			Title:              message.Title,
			MessageContentCode: message.MessageContentCode,
			Content:            message.Content,
			ProviderType:       message.ProviderType,
			CreatedAt:          message.CreatedAt,
			UpdatedAt:          message.UpdatedAt,
			DeletedAt:          message.DeletedAt.String,
		})
	}

	return res, err
}

func (uc NotificationUseCase) SendNotification(messageContentCode, from string, to []string) (err error) {
	messages, err := uc.BrowseBy("mc.message_content_code", messageContentCode)
	if err != nil {
		return err
	}

	for _, message := range messages {
		if message.ProviderType == "push" {
			payload := map[string]interface{}{}
			err = uc.FCMHelper.SendFcm(payload, "", message.Title, message.Content)
			if err != nil {
				return err
			}
		}else if message.ProviderType == "email" {
			for _, deliverTo := range to{
				err = uc.GoMailHelper.SendEmail(deliverTo,message.Title,message.Content)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
