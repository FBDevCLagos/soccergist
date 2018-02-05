const express = require('express');
const bodyParser = require('body-parser');
const dotenv = require('dotenv');

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
        res.status(201).json({
            "recipient": {
                "id": senderId
            },
            "message": {
                "text": `I have received your message: "${senderMessage}", and I've sent it to my Oga at the top: Oscar`
            }
        })
})

const server = app.listen(process.env.PORT || 3000, () => {
    console.log(`Listening on port ${server.address().port}`);
});
