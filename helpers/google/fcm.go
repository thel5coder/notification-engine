package google

import (
	"github.com/appleboy/go-fcm"
)

type FCM struct {
	Client *fcm.Client
}

func(fcmHelper FCM) SendFcm(payload map[string]interface{}, fcmToken, title, body string) error {
	msg := &fcm.Message{
		To: fcmToken,
		Notification: &fcm.Notification{
			Title: title,
			Body:  body,
		},
		Data: payload,
	}

	_,err := fcmHelper.Client.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
