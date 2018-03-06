const express = require('express');
const bodyParser = require('body-parser');
const dotenv = require('dotenv');
const request = require('request');
const { handleFeedback, sendTextMessage, getTable } = require('./helpers')

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

    // message object
    const message = req.body.entry[0].messaging[0];
    // So here we've got the request i.e req

    sendTextMessage(senderId, handleFeedback(message, getTable)) // Here we prepare and send off the response we want our bot to give the sender
    res.sendStatus(200) // Then we tell Facebook all went well        
})

const server = app.listen(process.env.PORT || 3000, () => {
    console.log(`Listening on port ${server.address().port}`);
});
