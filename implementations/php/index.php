<?php 
require 'vendor/autoload.php'; 
use  App\VerifyBot;
use App\ReceivedMessage;
use App\SendMessage;
use App\PremiershipHighlights;

$verifyBot = new VerifyBot();

 
$input = json_decode(file_get_contents('php://input'), true);


$receivedMessage = new ReceivedMessage($input);
$sendMessage = new SendMessage($input);

 
if(!empty($receivedMessage->textMessage)) {   
     $sendMessage->text("hello coder");
     $sendMessage->carousel();
}
 
if(!empty($receivedMessage->buttonMessage)) {

	$carouselOffset = explode("|",$receivedMessage->buttonMessage);
	$sendMessage->carousel($carouselOffset[1]);


}


 
