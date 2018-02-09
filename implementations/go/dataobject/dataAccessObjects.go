package dataobject

//JSONRequest struct
type JSONRequest struct {
	Object string      `json:"object"`
	Entry  []EntryItem `json:"entry"`
}

//JSONResponse - struct
type JSONResponse struct {
	Recipient Recipient       `json:"recipient"`
	Message   ResponseMessage `json:"message"`
}

//ResponseMessage struct
type ResponseMessage struct {
	Text       string      `json:"text"`
	Attachment *Attachment `json:"attachment"`
}

// ResponseMessageWithAttachment - function
type ResponseMessageWithAttachment struct {
	Text       string     `json:"text"`
	Attachment Attachment `json:"attachment"`
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
	MID  string `json:"mid"`  //message ID
	Seq  int    `json:"seq"`  //sequence
	Text string `json:"text"` //content of the message
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

//Payload Struct
type Payload struct {
	TemplateType string   `json:"template_type"`
	Text         string   `json:"text"`
	Buttons      []Button `json:"buttons"`
}

//Attachment struct this constructs the attachment thats added to a message
type Attachment struct {
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
}
