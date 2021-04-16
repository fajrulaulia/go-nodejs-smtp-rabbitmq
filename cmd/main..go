package main

import (
	"fmt"
	"log"
	"os"

	c "github.com/fajrulaulia/go-smtp-rabbitmq"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	log.Println("Starting Service")

	log.Println("-----------------------------------------------------------------------")
	log.Println("CONFIG_SMTP_HOST  :", os.Getenv("CONFIG_SMTP_HOST"))
	log.Println("CONFIG_SMTP_PORT  :", os.Getenv("CONFIG_SMTP_PORT"))
	log.Println("CONFIG_SENDER_NAME:", os.Getenv("CONFIG_SENDER_NAME"))
	log.Println("CONFIG_AUTH_EMAIL :", os.Getenv("CONFIG_AUTH_EMAIL"))

	rbm, err := c.RabbitMQconnector("send-mail")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Connector succesfully running, standby to subscribe channel send-mail")
	log.Println("-----------------------------------------------------------------------")

	frv := make(chan bool)
	go func() {
		for v := range rbm {
			res, err := c.ParsePayload(v.Body)
			if err != nil {
				log.Println(err)
				frv <- true
			}
			log.Println("Data yang dikirim : ", res["mail"])

			s := make([]string, len(res["mail"].([]interface{})))
			for i, v := range res["mail"].([]interface{}) {
				s[i] = fmt.Sprint(v)
			}
			c.SendMail(s, res["message"].(string))
		}
	}()
	<-frv
}
