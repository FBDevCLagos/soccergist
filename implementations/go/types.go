package main

type webhookPayload struct {
	Object string `json:"object"`
	Entry  []struct {
		ID        string `json:"id"`
		Messaging []struct {
			Message   messageEvent `json:"message"`
			Recipient struct {
				ID string `json:"id"`
			} `json:"recipient"`
			Sender struct {
				ID string `json:"id"`
			} `json:"sender"`
			Timestamp int `json:"timestamp"`
		} `json:"messaging"`
		Time int `json:"time"`
	} `json:"entry"`
}

type messageEvent struct {
	Mid  string
	Seq  int
	Text string
}

type textResponse struct {
	Recipient struct {
		ID string `json:"id"`
	} `json:"recipient"`
	Message struct {
		Text string `json:"text"`
	} `json:"message"`
}
