<?php

namespace App;
use App\Config;
use App\PremiershipHighlights;
class SendMessage {

	public $userID;
	public $accessToken;
  public $input;

	function __construct($input) {

    $this->userID = $input['entry'][0]['messaging'][0]['sender']['id'];
    $this->accessToken = Config::getAccessToken();
    $this->input = $input;

	}


  public function text($text) {

	  $userID = $this->userID;
	   
	  $response = [

	     'recipient' => [ 'id' => $userID ],

	     'message' => [ 'text' => $text ]

	  ];

    $this->curlResponse($response);

   
  }
	
  public function carousel($offset = 0) {

        
            
    $premiershipHighlights = new PremiershipHighlights();

    $dataBuffer  =  $premiershipHighlights->getAllMatches();

    $data  = [];

    $dataSize = count($dataBuffer);
     
     //maximum of 5 carousel displays at a time     
    $dataBuffer = array_slice($dataBuffer,$offset, 5);

   

    if(($dataSize - $offset) > 5) {// then we need a more matches button
       
      foreach ($dataBuffer as $dataItem) {

        array_push($data, $dataItem);
        
      }

      //second parameter tells the last point of array
      $response = $this->buildCarousel($data,$offset + 5);

    }

    else {

      $data = $dataBuffer;
      $response = $this->buildCarousel($data);
     
    }
    
     $this->curlResponse($response);

    
  }

  function buildCarousel($data,$moreData = 0)  {

    $userID = $this->userID; 
    $elements = []; 
    $button = [];
    $moreMatchesButton = [

      'title' => 'More Matches',
      'type' => 'postback',
      'payload' => 'moreMatches|'.$moreData
                                
    ];
                      
    $count = 1;

    foreach ($data as $item) { //build up each highlight

      $button = [

        [
          
          'title' => 'View Highlights',
          'type' => 'web_url',
          'url'  => $item['matchURL']
                                
        ]

      ];

      if($count == 5) {// add more matches button if need be

        $button[1] = $moreMatchesButton;
      }

                    

      $count++;

      $elements[] = [

        'title' => $item['matchTitle'],
        'image_url' => 'https://res.cloudinary.com/testi/image/upload/v1522831395/premier-league_jbqgi1.jpg',
        'buttons' => $button,
              
      ];


    }
         
                    
    $response = [//build full response

      'recipient' => [
      'id' =>$userID

      ],

      'message' => [

        'attachment' => [

          'type' => 'template',
          'payload' => [

            'template_type' => 'generic',
            'elements' => $elements

          ]
        ]
      ]

    ];

    return $response;
                       
    }

    public function curlResponse($response) {

      $accessToken = $this->accessToken;
	    $input=  $this->input;
	
	    $curl = curl_init('https://graph.facebook.com/v2.6/me/messages?access_token='.$accessToken);
	    curl_setopt($curl, CURLOPT_POST, 1);
	    curl_setopt($curl, CURLOPT_POSTFIELDS, json_encode($response));
	    curl_setopt($curl, CURLOPT_HTTPHEADER, ['Content-Type: application/json']);

      if(!empty($input)) {
	        
          curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);
	        $result = curl_exec($curl);
      }

      curl_close($curl); 

	  }


}
