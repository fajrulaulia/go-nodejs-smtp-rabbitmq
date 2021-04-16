var amqp = require('amqplib/callback_api');


const CONN_URL = 'amqp://guest:guest@localhost:5672/';

function Producer(message) {
  amqp.connect(CONN_URL, (error0, connection) => {
    if (error0) {
      throw error0;
    }
    connection.createChannel((error1, channel) => {
      if (error1) {
        throw error1;
      }
      let queue = "sendmail";
      channel.assertQueue(queue, {
        durable: 0,
      });
      channel.sendToQueue(queue, new Buffer(JSON.stringify(message)));
      console.log(" [x] Sent %s", message);
    });
  });
}
const timeElapsed = Date.now();
const today = new Date(timeElapsed);


Producer({
    mail:["ADDRESS_MAIL_YOU_WANT_SEND@gmail.com"],
    message:`Tanngal Kirim/ Date Sent ${today.toISOString()} | En: This Message send auto, Id:Ini Pesan dikirim Secara Otomatis`
})