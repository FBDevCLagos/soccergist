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
	Mid        string
	Seq        int
	Text       string
	QuickReply struct {
		Payload string `json:"payload,omitempty"`
	} `json:"quick_reply,omitempty"`
}

type postbackEvent struct {
	Title   string `json:"title,omitempty"`
	Payload string `json:"payload,omitempty"`
}

type quickReply struct {
	ContentType string `json:"content_type,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
	Payload     string `json:"payload,omitempty"`
	Title       string `json:"title,omitempty"`
}

type templateResponse struct {
	Recipient struct {
		ID string `json:"id,omitempty"`
	} `json:"recipient,omitempty"`
	Message struct {
		Text       string `json:"text,omitempty"`
		Attachment struct {
			Type    string `json:"type,omitempty"`
			Payload struct {
				TemplateType    string    `json:"template_type,omitempty"`
				TopElementStyle string    `json:"top_element_style,omitempty"`
				Text            string    `json:"text,omitempty"`
				Buttons         []button  `json:"buttons,omitempty"`
				Elements        []element `json:"elements,omitempty"`
			} `json:"payload,omitempty"`
		} `json:"attachment,omitempty"`
	} `json:"message,omitempty"`
}

type basicResponseTemplate struct {
	Recipient struct {
		ID string `json:"id,omitempty"`
	} `json:"recipient,omitempty"`
	Message struct {
		Text         string       `json:"text,omitempty"`
		QuickReplies []quickReply `json:"quick_replies,omitempty"`
	} `json:"message,omitempty"`
}

type button struct {
	Type    string `json:"type,omitempty"`
	Title   string `json:"title,omitempty"`
	Payload string `json:"payload,omitempty"`
	URL     string `json:"url,omitempty"`
}

type element struct {
	Buttons       []button `json:"buttons,omitempty"`
	Title         string   `json:"title,omitempty"`
	Subtitle      string   `json:"subtitle,omitempty"`
	ImageURL      string   `json:"image_url,omitempty"`
	MediaType     string   `json:"media_type,omitempty"`
	AttachmentID  string   `json:"attachment_id,omitempty"`
	DefaultAction *struct {
		Type                string `json:"type,omitempty"`
		URL                 string `json:"url,omitempty"`
		MessengerExtensions bool   `json:"messenger_extensions,omitempty"`
		WebviewHeightRatio  string `json:"webview_height_ratio,omitempty"`
	} `json:"default_action,omitempty"`
}
