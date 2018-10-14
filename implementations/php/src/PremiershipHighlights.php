<?php
namespace App;
use App\redditApi\Reddit;
use App\SendMessage;
//reddit wrapper



/**
*Get Premier Highlights
*/
class PremiershipHighlights {

	private $reddit;
	private $highlights;
	private $highlight_url_pattern;
	private $englishMatches;



	function __construct() {

		//AuthO credentials
		$this->reddit = new Reddit("prondubuisi", "blackman123", "XDTgTfN3Eyt7hQ", "Q6zGs4YpdDxmpjcKP9pzX4PQNAY");

        //API call
		$this->highlights = $this->reddit->getListing('footballhighlights', array("limit" => 100));

		//regex to get urls
        $this->highlight_url_pattern = '/href="([\w\/:?\.-=]+|[\w\/:\.-]+)"/';

        //array for english Matches
        $this->englishMatches = [];
		
	}

	function getAllMatches() {
        $highlights = $this->highlights;

		foreach ($highlights->body->data->children as $matchListing) {

            if($this->isMatchEnglish($matchListing)) {

	            	if($matchDetails = $this->getIndividualMatchDetails($matchListing)) {

	            		array_push($this->englishMatches, $matchDetails);
	            	}
            }	       
       }

       return $this->englishMatches;
       
	}

	function isMatchEnglish($matchListing) {

	    if(stripos($matchListing->data->title, 'premier') ) 
		    return true;
    }

    function getIndividualMatchDetails($matchListing) {

        $highlight_url_pattern = $this->highlight_url_pattern;

	    //get all urls linked to match
		preg_match_all($highlight_url_pattern, $matchListing->data->selftext_html, $matchURLs);

		
		$matchTitle = $matchListing->data->title;
		$matchURL = isset($matchURLs[1][0]) ? $matchURLs[1][0] : false;

		if($matchURL) 

			return [
				      
			'matchTitle' => $matchTitle,
			'matchURL' => $matchURL

		    ];

		return false;


   }


}













