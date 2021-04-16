# go-nodejs-smtp-rabbitmq

Simple Golang and NodeJS Message Broker using RabbitMQ

# Schema
this is repo how to use rabbitMQ
- Golang as Consumer
- NodeJS as Publish
- Node JS Send Email Address and Message, then throw a payload to worker and send it to smtp to send to client email

## Enviroment Consumer
```env
CONFIG_SMTP_HOST=smtp.gmail.com OR Anything
CONFIG_SMTP_PORT=587 OR Anything
CONFIG_SENDER_NAME=Fajrul <YourEmail@gmail.com> OR Anything
CONFIG_AUTH_EMAIL=INSERT YOUR EMAIL
CONFIG_AUTH_PASSWORD=INSERT YOUR PASSWORD
```

## Set Payload in Node JS as Publisher
```js
{
    mail:["ADDRESS_MAIL_YOU_WANT_SEND@gmail.com"],
    message:`Tanngal Kirim/ Date Sent ${today.toISOString()} | En: This Message send auto, Id:Ini Pesan dikirim Secara Otomatis`
}
```
