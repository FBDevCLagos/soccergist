from flask import Flask, request
import requests
import json
import os


VERIFY_TOKEN = "bots are awesome"
PAGE_ACCESS_TOKEN = os.environ["PAGE_ACCESS_TOKEN"]

app = Flask(__name__)


def send_text_message(recipient_id, message):
    data = json.dumps({
        "recipient": {"id": recipient_id},
        "message": {"text": "I received your message: {}, and "
                    "I've sent it to my Oga at the top: teenoh".format(message)}
    })

    params = {
        "access_token": PAGE_ACCESS_TOKEN
    }

    headers = {
        "Content-Type": "application/json"
    }

    r = requests.post("https://graph.facebook.com/v2.6/me/messages",
                      params=params, headers=headers, data=data)
    

@app.route('/webhook', methods=['GET'])
def verify():
    params = request.args
    if (params.get('hub.mode', '') == "subscribe"):
        if ((params.get('hub.verify_token', '') == VERIFY_TOKEN) and params.get("hub.challenge")):
            return params.get("hub.challenge"), 200

    return "Error", 403


@app.route('/webhook', methods=['POST'])
def handle_messages():
    
    data = request.get_json()

    
    if data["object"] == "page":
        message_object = data["entry"][0]["messaging"][0]
        sender_id = message_object["sender"]["id"]
        message = message_object["message"]["text"]

        send_text_message(sender_id, message)

    return 'OK', 200


if __name__ == "__main__":
    app.run(debug=True)
