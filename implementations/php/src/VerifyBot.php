<?php

namespace App;
use App\Config;
//verifies bot credentials with that of facebook
class VerifyBot
{

    function __construct()
    {

         $verifyToken = Config::getVerifyToken();

		 if ($_REQUEST['hub_verify_token'] === $verifyToken)
		  {

				  echo $_REQUEST['hub_challenge'];



		  } 
 

    }

    
    



}