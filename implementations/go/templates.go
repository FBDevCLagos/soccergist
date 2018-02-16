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
