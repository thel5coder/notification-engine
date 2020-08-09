package main

import (
	"github.com/appleboy/go-fcm"
	"github.com/joho/godotenv"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/server/grpc"
	"github.com/dev-legoas/notification-engine/helpers/google"
	"github.com/dev-legoas/notification-engine/helpers/mailing"
	"github.com/dev-legoas/notification-engine/helpers/str"
	"github.com/dev-legoas/notification-engine/model"
	legoas_srv_notification "github.com/dev-legoas/notification-engine/model/proto/notification"
	"github.com/dev-legoas/notification-engine/server/handler"
	"github.com/dev-legoas/notification-engine/usecase"
	"gopkg.in/gomail.v2"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading ..env file")
	}

	//create fcm instance
	fcmClient, err := fcm.NewClient(os.Getenv("FCM_SERVER_KEY"))
	if err != nil {
		log.Fatalln(err)
	}

	//go mail instance
	dialer := gomail.NewDialer(
		os.Getenv("MAIL_SMTP_HOST"),
		str.StringToInt(os.Getenv("MAIL_SMTP_PORT")),
		os.Getenv("MAIL_SENDER"),
		os.Getenv("MAIL_PASSWORD"),
	)
	goMailHelper := mailing.GoMail{Dialer: dialer}

	//create db instance
	db := model.Connection{
		Host:     os.Getenv("DB_HOST"),
		DbName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		SslMode:  os.Getenv("DB_SSL_MODE"),
	}
	database, err := db.DbConnect()
	if err != nil {
		panic(err)
	}

	//init usecase contract
	ucContract := usecase.Contract{
		DB:           database,
		FCMHelper:    google.FCM{Client: fcmClient},
		GoMailHelper: goMailHelper,
	}

	//create service instance and init service
	service := micro.NewService(
		micro.Name("notification.engine.dev"),
		micro.Version("0.1"),
		micro.Server(
			grpc.NewServer(
				server.Name(os.Getenv("APP_NAME")),
				),
			),
	)
	service.Init()
	legoas_srv_notification.RegisterNotificationHandler(service.Server(), &handler.NotificationHandler{Contract: &ucContract})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
