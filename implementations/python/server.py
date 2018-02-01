from flask import Flask, request

VERIFY_TOKEN = "bots are awesome"

app = Flask(__name__)


@app.route('/webhook', methods=['GET'])
def verify():
    params = request.args
    if (params.get('hub.mode', '') == "subscribe"):
        if ((params.get('hub.verify_token', '') == VERIFY_TOKEN) and params.get("hub.challenge")):
            return params.get("hub.challenge"), 200

    return "Error", 403


if __name__ == "__main__":
    app.run(debug=True)
