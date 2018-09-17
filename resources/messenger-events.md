# Messenger Events
_- Oscar_

Bots on Messenger, just like any webhook based web app, are driven (almost) entirely by a collection of events that a webhook can subscribe to. Events here refer to a collection of different activities that can occur on the platform for which the app on the other side of the webhook might be interested in.
Some of the events that can occur on the messenger platform (from a receiver's perspective) include `receiving a message`, `message sent is delivered`, `message sent is read`, ... etc

 As a user using messenger, you can easily rely on the faithful blue check mark beside a message to know it's current state as it cycles through the various states highlighted above. For a bot, it relies on the event callbacks. 

There are about 14 events available on Messenger. The common ones include:
- **messages**: This event is triggered when a bot receives a message
- **messaging_postbacks**: This is triggered when the user clicks a postback button in chat with the bot
- **message_reads**: This is triggered when the user reads the bot's message or response
- **message_deliveres**: This is triggered when the bot's response is delivered successfully to the user. It's akin to when you have the white check mark with the blue background

So when setting up your bot's webhook, you select which of these events you care about. You can select as little or as many as you want but it's wise to only go for the ones you have a need for. You can always update it later. 