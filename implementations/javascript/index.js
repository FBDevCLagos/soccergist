const express = require('express');
const bodyParser = require('body-parser');
const dotenv = require('dotenv');
const request = require('request');

dotenv.config()

const app = express();

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({
    extended: true
}))

// Webhook Validation

app.get('/webhook', (req, res) => {
    if (req.query['hub.mode'] === 'subscribe' &&
        req.query['hub.verify_token'] == (process.env.VERIFY_TOKEN) ) {
            res.status(200).send(req.query['hub.challenge']);
        } else {
            res.status(403).send('Error, you have passed wrong parameters')
        }
});

app.post('/webhook', (req, res) => {
    const senderId =  req.body.entry[0].messaging[0].sender.id;
    const senderMessage = req.body.entry[0].messaging[0].message.text;
    // So here we've got the request i.e req
        
    const responseText = `I have received your message: "${senderMessage}", and I've sent it to my Oga at the top: Oscar`

    sendTextMessage(senderId, responseText) // Here we prepare and send off the response we want our bot to give the sender
    res.sendStatus(200) // Then we tell Facebook all went well        
})

const sendTextMessage = (recipientId, messageText) => {
    // we package the bot response in FB required format
    const messageData = {
      recipient: {
        id: recipientId
      },
      message: {
        text: messageText,
      }
    };

    // We send off the response to FB
    request({
        uri: 'https://graph.facebook.com/v2.6/me/messages',
        qs: { access_token: process.env.PAGE_ACCESS_TOKEN },
        method: 'POST',
        json: messageData
    
      }, (error, response, body) => {
        if (!error && response.statusCode == 200) {
            console.log("Successfully sent message");
        } else {
          console.error("Failed calling Send API", response.statusCode, response.statusMessage, body.error);
        }
      });
}

const server = app.listen(process.env.PORT || 3000, () => {
    console.log(`Listening on port ${server.address().port}`);
});
