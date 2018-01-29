# HOW TO SETUP A WEBHOOK
_- Oscar_

The last piece required for a messenger bot is to setup a webhook. The other pieces are a [Facebook page](how-to-create-a-facebook-page.md) and a [Facebook app](how-to-create-a-facebook-app.md). 

Behind every bot, there is always a set of established logic (usually encapsulated in a webapp) that orchestrates and drives the experience/service the bot is offering. So say, you want your bot to provide users with fuel price updates, then you must need some backend service that can determine the current fuel price and returns this to users when they request for it.

Another way to look at is to consider an already existing business trying to extend their service to be available via the messenger platform. Let's assume this said business offers their customers Premier League updates. They already have an existing app logic for fetching the premier info. All that's required for their bot is to expose this in a structured way on messenger when users ask for it. And this is where the webhook comes in.

## What is a Webhook

A webhook is a URL that will be "called" once an event (or a set of events) occur. In the context of Facebook messenger, webhooks are used to send you a variety of events including messages, authentication events and callback from messages and interactions with your bot on messenger. The webhook url is the gateway between your bot setup on Facebook and your bot's brain (backend) on your hosted server.

Of the three main pieces that make up a bot, the webhook url connects the two pieces that live on Facebook (i.e. Facebook page and Facebook app) to your bot backend. A typical request flow is as follows
* A user interacts with your bot on messenger triggering an event
* Facebook checks if your app is subscribed to the event that was triggered. The flow terminates here is it's not subscribed otherwise it continues.
* Facebook packages information about the user event and sends to your backend via the registered webhook
* Your backend, on receipt of the event payload sends an acknowledgment to Facebook that it received the payload. This is in the form of a 200 response status.
* Your backend parses the event payload from the request and decides on a response
* Your backend packages the response in one of the many Facebook acceptable response template and sends this off to Facebook
* Facebook, on receipt of the response package, parses it (and if all is well and good) displays it to the user.

## Setting up a webhook

For security reasonses, Facebook requires that any webhook you supply pass a verification process. For the verification, Facebook sends a predefined verify token in a request to the supplied webhook url, the server at the webhook backend is expected to listen for the request from Facebook, check that the request contains the predefined verify token and then respond with a challenge value that would be part of the request.

Sample validation request payload
```json
{
  "hub.mode": "subscribe",
  "hub.verify_token": "the_verify_token_you_specify",
  "hub.challenge": "the_challenge_param_to_send_back_to_facebook",
}
```


Let's see how to handle this in a nodejs codebase using express. 
* Step 1 : Create a Get http endpoint
```js
app.get('/webhook', function(req, res) {
});
```

Step 2 : In the route handler function, test for the verify token
```js
app.get('/webhook', function(req, res) {
    if (req.query['hub.mode'] === 'subscribe' &&
          req.query['hub.verify_token'] === <VERIFY_TOKEN>) {
        console.log("Successful token validation");
      } else {
        console.error("Failed validation. Make sure the validation token match.");
        res.sendStatus(403);
      }
});
```

Step 3 : Create a response with the challenge string and send back to Facebook
```js 
    app.get('/webhook', function(req, res) {
      if (req.query['hub.mode'] === 'subscribe' &&
          req.query['hub.verify_token'] === <VERIFY_TOKEN>) {
          res.status(200).send(req.query['hub.challenge']);
        } else {
          res.send('Error, wrong validation token');  
        }
    });
```

Putting this all together in a simple node app

```js
'use strict';
const express = require('express');
const bodyParser = require('body-parser');

// The rest of the code implements the routes for our Express server.
let app = express();

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({
  extended: true
}));

// Webhook validation
app.get('/webhook', function(req, res) {
    // replace <VERIFY_TOKEN> here with the verify token you specified on messenger dashboard
    if (req.query['hub.mode'] === 'subscribe' &&
        req.query['hub.verify_token'] === '<VERIFY_TOKEN>') {
        res.status(200).send(req.query['hub.challenge']);
    } else {
        res.send('Error, wrong validation token');  
    }
});

// Set Express to listen out for HTTP requests
var server = app.listen(process.env.PORT || 3000, function () {
  console.log("Listening on port %s", server.address().port);
});

```
_note 1: for nodejs you'll need to npm install express, body-parser_

_note 2: similar implementation are available for PHP_


### Exposing the web app over the internet

_skip to the next section if you already have your app hosted_

You app needs to be accessible over the internet for Facebook to access it. While in development stage, you can use [ngrok](https://ngrok.com/) to setup a tunnel to allow external access to your local app server. Here's a [quick tutorial on how to install and setup ngrok](https://medium.com/@Oskarr3/developing-messenger-bot-with-ngrok-5d23208ed7c8).

Note: One requirement of the provided url is that it must be https. Ngrok provides ssl url option which you will use. And if using a hosting provider, ensure you have [SSL](http://info.ssl.com/article.aspx?id=10241) installed.

With your app running and ngrok working and pointing to your app, copy the ngrok url and head over to the app dashboard.

### Verifying the webhook
On the dashboard for your Facebook app, add the messenger product if you've not added it. 
<img width="1280" alt="screen shot 2017-09-18 at 12 47 23 pm" src="https://user-images.githubusercontent.com/11221027/30541414-3e0c59d6-9c73-11e7-98b8-649bfabc8138.png">
* Click on the settings for the messenger product and click the setup webhook option. 
* In the callback URL field enter the **https** url to the webhook verification path on your bot. For example, using our example above with ngrok, the url would be `https://<some-random-string>.ngrok.io/webhook`. Then for verfication token, enter any random string of your choice. Note that if must match what you setup here in your code. 
* Under subscription fields, select `messages` and `messaging_postbacks`. _(More on this later. And yes, you can always update it)_
<img width="812" alt="screen shot 2017-09-18 at 12 48 33 pm" src="https://user-images.githubusercontent.com/11221027/30541415-3e0d8c3e-9c73-11e7-899d-6b75cb92f029.png">
* Click verify and save to have verify your webhook

<img width="918" alt="screen shot 2017-09-18 at 12 50 01 pm" src="https://user-images.githubusercontent.com/11221027/30541420-3e47f126-9c73-11e7-9001-2896d9201d92.png">

 Once you get the green checkmark, you're good to go. And the next place to go is to obtain your access token.

 If verification fails, you'll get a red indicator with some helpful error message.
 <img width="807" alt="screen shot 2017-09-18 at 12 48 48 pm" src="https://user-images.githubusercontent.com/11221027/30541416-3e0eee58-9c73-11e7-99da-6ab5bf180b79.png">

<img width="813" alt="screen shot 2017-09-18 at 12 49 20 pm" src="https://user-images.githubusercontent.com/11221027/30541419-3e20dfa0-9c73-11e7-8099-e7141915a8df.png">


 Note: if you run in to errors setting this up please reach out on the [Facebook DevC Lagos group](https://www.facebook.com/groups/DevCLagos/) for assistance.
