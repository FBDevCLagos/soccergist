package main

func buildPostbackBtn(title, payload string) button {
	return button{
		Type:    "postback",
		Title:   title,
		Payload: payload,
	}
}

func buildBasicElement(title, subtitle, imageURL string) element {
	return element{
		Title:    title,
		ImageURL: imageURL,
		Subtitle: subtitle,
	}
}

func buildTextMsg(senderID, text string) (msg basicResponseTemplate) {
	msg.Message.Text = text
	msg.Recipient.ID = senderID
	return msg
}

func buildQuickReply(text, senderID string, quickReplies []quickReply) (msg basicResponseTemplate) {
	msg.Recipient.ID = senderID
	msg.Message.Text = text
	msg.Message.QuickReplies = quickReplies
	return
}
