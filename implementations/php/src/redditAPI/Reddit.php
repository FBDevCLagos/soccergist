<?php
 namespace App\redditApi;
    class Reddit {
        private $apiHost = "https://oauth.reddit.com";
        private $token;
        private $username;
        private $password;
        private $clientid;
        private $clientsecret;
        private $headers;
        
        public function __construct($username, $password, $clientid, $clientsecret) {
            
            $this->login($username, $password, $clientid, $clientsecret);
            $this->username = $username;
            $this->password = $password;
            $this->clientid = $clientid;
            $this->clientsecret = $clientsecret;
        }

        public function login($username, $password, $clientid, $clientsecret) {
            $data = array(
                "grant_type" => "password",
                "username" => $username,
                "password" => $password            
            );

            $loginresult = $this->runCurl("https://www.reddit.com/api/v1/access_token" , $data,"{$clientid}:{$clientsecret}");

            if($this->hasError($loginresult->body)) {
                die("Not authenticated: {$loginresult->body->error}");
            } else {
                $this->token = $loginresult->body->access_token;
                return true; 
            }
        }        
      
        // Get the raw json from anywhere in reddit as an object. 
        public function getRaw($rawurl) {
       
        try {
              $response = $this->sendAPIRequest("{$this->apiHost}/{$rawurl}");
              return $response; 
            }
            catch(Exception $e) {
                return 'Error getting raw url: ' .$e->getMessage();
            }   
                return $response;
        }
        
       
        // Get me
        public function getMe() {

        try {
              $response = $this->sendAPIRequest("{$this->apiHost}/api/v1/me", null, null, $headers);
              return $response; 
            }
            catch(Exception $e) {
                return 'Error getting my data: ' .$e->getMessage();
            }   
                return $response;
        }

        // Get my karma 
        public function getMeKarma() {

        try {
              $response = $this->sendAPIRequest("{$this->apiHost}/api/v1/me/karma", null, null, $headers);
              return $response; 
            }
            catch(Exception $e) {
                return 'Error getting my karma data: ' .$e->getMessage();
            }   
                return $response;
        }
        
        // Get my prefs
        public function getMePrefs() {

        try {
              $response = $this->sendAPIRequest("{$this->apiHost}/api/v1/me/prefs", null, null, $headers);
              return $response; 
            }
            catch(Exception $e) {
                return 'Error getting my preferences: ' .$e->getMessage();
            }   
                return $response;
        }    
        
        
        // Get a listing of a subreddit.
        public function getListing($subreddit = "all", $parameters = array()) {
            
            $data = http_build_query($parameters);
            
            try {
              $response = $this->sendAPIRequest("{$this->apiHost}/r/$subreddit/.json?$data");
              return $response; 
            }
            catch(Exception $e) {
                return 'Error getting listing: ' .$e->getMessage();
            }   
        }
        
        // Submit a comment 
        public function submitComment($thing_id = null, $comment = null) {
  
            $data = array(
                "thing_id" => $thing_id,
                "text" => $comment,
                "api_type" => "json"
            );
            
            try {
              $response = $this->sendAPIRequest("{$this->apiHost}/api/comment", $data);
              return $response; 
            }
            catch(Exception $e) {
                return 'Error submitting comment: ' .$e->getMessage();
            }         
                 
        }        
        
        // Send the request to the api and handle possible errors. 
        private function sendAPIRequest($url, $data = null) {
            
            // We do some error checking below, it is still very basic and just fails when there is an issue with the authorisation which also means that it doesn't check for expired tokens as of yet.  
            // The first warrants a login attempt, the latter should probably shut the whole thing down with an error. 
            $headers = array("Authorization: bearer {$this->token}");

            $result =  $this->runCurl($url, $data, null, $headers);
            
            if ($this->hasError($result->body)) {
                $resultError = $result->body->error;
                if ($resultError === 401 || $resultError === "invalid_grant" ) {
                    if ($this->login($this->username, $this->password, $this->clientid, $this->clientsecret)) {
                        return $this->sendAPIRequest($url, $data);
                    } else {
                        // In practice this should never be triggered since login failed, better save than sorry though.
                        die("Authentication revoked or expired: {$resultError}");
                    }
                } else { 
                    throw new Exception("reddit server error: {$resultError} \n");
                }
            } elseif (property_exists($result->body, 'json') && $this->hasApiErrors($result->body->json)) {
                // This assumes that there is only one value in errors. It is very likely that this is not the case, but this needs to be double checked.
                throw new Exception(implode(", ", $result->body->json->errors[0]));
            } 
            return $result;
        }

        // Curl stuff, I probably need to go over this again. 
        private function runCurl($url, $postVals = null, $auth = null, $headers = null){
            $curl = curl_init();

                    if ($postVals) {
                        curl_setopt($curl, CURLOPT_POST, count($postVals));
                        curl_setopt($curl, CURLOPT_POSTFIELDS, $postVals);
                        curl_setopt($curl, CURLOPT_CUSTOMREQUEST, "POST");
                    }
                    
                    if ($auth) {
                        curl_setopt($curl, CURLOPT_USERPWD, $auth);
                        curl_setopt($curl, CURLOPT_HTTPAUTH, CURLAUTH_BASIC);
                    }
                    
                    if ($headers) {
                        $headers[] = 'Expect:';                                           
                    } else {
                        $headers = array('Expect:');    
                    }
                    
                    curl_setopt($curl, CURLOPT_HTTPHEADER, $headers);
                    
                    curl_setopt($curl, CURLOPT_SSLVERSION, 4);
                    curl_setopt($curl, CURLOPT_SSL_VERIFYPEER,false);
                    curl_setopt($curl, CURLOPT_SSL_VERIFYHOST, 2);
                    
                    curl_setopt($curl, CURLOPT_USERAGENT, "PHPapiWrap/0.1 by creesch");            
                    curl_setopt($curl, CURLOPT_FOLLOWLOCATION, true);
                    curl_setopt($curl, CURLOPT_RETURNTRANSFER, true );
                    curl_setopt($curl, CURLOPT_VERBOSE, 0 );
                    curl_setopt($curl, CURLOPT_HEADER, 1 );
                    curl_setopt($curl, CURLOPT_URL, $url);

                    
                    $response = curl_exec( $curl );
                    $headersize = curl_getinfo($curl, CURLINFO_HEADER_SIZE);
                    curl_close($curl);
                    $header = substr($response, 0, $headersize);
                    $body = substr($response, $headersize);
                    
                    $output = new \stdClass();
                    $headerArray = array();
                    foreach (explode("\r\n", $header) as $i => $line) {
                        if ($i === 0) {
                            $headerArray['http_code'] = $line;
                        } else {
                            if($line !=='') {                    
                            list ($key, $value) = explode(': ', $line);
                            $headerArray[$key] = $value;
                            }
                        }
                    }
                    
                    $output->header = (object) $headerArray;
                    $output->body = json_decode($body);

                    return $output;


        }
        
        // check for errors directly in the response, not for errors in the data->errors TODO: figure out which situation trigger what error response and handle those properly. 
        private function hasError($response) {
            if(!isset($response->error) || empty($response->error)) {
                return false;
            } else {
                return true;
            }
        }
        
        private function hasApiErrors($response) {
            if(!isset($response->errors) || empty($response->errors)) {
                return false;
            } else {
                return true;
            }
        }


    }

?>
