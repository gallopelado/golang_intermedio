package main

import "fmt"

// SMS, email

type INotificactionFactory interface {
	SendNotification()
	GetSender() Isender
}

type Isender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification via SMS")
}

func (SMSNotification) GetSender() Isender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

// Lo mismo pero para Email

type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending Notification via Email")
}

func (EmailNotification) GetSender() Isender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "EMAIL"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

func getNotificationFactory(notificationType string) (INotificactionFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}
	if notificationType == "EMAIL" {
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("No Notification Type")
}

func sendNotification(f INotificactionFactory) {
	f.SendNotification()
}

func getMethod(f INotificactionFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("EMAIL")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
