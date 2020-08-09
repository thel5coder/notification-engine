package usecase

import (
	"database/sql"
	"github.com/dev-legoas/notification-engine/helpers/google"
	"github.com/dev-legoas/notification-engine/helpers/mailing"
)

type Contract struct {
	DB           *sql.DB
	FCMHelper    google.FCM
	GoMailHelper mailing.GoMail
}
