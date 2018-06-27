# The Button Template
_- Oscar_

"The button template sends a text message with up to three attached buttons. This template is useful for offering the message recipient options to choose from, such as pre-determined responses to a question, or actions to take."

The buttons in the response can be any of the 6 out of the (about) 8 standard buttons available on the Messenger:
- **URL button**: The URL Button opens a web page in the Messenger webview when the user clicks it.
- **Postback button**: The postback button sends a `messaging_postbacks `type of event to your webhook. You configure it such that when the user clicks the button, a certain message will be sent to your backend
- **Call button**: The call button dials a phone number when tapped
- **Log in button**: The log in button is used in the account linking flow to link the message recipient's identity
- **Log out button**: The log out button is used in the account linking flow to unlink the message recipient's identity

The other buttons not supported by the Button template include: Share button, Buy button, Game button.