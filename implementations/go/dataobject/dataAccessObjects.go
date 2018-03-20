package dataobject

//TeamLogos - Contains the logo of each team
var TeamLogos = map[string]string{
	"Manchester City FC":      "https://logoeps.com/wp-content/uploads/2011/08/manchester-city-logo-vector.png",
	"Manchester United FC":    "https://logoeps.com/wp-content/uploads/2011/08/manchester-united-logo-vector.png",
	"Liverpool FC":            "https://logoeps.com/wp-content/uploads/2011/08/liverpool-logo-vector.png",
	"Chelsea FC":              "https://logoeps.com/wp-content/uploads/2011/08/chelsea-logo-vector.png",
	"Tottenham Hotspur FC":    "https://logoeps.com/wp-content/uploads/2012/02/tottenham-hotspur-fc-logo-vector.jpg",
	"Arsenal FC":              "https://logoeps.com/wp-content/uploads/2011/05/arsenal-logo-vector.png",
	"Burnley FC":              "http://logovector.net/wp-content/uploads/2014/05/363404-burnley-fc-logo.gif",
	"Leicester City FC":       "http://logovector.net/wp-content/uploads/2013/06/221992-leicester-city-fc-1-logo.gif",
	"Everton FC":              "https://logoeps.com/wp-content/uploads/2012/02/everton-fc-logo-vector.jpg",
	"AFC Bournemouth":         "http://logovector.net/wp-content/uploads/2010/04/221104-bournemouth-fc-logo.gif",
	"Watford FC":              "http://logovector.net/wp-content/uploads/2012/02/222037-watford-fc-0-logo.gif",
	"West Ham United FC":      "https://logoeps.com/wp-content/uploads/2012/12/west-ham-united-logo-vector.png",
	"Newcastle United FC":     "https://logoeps.com/wp-content/uploads/2011/08/newcastle-united-fc-logo-200x200.jpg",
	"Brighton & Hove Albion":  "http://logovector.net/wp-content/uploads/2014/02/326020-brighton-hove-albion-fc-logo.jpg",
	"Crystal Palace FC":       "http://logovector.net/wp-content/uploads/2010/01/350045-crystal-palace-fc-logo.gif",
	"Swansea City FC":         "https://logoeps.com/wp-content/uploads/2012/04/swansea-city-vector.gif",
	"Huddersfield Town":       "http://logovector.net/wp-content/uploads/2013/01/348872-huddersfield-town-fc-1-logo.png",
	"Southampton FC":          "https://logoeps.com/wp-content/uploads/2012/11/southampton-f.c-logo-vector.png",
	"Stoke City FC":           "https://logoeps.com/wp-content/uploads/2012/04/stoke-city-fc-vector.gif",
	"West Bromwich Albion FC": "https://logoeps.com/wp-content/uploads/2012/10/west-brom-logo-vector.png",
}

//JSONRequest struct
type JSONRequest struct {
	Object string      `json:"object"`
	Entry  []EntryItem `json:"entry"`
}

//JSONResponse - struct
type JSONResponse struct {
	Recipient Recipient `json:"recipient"`
	// Message   ResponseMessage `json:"message"`
	Message interface{} `json:"message"`
}

//ResponseMessage struct
type ResponseMessage struct {
	Text       string      `json:"text"`
	Attachment *Attachment `json:"attachment"`
}

//QuickResponseMessage Struct
type QuickResponseMessage struct {
	Text         string       `json:"text"`
	QuickReplies []QuickReply `json:"quick_replies"`
}

//QuickReply Template
type QuickReply struct {
	ContentType string `json:"content_type"`
	Title       string `json:"title"`
	Payload     string `json:"payload"`
}

//EntryItem Struct
type EntryItem struct {
	ID        string          `json:"id"`
	Time      int64           `json:"time"`
	Messaging []MessagingItem `json:"messaging"`
}

//MessagingItem struct for handling the messaging values
type MessagingItem struct {
	Sender    Sender    `json:"sender"`
	Recipient Recipient `json:"recipient"`
	Timestamp int64     `json:"timestamp"`
	Message   Message   `json:"message"`
	PostBack  PostBack  `json:"postback"`
	Read      Read      `json:"read"`
	Delivery  Delivery  `json:"delivery"`
}

//Message struct handles the message that is being sent across.
type Message struct {
	MID        string            `json:"mid"`  //message ID
	Seq        int               `json:"seq"`  //sequence
	Text       string            `json:"text"` //content of the message
	QuickReply MessageQuickReply `json:"quick_reply"`
}

//MessageQuickReply - struct to save the quick reply returned.
type MessageQuickReply struct {
	Payload string `json:"payload"`
}

//PostBack struct
type PostBack struct {
	Payload string `json:"payload"`
	Title   string `json:"title"`
}

//Delivery struct
type Delivery struct {
	Mids      []string `json:"mids"`
	Watermark int64    `json:"watermark"`
	Seq       int      `json:"seq"`
}

//Read - struct for read message
type Read struct {
	Watermark int64 `json:"watermark"`
	Seq       int   `json:"seq"`
}

//Sender struct - To Manage the sender
type Sender struct {
	ID string `json:"id"`
}

//Recipient struct to manage the recipient
type Recipient struct {
	ID string `json:"id"`
}

//Button Template Struct
type Button struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Payload string `json:"payload"`
}

//ButtonPayload Struct
type ButtonPayload struct {
	TemplateType string   `json:"template_type"`
	Text         string   `json:"text"`
	Buttons      []Button `json:"buttons"`
}

//ListPayload struct
type ListPayload struct {
	TemplateType    string    `json:"template_type"`
	TopElementStyle string    `json:"top_element_style"`
	Elements        []Element `json:"elements"`
}

//Attachment struct this constructs the attachment thats added to a message
type Attachment struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

//Element struct
type Element struct {
	Title    string   `json:"title"`
	SubTitle string   `json:"subtitle"`
	ImageURL string   `json:"image_url"`
	Buttons  []Button `json:"buttons"`
}
