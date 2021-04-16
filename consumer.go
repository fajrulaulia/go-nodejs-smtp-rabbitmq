package broker

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func RabbitMQconnector(consumer string) (<-chan amqp.Delivery, error) {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, errors.New("Failed to Connect ")
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.New("failed to connect channels")
	}

	msg, err := ch.Consume("sendmail", "", true, false, false, false, nil)
	if err != nil {
		return nil, errors.New("sendmail channel not create")
	}
	return msg, nil
}

func ParsePayload(payload []byte) (map[string]interface{}, error) {
	var objmap map[string]interface{}
	if err := json.Unmarshal(payload, &objmap); err != nil {
		return objmap, errors.New("sendmail channel not create")
	}

	return objmap, nil

}

func SendMail(to []string, message string) {

	subject := "Pesan Otomatis Fajrul Worker"
	body := "From: " + os.Getenv("CONFIG_SENDER_NAME") + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	smtpAddr := fmt.Sprintf("%s:%s", os.Getenv("CONFIG_SMTP_HOST"), os.Getenv("CONFIG_SMTP_PORT"))
	err := smtp.SendMail(
		smtpAddr,
		smtp.PlainAuth("",
			os.Getenv("CONFIG_AUTH_EMAIL"),
			os.Getenv("CONFIG_AUTH_PASSWORD"),
			os.Getenv("CONFIG_SMTP_HOST")),
		os.Getenv("CONFIG_AUTH_EMAIL"),
		to, []byte(body))
	if err != nil {
		log.Fatal(err.Error())
	}

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
