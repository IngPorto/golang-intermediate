/*
*

Software que envía notificaciones -> SMS, Push Notifications, Email
*/
package main

import "fmt"

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

// Para SMS
type SMSNotification struct {
	
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification via SMS")
}

func (SMSNotification) GetSender() ISender {
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

// Para el Email

type EmailNotification struct {
	
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending Notification via Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
	
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "Gmail"
}

// --- CLASES CONCRETAS ---

func getNotificationFactory (notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	} else if notificationType == "Email" {
		return &EmailNotification{}, nil
	} else {
		return nil, fmt.Errorf("Invalid notification type: %s", notificationType)
	}
}

// Se puede enviar un mensaje por SMS o por Email
func sendNotification (factory INotificationFactory) {
	factory.SendNotification()
}

func getMethod (factory INotificationFactory) {
	fmt.Println(factory.GetSender().GetSenderMethod())
}

func getChannel (factory INotificationFactory) {
	fmt.Println(factory.GetSender().GetSenderChannel())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	getMethod(smsFactory)
	getChannel(smsFactory)

	sendNotification(emailFactory)
	getMethod(emailFactory)
	getChannel(emailFactory)
}