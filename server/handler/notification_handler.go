package handler

import (
	"context"
	legoas_srv_notification "github.com/dev-legoas/notification-engine/model/proto/notification"
	"github.com/dev-legoas/notification-engine/usecase"
)

type NotificationHandler struct {
	*usecase.Contract
}

func (handler *NotificationHandler) Send(ctx context.Context, req *legoas_srv_notification.SendRequest, resp *legoas_srv_notification.SendResponse) error {
	uc := usecase.NotificationUseCase{Contract:handler.Contract}
	err := uc.SendNotification(req.MessageCode,req.From,req.To)
	if err != nil {
		resp.Message = err.Error()
		resp.IsSuccess = false
		resp.StatusCode = 244
		return err
	}

	resp.IsSuccess = true
	return nil
}
