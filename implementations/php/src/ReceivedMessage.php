<?php

//receives message
//returns message type by control i.e text, image, video, etc
//from type we know kind of response
//we build response
//we send response
namespace App;
class ReceivedMessage{


	public $textMessage; //receives text messages
	public $buttonMessage; //receives button messages
	public $quickReplyMessage; //receives quick reply messages
	public $attachmentMessage;//receives attachment messages i.e Video,picture,audio

  function __construct($input) {
     
    $this->textMessage = $input['entry'][0]['messaging'][0]['message']['text'];
    $this->buttonMessage = $input['entry'][0]['messaging'][0]['postback']['payload'];
    $this->quickReplyMessage = $input['entry'][0]['messaging'][0]['message']['quick_reply']['payload'];
    $this->attachmentMessage = $input['entry'][0]['messaging'][0]['message']['attachments'][0]['type'];


  }


}