package main

type webhookPayload struct {
	Object string `json:"object,omitempty"`
	Entry  []struct {
		ID        string `json:"id,omitempty"`
		Messaging []struct {
			Message   messageEvent `json:"message,omitempty"`
			Recipient struct {
				ID string `json:"id,omitempty"`
			} `json:"recipient,omitempty"`
			Sender struct {
				ID string `json:"id,omitempty"`
			} `json:"sender,omitempty"`
			Timestamp int           `json:"timestamp,omitempty"`
			Postback  postbackEvent `json:"postback,omitempty"`
		} `json:"messaging,omitempty"`
		Time int `json:"time,omitempty"`
	} `json:"entry,omitempty"`
}

type messageEvent struct {
	Mid  string
	Seq  int
	Text string
}

type postbackEvent struct {
	Title   string `json:"title,omitempty"`
	Payload string `json:"payload,omitempty"`
}

type textResponse struct {
	Recipient struct {
		ID string `json:"id,omitempty"`
	} `json:"recipient,omitempty"`
	Message struct {
		Text string `json:"text,omitempty"`
	} `json:"message,omitempty"`
}

type templateResponse struct {
	Recipient struct {
		ID string `json:"id,omitempty"`
	} `json:"recipient,omitempty"`
	Message struct {
		Attachment struct {
			Type    string `json:"type,omitempty"`
			Payload struct {
				TemplateType string   `json:"template_type,omitempty"`
				Text         string   `json:"text,omitempty"`
				Buttons      []button `json:"buttons,omitempty"`
			} `json:"payload,omitempty"`
		} `json:"attachment,omitempty"`
	} `json:"message,omitempty"`
}

type button struct {
	Type    string `json:"type,omitempty"`
	Title   string `json:"title,omitempty"`
	Payload string `json:"payload,omitempty"`
}
