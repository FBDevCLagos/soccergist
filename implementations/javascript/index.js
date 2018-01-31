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

const server = app.listen(process.env.PORT || 3000, () => {
    console.log(`Listening on port ${server.address().port}`);
});
